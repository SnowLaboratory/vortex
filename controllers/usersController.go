package controllers

import (
	"fmt"
	"net/http"
	"os"
	"snowlabs/vortex/components"
	"snowlabs/vortex/initializers"
	"snowlabs/vortex/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var requestBody struct {
		Email    string
		Password string
	}

	if c.Bind(&requestBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body.",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}
	user := models.User{Email: requestBody.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	var requestBody struct {
		Email    string
		Password string
	}

	if c.Bind(&requestBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body.",
		})
		return
	}
	fmt.Println("req body", requestBody)
	var user models.User
	initializers.DB.First(&user, "email = ?", requestBody.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email/password.",
		})
		return
	}
	fmt.Println("the user", user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email/password.",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject":    user.ID,
		"expiration": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token.",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.Redirect(http.StatusFound, "/dashboard")
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "", components.Login())
}

func Dashboard(c *gin.Context) {
	// user, _ := c.Get("user")
	// user.(models.User).Email
	c.HTML(http.StatusOK, "", components.Dashboard())
}
