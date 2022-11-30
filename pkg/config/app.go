package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// returned variable that helps interact with the database
var(
	db *gorm.DB
)

//helps open a connection with database
func Connect(){
	//					db		username:password		table
	d, err := gorm.Open("mysql","root:G22mEyct140522FU/bookstore?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db 
}