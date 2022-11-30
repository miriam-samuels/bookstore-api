package modules

import(
	"github.com/jinzhu/gorm"
	"github.com/miriam-samuels/bookstore/pkg/config"
)
//operations with db resides in modules
var db *gorm.DB

type Book struct{
	gorm.Model 
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init()  {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func(b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}