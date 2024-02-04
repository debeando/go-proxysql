package proxysql

import (
	"errors"
	"fmt"
)

const QueryConnectionPoolReset = "SELECT * FROM stats_mysql_connection_pool_reset;"

type Stats struct {
	connection      *Connection
	HostgroupID     int    `db:"hostgroup"`
	Hostname        string `db:"srv_host"`
	Port            int    `db:"srv_port"`
	Status          string `db:"status"`
	ConnUsed        uint64 `db:"ConnUsed"`
	ConnFree        uint64 `db:"ConnFree"`
	ConnOK          uint64 `db:"ConnOK"`
	ConnERR         uint64 `db:"ConnERR"`
	MaxConnUsed     uint64 `db:"MaxConnUsed"`
	Queries         uint64 `db:"Queries"`
	QueriesGTIDSync uint64 `db:"Queries_GTID_sync"`
	BytesDataSent   uint64 `db:"Bytes_data_sent"`
	BytesDataRecv   uint64 `db:"Bytes_data_recv"`
	Latency         uint64 `db:"Latency_us"`
}

func (s *Stats) New(c *Connection) {
	s.connection = c
}

func (s *Stats) Select() Stats {
	if s.connection == nil {
		return *s
	}

	s.connection.Instance.QueryRow(s.QuerySelect()).Scan(
		&s.HostgroupID,
		&s.Hostname,
		&s.Port,
		&s.Status,
		&s.ConnUsed,
		&s.ConnFree,
		&s.ConnOK,
		&s.ConnERR,
		&s.MaxConnUsed,
		&s.Queries,
		&s.QueriesGTIDSync,
		&s.BytesDataSent,
		&s.BytesDataRecv,
		&s.Latency)

	return *s
}

func (s *Stats) QuerySelect() string {
	return fmt.Sprintf(
		"SELECT hostgroup, srv_host, srv_port, status, ConnUsed, ConnFree, ConnOK, ConnERR, MaxConnUsed, Queries, Queries_GTID_sync, Bytes_data_sent, Bytes_data_recv, Latency_us "+
			"FROM stats_mysql_connection_pool WHERE hostgroup = %d AND srv_host = '%s' LIMIT 1;",
		s.HostgroupID,
		s.Hostname,
	)
}

func (s *Stats) Reset() error {
	if s.connection == nil {
		return errors.New("Connection does not exist")
	}

	s.connection.Instance.Query(QueryConnectionPoolReset)

	return nil
}
