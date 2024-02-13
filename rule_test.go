package proxysql_test

import (
	"testing"

	"github.com/debeando/go-proxysql"

	"github.com/stretchr/testify/assert"
)

func TestRuleEqual(t *testing.T) {
	rule := proxysql.Rule{ID: 100}

	assert.True(t, rule.Equal(proxysql.Rule{ID: 100}))
}

func TestRuleCRUD(t *testing.T) {
	con := proxysql.Connection{
		Host:     "127.0.0.1",
		Port:     6032,
		Username: "radmin",
		Password: "radmin",
	}
	con.Connect()

	rule := proxysql.Rule{
		ID:          100,
		Active:      1,
		Apply:       1,
		HostgroupID: 11,
		MatchDigest: `^SELECT.*WHERE.* IN \(.*$`,
		Username:    "foo",
	}
	rule.New(&con)

	t.Run("SelectAndIsEmpty", func(t *testing.T) {
		assert.Error(t, rule.Select())
	})

	t.Run("Insert", func(t *testing.T) {
		assert.NoError(t, rule.Insert())
	})

	t.Run("SelectAndIsNotEmpty", func(t *testing.T) {
		assert.Nil(t, rule.Select())
		assert.Equal(t, rule.Apply, 1)
		assert.Equal(t, rule.Active, 1)
		assert.Equal(t, rule.HostgroupID, 11)
		assert.Equal(t, rule.MatchDigest, `^SELECT.*WHERE.* IN \(.*$`)
		assert.Equal(t, rule.Username, "foo")
	})

	t.Run("Update", func(t *testing.T) {
		rule.Active = 0
		rule.Apply = 0
		assert.Nil(t, rule.Update())
	})

	t.Run("SelectAndVerify", func(t *testing.T) {
		rule = proxysql.Rule{
			ID:          100,
			Active:      1,
			Apply:       1,
			HostgroupID: 11,
			MatchDigest: `^SELECT.*WHERE.* IN \(.*$`,
			Username:    "foo",
		}
		rule.New(&con)

		assert.Nil(t, rule.Select())
		assert.Equal(t, rule.Active, 0)
		assert.Equal(t, rule.Apply, 0)
		assert.Equal(t, rule.HostgroupID, 11)
		assert.Equal(t, rule.MatchDigest, `^SELECT.*WHERE.* IN \(.*$`)
		assert.Equal(t, rule.Username, "foo")
	})

	t.Run("Delete", func(t *testing.T) {
		assert.NoError(t, rule.Delete())
	})
}
