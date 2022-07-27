package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type File struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Author    User      `json:"author"`
	Content   string    `gorm:"size:255;not null;" json:"content"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:TIMESTAMP" json:"created_at"`
}

func (f *File) Prepare() {
	f.ID = 0 
	f.Content = html.EscapeString(strings.TrimSpace(f.Content))
	f.Author = User{}
	f.CreatedAt = time.Now() 
}

func (f *File) Validate() error {
 
	if f.Content == "" {
		return errors.New("Required Content")
	}
	if f.AuthorID < 1 {
		return errors.New("Required Author")
	}
	return nil
}


func (f *File) SaveFile(db *gorm.DB) (*File, error) {
	var err error
	err = db.Debug().Model(&File{}).Create(&f).Error
	if err != nil {
		return &File{}, err
	}
	if f.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", f.AuthorID).Take(&f.Author).Error
		if err != nil {
			return &File{}, err
		}
	}
	return f, nil
}

func (f *File) FindAllFiles(db *gorm.DB) (*[]File, error) {
	var err error
	files := []File{}
	err = db.Debug().Model(&File{}).Limit(100).Find(&files).Error
	if err != nil {
		return &[]File{}, err
	}
	if len(files) > 0 {
		for i, _ := range files {
			err := db.Debug().Model(&User{}).Where("id = ?", files[i].AuthorID).Take(&files[i].Author).Error
			if err != nil {
				return &[]File{}, err
			}
		}
	}
	return &files, nil
}


func (f *File) DeleteFile(db *gorm.DB, fid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&File{}).Where("id = ? and author_id = ?", fid, uid).Take(&File{}).Delete(&File{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("File not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
