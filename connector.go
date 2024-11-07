package tidb_sql_driver

import (
	"context"
	"database/sql/driver"
)

type tidbConnector struct {
	driver.Connector
	txnCausalEnabled bool
}

func (c tidbConnector) Connect(ctx context.Context) (driver.Conn, error) {
	conn, err := c.Connector.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return newTidbConn(conn, c.txnCausalEnabled), nil
}

func (c tidbConnector) Driver() driver.Driver {
	return &TidbDriver{}
}
