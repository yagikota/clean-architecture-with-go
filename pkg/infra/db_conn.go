package infra

import (
	"database/sql"
	"fmt"

	"github.com/yagikota/clean_architecture_wtih_go/pkg/config"
)

const driverName = "mysql"

type MySQLConnector struct {
	Conn *sql.DB
}

func NewMySQLConnector() *MySQLConnector {
	conf := config.LoadConfig()

	dsn := mysqlConnInfo(*conf.MySQLInfo)
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(); err != nil {
		panic(err)
	}

	return &MySQLConnector{
		Conn: conn,
	}
}

func mysqlConnInfo(mysqlInfo config.MysqlInfo) string {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local",
		mysqlInfo.MySQLUser,
		mysqlInfo.MySQLPassWord,
		mysqlInfo.MySQLAddr,
		mysqlInfo.MySQLDBName)

	return dataSourceName
}
