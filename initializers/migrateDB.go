package initializers

import (
	"fmt"
	"snowlabs/vortex/models"
)

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
	fmt.Println("DB migrated")
}
