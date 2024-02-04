package proxysql_test

import (
	"testing"

	"github.com/debeando/go-proxysql"

	"github.com/stretchr/testify/assert"
)

func TestProxySQLNew(t *testing.T) {
	proxy := proxysql.Instance()

	assert.NotNil(t, proxy)
}

func TestProxySQLConnect(t *testing.T) {
	proxy := proxysql.Instance()
	err := proxy.Connect(&proxysql.Connection{
		Host:     "127.0.0.1",
		Port:     6032,
		Username: "radmin",
		Password: "radmin",
	})

	assert.Nil(t, err)
}

func TestProxySQLServers(t *testing.T) {
	proxy := proxysql.Instance()
	servers := proxy.Servers()

	assert.NotNil(t, servers)
}
