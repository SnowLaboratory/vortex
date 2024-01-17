package initializers

import "snowlabs/vortex/models"

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}
