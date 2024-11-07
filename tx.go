package tidb_sql_driver

import (
	"database/sql/driver"

	"github.com/go-sql-driver/mysql"
)

type TiDBTx struct {
	driver.Execer
}

func (t TiDBTx) Commit() error {
	if t.Execer == nil {
		return mysql.ErrInvalidConn
	}

	_, err := t.Execer.Exec("COMMIT", nil)
	t.Execer = nil

	return err
}

func (t TiDBTx) Rollback() error {
	if t.Execer == nil {
		return mysql.ErrInvalidConn
	}

	_, err := t.Execer.Exec("ROLLBACK", nil)
	t.Execer = nil

	return err
}
