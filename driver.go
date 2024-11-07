package tidb_sql_driver

import (
	"context"
	"database/sql/driver"
	"fmt"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

type TidbDriver struct {
	mysql.MySQLDriver
}

func (d TidbDriver) Open(dsn string) (driver.Conn, error) {
	var err error
	var connector driver.Connector

	if connector, err = d.getTidbConnector(dsn); err != nil {
		return nil, err
	}

	return connector.Connect(context.Background())
}

func (d TidbDriver) OpenConnector(dsn string) (driver.Connector, error) {
	return d.getTidbConnector(dsn)
}

func (d TidbDriver) getTidbConnector(dsn string) (driver.Connector, error) {
	var err error
	var tidbCfg *mysql.Config
	var mySqlconnector driver.Connector
	var txnCausalEnabled bool

	if tidbCfg, err = mysql.ParseDSN(dsn); err != nil {
		return nil, fmt.Errorf("failed to parse dsn %s: %v", dsn, err)
	}

	if causal, ok := tidbCfg.Params["tidb_txn_causal"]; ok {
		delete(tidbCfg.Params, "tidb_txn_causal")

		if txnCausalEnabled, err = strconv.ParseBool(causal); err != nil {
			return nil, fmt.Errorf("failed to parse tidb_txn_causal %q: %w", causal, err)
		}
	}

	mysqlDsn := tidbCfg.FormatDSN()

	if mySqlconnector, err = d.MySQLDriver.OpenConnector(mysqlDsn); err != nil {
		return nil, err
	}

	return tidbConnector{
		Connector:        mySqlconnector,
		txnCausalEnabled: txnCausalEnabled,
	}, nil
}
