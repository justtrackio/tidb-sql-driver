package tidb_sql_driver

import (
	"context"
	"database/sql/driver"
)

type tidbConn struct {
	driver.Conn
	driver.ConnBeginTx
	driver.Execer
	txnCausalEnabled bool
}

func newTidbConn(conn driver.Conn, txnCausalEnabled bool) tidbConn {
	return tidbConn{
		Conn:             conn,
		ConnBeginTx:      conn.(driver.ConnBeginTx),
		Execer:           conn.(driver.Execer),
		txnCausalEnabled: txnCausalEnabled,
	}
}

func (c tidbConn) Begin() (driver.Tx, error) {
	return c.BeginTx(context.Background(), driver.TxOptions{})
}

func (c tidbConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	sql := "START TRANSACTION"

	if c.txnCausalEnabled {
		sql = "START TRANSACTION WITH CAUSAL CONSISTENCY ONLY"
	}

	_, err := c.Execer.Exec(sql, nil)

	return &TiDBTx{c.Execer}, err
}

func (c tidbConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	return c.Execer.Exec(query, args)
}
