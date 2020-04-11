package portals_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_portals_username = "mysql_portals_username"
	mysql_portals_password = "mysql_portals_password"
	mysql_portals_host     = "mysql_portals_host"
	mysql_portals_schema   = "mysql_portals_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_portals_username)
	password = os.Getenv(mysql_portals_password)
	host     = os.Getenv(mysql_portals_host)
	schema   = os.Getenv(mysql_portals_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var err error

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("DB succesfully connected")
}
