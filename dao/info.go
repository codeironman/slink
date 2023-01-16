<<<<<<< HEAD
package model

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shortlink/api/utils"
	"shortlink/model"
	"sync"
)

var Db *gorm.DB
var Lock sync.Mutex
func Getdb() *gorm.DB {
	return Db
}
func Init() {
	//Configure the viper
	Config := viper.New()
	Config.SetConfigFile("./config/config.yaml")
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./config")
	err := Config.ReadInConfig() // 查找并读取配置文件
	if err != nil {              // 处理读取配置文件的错误
		logrus.Error("Read config file failed: %s \n", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Info("no error in config file")
		} else {
			logrus.Error("found error in config file\n", ok)
		}
	}
<<<<<<<< HEAD:dao/info.go
	utils.Salt = Config.GetString("Salt")
	jwtInfo := Config.GetStringMapString("Jwt")
	utils.JwtSecret = jwtInfo["secret"]
========
	dao.Salt = Config.GetString("Salt")
	jwtInfo := Config.GetStringMapString("Jwt")
	dao.JwtSecret = jwtInfo["secret"]
>>>>>>>> refs/remotes/origin/master:model/database.go
	loginInfo := Config.GetStringMapString("mysql")
	Dsn := loginInfo["predsn"] + loginInfo["database"] + loginInfo["mode"]
	dao.Db, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("gorm OPENS MySQL failed")
	}
<<<<<<<< HEAD:dao/info.go
	err = Getdb().AutoMigrate(&model.User{}, &model.UrlInfo{}, &model.LoginInfo{}, &model.Cookie{}, &model.PauseUrl{})
========
	err = dao.Getdb().AutoMigrate(&User{}, &UrlInfo{}, &LoginInfo{}, &Cookie{}, &PauseUrl{})
>>>>>>>> refs/remotes/origin/master:model/database.go
	if err != nil {
		logrus.Error("build tables corrupt!\n", err)
	}
	DefaultUser.Id = NOTLOGIN
	DefaultUser.Name = "origin"
}

=======
>>>>>>> refs/remotes/origin/master
/*
* @brief init the config of viper and database
set default viper config path and read data from config.yaml
to get mysql userInfo,jwt secretKey
open mysql database with gorm and create table with user Struct
*/
