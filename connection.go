package proxysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Instance *sql.DB
	Host     string `json:"host" yaml:"host"`
	Password string `json:"password" yaml:"password"`
	Port     uint16 `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
}

func (c *Connection) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
	)
}

func (c *Connection) DSNSecret() string {
	return fmt.Sprintf(
		"%s:***@tcp(%s:%d)/",
		c.Username,
		c.Host,
		c.Port,
	)
}

func (c *Connection) Connect() (err error) {
	c.Instance, err = sql.Open("mysql", c.DSN())
	if err != nil {
		return err
	}

	return c.Instance.Ping()
}
