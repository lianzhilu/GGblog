package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Mode   string
	Port   string
	JwtKey string
}

type MySQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type QiniuConfig struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Server    string
}

var (
	AppConf   AppConfig
	MySQLConf MySQLConfig
	QiniuConf QiniuConfig
)

func Init(dir string) {
	// 指定配置文件路径
	viper.SetConfigFile(dir)

	// 查找并读取配置文件
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	appMap := viper.GetStringMapString("server")
	{
		AppConf.Mode = appMap["mode"]
		AppConf.Port = appMap["port"]
		AppConf.JwtKey = appMap["jwtkey"]
	}

	mysqlMap := viper.GetStringMapString("mysql")
	{
		MySQLConf.Host = mysqlMap["host"]
		MySQLConf.Port = mysqlMap["port"]
		MySQLConf.User = mysqlMap["user"]
		MySQLConf.Password = mysqlMap["password"]
		MySQLConf.Name = mysqlMap["name"]
	}

	qiniuMap := viper.GetStringMapString("qiniu")
	{
		QiniuConf.AccessKey = qiniuMap["accesskey"]
		QiniuConf.SecretKey = qiniuMap["secretkey"]
		QiniuConf.Bucket = qiniuMap["bucket"]
		QiniuConf.Server = qiniuMap["server"]
	}
}
