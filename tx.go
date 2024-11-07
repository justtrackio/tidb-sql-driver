package tidb_sql_driver

import "database/sql/driver"

type TiDBTx struct {
	driver.Execer
}

func (t TiDBTx) Commit() error {
	_, err := t.Execer.Exec("COMMIT", nil)
	return err
}

func (t TiDBTx) Rollback() error {
	_, err := t.Execer.Exec("ROLLBACK", nil)
	return err
}
