package proxysql_test

import (
	"testing"

	"github.com/debeando/go-proxysql"

	"github.com/stretchr/testify/assert"
)

func TestServersReset(t *testing.T) {
	servers := proxysql.Servers{}
	servers.New(&proxysql.Connection{})
	servers.Add(proxysql.Server{
		HostgroupID:       11,
		Hostname:          "127.0.0.1",
		MaxConnections:    100,
		MaxReplicationLag: 60,
		Port:              3307,
		Status:            proxysql.ONLINE,
		Weight:            100,
	})

	assert.Equal(t, servers.Count(), 1)
	servers.Reset()
	assert.Equal(t, servers.Count(), 0)
}

func TestServersAdd(t *testing.T) {
	servers := proxysql.Servers{}
	servers.New(&proxysql.Connection{})
	servers.Add(proxysql.Server{
		HostgroupID:       11,
		Hostname:          "127.0.0.1",
		MaxConnections:    100,
		MaxReplicationLag: 60,
		Port:              3307,
		Status:            proxysql.ONLINE,
		Weight:            100,
	})

	assert.Equal(t, servers.Count(), 1)
}

func TestServersCount(t *testing.T) {
	servers := proxysql.Servers{}
	assert.Equal(t, servers.Count(), 0)
}

func TestServersFirst(t *testing.T) {
	servers := proxysql.Servers{}
	servers.New(&proxysql.Connection{})
	servers.Reset()

	t.Run("AddSeveral", func(t *testing.T) {
		servers.Add(proxysql.Server{
			HostgroupID: 10,
			Hostname:    "127.0.0.1",
			Port:        3307,
			Weight:      123,
			Status:      proxysql.ONLINE,
		})
		servers.Add(proxysql.Server{
			HostgroupID: 11,
			Hostname:    "127.0.0.1",
			Port:        3307,
			Weight:      100,
			Status:      proxysql.OFFLINE_SOFT,
		})

		assert.Equal(t, servers.Count(), 2)
	})

	t.Run("GetFirst", func(t *testing.T) {
		assert.Equal(t, servers.First(), servers.Get(
			proxysql.Server{
				HostgroupID: 10,
				Hostname:    "127.0.0.1",
				Port:        3307,
			}),
		)
	})
}

func TestServersGet(t *testing.T) {
	servers := proxysql.Servers{}
	servers.New(&proxysql.Connection{})
	servers.Reset()

	t.Run("AddSeveral", func(t *testing.T) {
		servers.Add(proxysql.Server{
			HostgroupID: 10,
			Hostname:    "127.0.0.1",
			Port:        3307,
			Weight:      123,
			Status:      proxysql.ONLINE,
		})
		servers.Add(proxysql.Server{
			HostgroupID: 11,
			Hostname:    "127.0.0.1",
			Port:        3307,
			Weight:      100,
			Status:      proxysql.OFFLINE_SOFT,
		})

		assert.Equal(t, servers.Count(), 2)
	})

	t.Run("GetFirst", func(t *testing.T) {
		server := servers.Get(proxysql.Server{
			HostgroupID: 10,
			Hostname:    "127.0.0.1",
			Port:        3307,
		})

		assert.Equal(t, server.Weight, 123)
		assert.Equal(t, server.Status, proxysql.ONLINE)
	})

	t.Run("GetSecond", func(t *testing.T) {
		server := servers.Get(proxysql.Server{
			HostgroupID: 11,
			Hostname:    "127.0.0.1",
			Port:        3307,
		})

		assert.Equal(t, server.Weight, 100)
		assert.Equal(t, server.Status, proxysql.OFFLINE_SOFT)
	})

	t.Run("ModifyFirst", func(t *testing.T) {
		server := servers.Get(proxysql.Server{
			HostgroupID: 10,
			Hostname:    "127.0.0.1",
			Port:        3307,
		})

		server.Weight = 1
		server.Status = proxysql.OFFLINE_SOFT
	})

	t.Run("IsModified", func(t *testing.T) {
		server := servers.Get(proxysql.Server{
			HostgroupID: 10,
			Hostname:    "127.0.0.1",
			Port:        3307,
		})

		assert.Equal(t, server.Weight, 1)
		assert.Equal(t, server.Status, proxysql.OFFLINE_SOFT)
	})

	t.Run("IsEmpty", func(t *testing.T) {
		server := servers.Get(proxysql.Server{
			HostgroupID: 20,
			Hostname:    "127.0.0.1",
			Port:        3306,
		})

		assert.Empty(t, server, proxysql.Server{})
	})
}

func TestServersLoad(t *testing.T) {

}

func TestServersSave(t *testing.T) {

}
