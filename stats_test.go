package proxysql_test

import (
	"testing"

	"github.com/debeando/go-proxysql"

	"github.com/stretchr/testify/assert"
)

func TestStatsSelect(t *testing.T) {
	con := proxysql.Connection{
		Host:     "127.0.0.1",
		Port:     6032,
		Username: "radmin",
		Password: "radmin",
	}
	con.Connect()

	srvs := proxysql.Servers{}
	srvs.New(&con)

	srv := proxysql.Server{
		HostgroupID: 10,
		Hostname:    "127.0.0.1",
		Port:        3306,
	}
	srv.New(&con)
	srv.Insert()

	srvs.LoadToRunTime()

	assert.Equal(t, srv.Stats().Select().Status, proxysql.ONLINE)

	srv.Delete()
	srvs.LoadToRunTime()
}
