package Connection

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
)
var db *gorm.DB
func init() {
	fmt.Println("Fail to parse 'mysql'")
	//db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	//	user,
	//	password,
	//	host,
	//	dbName))
	//
	//if err != nil {
	//	log.Println(err)
	//}
	connect()
}


func New(){
	fmt.Println("mypackage.mysql")
}


func connect() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  		if err != nil {
		log.Println(err)
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return tablePrefix + defaultTableName
	//}
	//
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
