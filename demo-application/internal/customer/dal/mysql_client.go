package dal

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
)

type MySQLClientConfig struct {
	User      string
	Password  string
	Address   string
	DB        string
	TimeoutMs int
}

func NewMySQLClient(inputConf *MySQLClientConfig) (*sql.DB, error) {
	conf := mysql.NewConfig()
	conf.User = inputConf.User
	conf.Passwd = inputConf.Password
	conf.Addr = inputConf.Address
	conf.DBName = inputConf.DB
	conf.Timeout = time.Duration(inputConf.TimeoutMs) * time.Millisecond

	connector, err := mysql.NewConnector(conf)
	if err != nil {
		return nil, fmt.Errorf("new connector: %w", err)
	}

	return sql.OpenDB(connector), nil
}
