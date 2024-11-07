package tidb_sql_driver

import (
	"context"
	"database/sql/driver"

	"github.com/go-sql-driver/mysql"
)

type TiDbDriver struct {
	mysql.MySQLDriver
}

func (d TiDbDriver) Open(dsn string) (driver.Conn, error) {
	conn, err := d.MySQLDriver.OpenConnector(dsn)
	if err != nil {
		return nil, err
	}

	tidbConn := TiDbConnector{conn}

	return tidbConn.Connect(context.Background())
}

func (d TiDbDriver) OpenConnector(dsn string) (driver.Connector, error) {
	conn, _ := d.MySQLDriver.OpenConnector(dsn)

	return TiDbConnector{conn}, nil
}
