package tidb_sql_driver

import (
	"context"
	"database/sql/driver"
)

type TiDbConnector struct {
	driver.Connector
}

func (c TiDbConnector) Connect(ctx context.Context) (driver.Conn, error) {
	conn, err := c.Connector.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return NewTiDbConn(conn), nil
}

func (c TiDbConnector) Driver() driver.Driver {
	return &TiDbDriver{}
}
