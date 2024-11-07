package tidb_sql_driver

import (
	"context"
	"database/sql/driver"
)

type TiDbConn struct {
	driver.Conn
	driver.ConnBeginTx
	driver.Execer
}

func NewTiDbConn(conn driver.Conn) TiDbConn {
	return TiDbConn{
		Conn:        conn,
		ConnBeginTx: conn.(driver.ConnBeginTx),
		Execer:      conn.(driver.Execer),
	}
}

func (c TiDbConn) Begin() (driver.Tx, error) {
	_, err := c.Execer.Exec("START TRANSACTION WITH CAUSAL CONSISTENCY ONLY", nil)

	return &TiDBTx{c.Execer}, err
}

func (c TiDbConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	_, err := c.Execer.Exec("START TRANSACTION WITH CAUSAL CONSISTENCY ONLY", nil)

	return &TiDBTx{c.Execer}, err
}

func (c TiDbConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	return c.Execer.Exec(query, args)
}
