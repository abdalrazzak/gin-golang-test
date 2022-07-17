package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/abdalrazzak/gin-golang-test/api/models"
)

var users = []models.User{
	models.User{ 
		Email:    "abboudbath4@gmail.com",
		Password: "password",
	},
	models.User{ 
		Email:    "test43@gmail.com",
		Password: "password",
	},
}

var files = []models.File{
	models.File{ 
		Content: "3fewfew52532.jpg",
	},
	models.File{
		Content: "32r32rfsfds.png",
	},
}
 

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.File{},&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.File{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	} 
	err = db.Debug().Model(&models.File{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		files[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.File{}).Create(&files[i]).Error
		if err != nil {
			log.Fatalf("cannot seed files table: %v", err)
		}
	}
}
