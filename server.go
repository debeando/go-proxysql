package proxysql

import (
	"errors"
	"fmt"
)

type Server struct {
	connection        *Connection
	HostgroupID       int    `yaml:"hostgroup_id"`
	Hostname          string `yaml:"hostname"`
	MaxConnections    int    `yaml:"max_connections"`
	MaxReplicationLag int    `yaml:"max_replication_lag"`
	Port              int    `yaml:"port"`
	Status            string `yaml:"status"`
	Weight            int    `yaml:"weight"`
	stats             Stats
}

func (s *Server) New(c *Connection) {
	s.defaults()
	s.connection = c
	s.stats.New(s.connection)
	s.stats.HostgroupID = s.HostgroupID
	s.stats.Hostname = s.Hostname
	s.stats.Port = s.Port
}

func (s *Server) defaults() {
	if s.Hostname == "" {
		s.Hostname = "127.0.0.1"
	}

	if s.Port == 0 {
		s.Port = 3306
	}

	if s.Status == "" {
		s.Status = "ONLINE"
	}

	if s.MaxConnections == 0 {
		s.MaxConnections = 1000
	}

	if s.Weight == 0 {
		s.Weight = 1
	}
}

func (s *Server) Stats() *Stats {
	return &(s).stats
}

func (s *Server) Equal(in Server) bool {
	return s.HostgroupID == in.HostgroupID &&
		s.Hostname == in.Hostname &&
		s.Port == in.Port
}

func (s *Server) IsOnLine() bool {
	stats := s.Stats()
	stats.Select()

	return s.Status == ONLINE && stats.Status == ONLINE
}

func (s *Server) Select() error {
	if s.connection == nil {
		return errors.New("Connection does not exist")
	}

	return s.connection.Instance.QueryRow(s.QuerySelect()).Scan(
		&s.HostgroupID,
		&s.Hostname,
		&s.Port,
		&s.Status,
		&s.Weight,
		&s.MaxConnections,
		&s.MaxReplicationLag)
}

func (s *Server) Insert() error {
	if s.connection == nil {
		return errors.New("Connection does not exist")
	}

	_, err := s.connection.Instance.Query(s.QueryInsert())
	return err
}

func (s *Server) Update() error {
	if s.connection == nil {
		return errors.New("Connection does not exist")
	}

	_, err := s.connection.Instance.Query(s.QueryUpdate())
	return err
}

func (s *Server) Delete() error {
	if s.connection == nil {
		return errors.New("Connection does not exist")
	}

	_, err := s.connection.Instance.Query(s.QueryDelete())
	return err
}

func (s *Server) QuerySelect() string {
	return fmt.Sprintf(
		"SELECT hostgroup_id, hostname, port, status, weight, max_connections, max_replication_lag FROM mysql_servers WHERE hostgroup_id = %d AND hostname = '%s' LIMIT 1;",
		s.HostgroupID,
		s.Hostname,
	)
}

func (s *Server) QueryInsert() string {
	return fmt.Sprintf(
		"INSERT INTO mysql_servers (hostgroup_id, hostname, port, status, weight, max_connections, max_replication_lag) VALUES (%d, '%s', %d, '%s', %d, %d, %d);",
		s.HostgroupID,
		s.Hostname,
		s.Port,
		s.Status,
		s.Weight,
		s.MaxConnections,
		s.MaxReplicationLag,
	)
}

func (s *Server) QueryUpdate() string {
	return fmt.Sprintf(
		"UPDATE mysql_servers "+
			"SET hostgroup_id = %d, hostname = '%s', port = %d, status = '%s', weight = %d, max_connections = %d, max_replication_lag = %d "+
			"WHERE hostgroup_id = %d AND hostname = '%s';",
		s.HostgroupID,
		s.Hostname,
		s.Port,
		s.Status,
		s.Weight,
		s.MaxConnections,
		s.MaxReplicationLag,
		s.HostgroupID,
		s.Hostname,
	)
}

func (s *Server) QueryDelete() string {
	return fmt.Sprintf(
		"DELETE FROM mysql_servers WHERE hostgroup_id = %d AND hostname = '%s';",
		s.HostgroupID,
		s.Hostname,
	)
}
