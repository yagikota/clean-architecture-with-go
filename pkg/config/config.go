package config

import (
	"os"
)

type appConfig struct {
	HTTPInfo  *HTTPInfo
	MySQLInfo *MysqlInfo
}

type HTTPInfo struct {
	Addr string
}

type MysqlInfo struct {
	MySQLUser     string
	MySQLPassWord string
	MySQLAddr     string
	MySQLDBName   string
}

func LoadConfig() *appConfig {
	addr := ":" + os.Getenv("PORT")

	httpInfo := &HTTPInfo{
		Addr: addr,
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	mysqlDBName := os.Getenv("MYSQL_DATABASE")

	dbInfo := &MysqlInfo{
		MySQLUser:     mysqlUser,
		MySQLPassWord: mysqlPassword,
		MySQLAddr:     mysqlAddr,
		MySQLDBName:   mysqlDBName,
	}

	conf := appConfig{
		MySQLInfo: dbInfo,
		HTTPInfo:  httpInfo,
	}

	return &conf
}
