package models

import (
	"hanxiaolin/gin-demo/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
}

func Setup() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = setting.DatabaseSetting.Type
	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	tablePrefix = setting.DatabaseSetting.TablePrefix
	fmt.Println(setting.DatabaseSetting)

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName;
	}
	fmt.Println(dbType, 222)
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
}

func CloseDb() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
//func updateTimeStampForCreateCallback(scope *gorm.Scope) {
//	if !scope.HasError() {
//		nowTime := time.Now().Unix()
//		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
//			if createTimeField.IsBlank {
//				createTimeField.Set(nowTime)
//			}
//		}
//
//		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
//			if modifyTimeField.IsBlank {
//				modifyTimeField.Set(nowTime)
//			}
//		}
//	}
//}
//
//// updateTimeStampForUpdateCallback will set `ModifyTime` when updating
//func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
//	if _, ok := scope.Get("gorm:update_column"); !ok {
//		scope.SetColumn("ModifiedOn", time.Now().Unix())
//	}
//}
