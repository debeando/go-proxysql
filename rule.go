package proxysql

import (
	"errors"
	"fmt"
)

type Rule struct {
	connection  *Connection
	Active      int    `yaml:"active"`
	Apply       int    `yaml:"apply"`
	HostgroupID int    `yaml:"destination_hostgroup"`
	ID          int    `yaml:"rule_id"`
	MatchDigest string `yaml:"match_digest"`
	Username    string `yaml:"username"`
}

func (r *Rule) New(c *Connection) {
	r.connection = c
}

func (r *Rule) Equal(in Rule) bool {
	return r.ID == in.ID
}

func (r *Rule) Select() error {
	if r.connection == nil {
		return errors.New("Connection does not exist")
	}

	return r.connection.Instance.QueryRow(r.QuerySelect()).Scan(
		&r.ID,
		&r.Active,
		&r.Apply,
		&r.HostgroupID,
		&r.MatchDigest,
		&r.Username)
}

func (r *Rule) Insert() error {
	if r.connection == nil {
		return errors.New("Connection does not exist")
	}

	_, err := r.connection.Instance.Query(r.QueryInsert())
	return err
}

func (r *Rule) Update() error {
	if r.connection == nil {
		return errors.New("Connection does not exist")
	}

	_, err := r.connection.Instance.Query(r.QueryUpdate())
	return err
}

func (r *Rule) Delete() error {
	if r.connection == nil {
		return errors.New("Connection does not exist")
	}

	_, err := r.connection.Instance.Query(r.QueryDelete())
	return err
}

func (r *Rule) QuerySelect() string {
	return fmt.Sprintf(
		"SELECT rule_id, active, apply, destination_hostgroup, match_digest, username FROM mysql_query_rules WHERE rule_id = %d;",
		r.ID,
	)
}

func (r *Rule) QueryInsert() string {
	return fmt.Sprintf(
		"INSERT INTO mysql_query_rules (rule_id, active, apply, destination_hostgroup, username, match_digest) VALUES (%d, %d, %d, %d, '%s', '%s');",
		r.ID,
		r.Active,
		r.Apply,
		r.HostgroupID,
		r.Username,
		r.MatchDigest,
	)
}

func (r *Rule) QueryUpdate() string {
	return fmt.Sprintf(
		"UPDATE mysql_query_rules "+
			"SET active = %d, apply = %d, destination_hostgroup = %d, username = '%s', match_digest = '%s' "+
			"WHERE rule_id = %d;",
		r.Active,
		r.Apply,
		r.HostgroupID,
		r.Username,
		r.MatchDigest,
		r.ID,
	)
}

func (r *Rule) QueryDelete() string {
	return fmt.Sprintf(
		"DELETE FROM mysql_query_rules WHERE rule_id = %d;",
		r.ID,
	)
}
