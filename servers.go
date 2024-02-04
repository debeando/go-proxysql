package proxysql

import (
	"errors"
)

type Servers struct {
	connection *Connection
	items      *[]Server
}

func (s *Servers) New(c *Connection) {
	s.connection = c
	s.items = &[]Server{}
}

func (s *Servers) Reset() {
	*s.items = (*s.items)[:0]
}

func (s *Servers) Add(in Server) {
	in.New(s.connection)

	*s.items = append(*s.items, in)
}

func (s *Servers) Count() int {
	if s.items == nil {
		return 0
	}

	return len(*s.items)
}

func (s *Servers) First() *Server {
	if s.Count() > 0 {
		return &(*s.items)[0]
	}
	return &Server{}
}

func (s *Servers) Get(in Server) *Server {
	for index, server := range *s.items {
		if server.Equal(in) {
			return &(*s.items)[index]
		}
	}
	return &Server{}
}

func (s *Servers) Load() error {
	if s.connection == nil {
		return errors.New("Connection does not exist")
	}

	results, err := s.connection.Instance.Query(s.QuerySelect())
	if err != nil {
		return err
	}

	s.Reset()

	for results.Next() {
		server := Server{}

		err = results.Scan(
			&server.HostgroupID,
			&server.Hostname,
			&server.Port,
			&server.Status,
			&server.Weight,
			&server.MaxConnections,
			&server.MaxReplicationLag,
		)
		if err != nil {
			return err
		}

		s.Add(server)
	}

	return nil
}

func (s *Servers) Save() {
	for _, server := range *s.items {
		if err := server.Select(); err != nil {
			server.Insert()
		} else {
			server.Update()
		}
	}
}

func (s *Servers) LoadToRunTime() error {
	if s.connection == nil {
		return errors.New("Connection does not exist")
	}

	s.connection.Instance.Query("LOAD MYSQL SERVERS TO RUNTIME;")

	return nil
}

func (s *Servers) SaveToDisk() error {
	if s.connection == nil {
		return errors.New("Connection does not exist")
	}

	s.connection.Instance.Query("SAVE MYSQL SERVERS TO DISK;")

	return nil
}

func (s *Servers) QuerySelect() string {
	return "SELECT hostgroup_id, hostname, port, status, weight, max_connections, max_replication_lag FROM mysql_servers;"
}
