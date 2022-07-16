package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/abdalrazzak/gin-golang-test/api/models"
)

var users = []models.User{
	models.User{
		Age: 	  30,
		Email:    "abboudbath4@gmail.com",
		Password: "password",
	},
	models.User{
		Age: 	   30,
		Email:    "test43@gmail.com",
		Password: "password",
	},
}
 

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	} 

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
