package Database

import "newappp/Model"

func SyncDatabase() {
	DB.AutoMigrate(&Model.User{})
}
