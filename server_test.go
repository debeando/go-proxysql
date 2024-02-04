package proxysql_test

import (
	"testing"

	"github.com/debeando/go-proxysql"

	"github.com/stretchr/testify/assert"
)

func TestServerEqual(t *testing.T) {
	srv := proxysql.Server{
		HostgroupID: 10,
		Hostname:    "127.0.0.1",
		Port:        3306,
	}

	assert.True(t, srv.Equal(proxysql.Server{
		HostgroupID: 10,
		Hostname:    "127.0.0.1",
		Port:        3306,
	}))

	assert.False(t, srv.Equal(proxysql.Server{
		HostgroupID: 11,
		Hostname:    "127.0.0.1",
		Port:        3306,
	}))
}

func TestServerCRUD(t *testing.T) {
	con := proxysql.Connection{
		Host:     "127.0.0.1",
		Port:     6032,
		Username: "radmin",
		Password: "radmin",
	}
	con.Connect()

	srv := proxysql.Server{
		HostgroupID: 10,
		Hostname:    "127.0.0.1",
		Port:        3306,
	}
	srv.New(&con)

	t.Run("SelectAndIsEmpty", func(t *testing.T) {
		assert.Error(t, srv.Select())
	})

	t.Run("Insert", func(t *testing.T) {
		assert.NoError(t, srv.Insert())
	})

	t.Run("SelectAndIsNotEmpty", func(t *testing.T) {
		assert.Nil(t, srv.Select())
	})

	t.Run("Update", func(t *testing.T) {
		srv.Status = proxysql.OFFLINE_SOFT
		assert.Nil(t, srv.Update())
	})

	t.Run("SelectAndVerify", func(t *testing.T) {
		srv = proxysql.Server{
			HostgroupID: 10,
			Hostname:    "127.0.0.1",
			Port:        3306,
		}
		srv.New(&con)

		assert.Nil(t, srv.Select())
		assert.Equal(t, srv.Status, proxysql.OFFLINE_SOFT)
	})

	t.Run("Delete", func(t *testing.T) {
		err := srv.Delete()
		assert.Nil(t, err)
	})

}
