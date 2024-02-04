package proxysql_test

import (
	"testing"

	"github.com/debeando/go-proxysql"

	"github.com/stretchr/testify/assert"
)

func TestConnectionDSN(t *testing.T) {
	c := proxysql.Connection{
		Host:     "127.0.0.1",
		Port:     6032,
		Username: "radmin",
		Password: "radmin",
	}

	assert.Equal(t, c.DSN(), "radmin:radmin@tcp(127.0.0.1:6032)/")
}

func TestConnectionDSNSecret(t *testing.T) {
	c := proxysql.Connection{
		Host:     "127.0.0.1",
		Port:     6032,
		Username: "radmin",
		Password: "radmin",
	}

	assert.Equal(t, c.DSNSecret(), "radmin:***@tcp(127.0.0.1:6032)/")
}

func TestConnectionConnect(t *testing.T) {
	c := proxysql.Connection{
		Host:     "127.0.0.1",
		Port:     6032,
		Username: "radmin",
		Password: "radmin",
	}

	assert.Nil(t, c.Connect())
}
