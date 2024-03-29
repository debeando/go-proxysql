package proxysql_test

import (
	"testing"

	"github.com/debeando/go-proxysql"

	"github.com/stretchr/testify/assert"
)

func TestRulesReset(t *testing.T) {
	rules := proxysql.Rules{}
	rules.New(&proxysql.Connection{})
	rules.Add(proxysql.Rule{ID: 101})

	assert.Equal(t, rules.Count(), 1)
	rules.Reset()
	assert.Equal(t, rules.Count(), 0)
}

func TestRulesAdd(t *testing.T) {
	rules := proxysql.Rules{}
	rules.New(&proxysql.Connection{})
	rules.Add(proxysql.Rule{ID: 101})
	rules.Add(proxysql.Rule{ID: 102})

	assert.Equal(t, rules.Count(), 2)
}

func TestRulesCount(t *testing.T) {
	rules := proxysql.Rules{}
	rules.New(&proxysql.Connection{})
	rules.Add(proxysql.Rule{ID: 101})
	rules.Add(proxysql.Rule{ID: 102})
	rules.Add(proxysql.Rule{ID: 103})

	assert.Equal(t, rules.Count(), 3)
}

func TestRulesGet(t *testing.T) {
	rules := proxysql.Rules{}
	rules.New(&proxysql.Connection{})

	t.Run("AddSeveral", func(t *testing.T) {
		rules.Add(proxysql.Rule{ID: 101})
		rules.Add(proxysql.Rule{ID: 102})

		assert.Equal(t, rules.Count(), 2)
	})

	t.Run("GetFirst", func(t *testing.T) {
		rule := rules.Get(proxysql.Rule{ID: 101})

		assert.Equal(t, rule.ID, 101)
	})

	t.Run("GetSecond", func(t *testing.T) {
		rule := rules.Get(proxysql.Rule{ID: 102})

		assert.Equal(t, rule.ID, 102)
	})

}

func TestRulesDisable(t *testing.T) {
	rules := proxysql.Rules{}
	rules.New(&proxysql.Connection{})

	rules.Add(proxysql.Rule{ID: 101, Active: 1})
	rules.Add(proxysql.Rule{ID: 102, Active: 1})
	rules.Add(proxysql.Rule{ID: 103, Active: 1})

	assert.Equal(t, rules.Count(), 3)
	assert.Equal(t, rules.Get(proxysql.Rule{ID: 101}).Active, 1)
	assert.Equal(t, rules.Get(proxysql.Rule{ID: 102}).Active, 1)
	assert.Equal(t, rules.Get(proxysql.Rule{ID: 103}).Active, 1)

	rules.Disable()

	assert.Equal(t, rules.Count(), 3)
	assert.Equal(t, rules.Get(proxysql.Rule{ID: 101}).Active, 0)
	assert.Equal(t, rules.Get(proxysql.Rule{ID: 102}).Active, 0)
	assert.Equal(t, rules.Get(proxysql.Rule{ID: 103}).Active, 0)
}

func TestRulesEnable(t *testing.T) {
	rules := proxysql.Rules{}
	rules.New(&proxysql.Connection{})

	rules.Add(proxysql.Rule{ID: 101})
	rules.Add(proxysql.Rule{ID: 102})
	rules.Add(proxysql.Rule{ID: 103})

	assert.Equal(t, rules.Count(), 3)
	assert.Equal(t, rules.Get(proxysql.Rule{ID: 101}).Active, 0)
	assert.Equal(t, rules.Get(proxysql.Rule{ID: 102}).Active, 0)
	assert.Equal(t, rules.Get(proxysql.Rule{ID: 103}).Active, 0)

	rules.Enable()

	assert.Equal(t, rules.Count(), 3)
	assert.Equal(t, rules.Get(proxysql.Rule{ID: 101}).Active, 1)
	assert.Equal(t, rules.Get(proxysql.Rule{ID: 102}).Active, 1)
	assert.Equal(t, rules.Get(proxysql.Rule{ID: 103}).Active, 1)
}

func TestRulesLoad(t *testing.T) {
	con := proxysql.Connection{
		Host:     "127.0.0.1",
		Port:     6032,
		Username: "radmin",
		Password: "radmin",
	}
	con.Connect()

	assert.Nil(t, con.Connect())

	rules := proxysql.Rules{}
	rules.New(&con)
	rules.Add(proxysql.Rule{
		ID:          100,
		Active:      1,
		Apply:       1,
		HostgroupID: 11,
		MatchDigest: `^SELECT.*WHERE.* IN \(.*$`,
		Username:    "foo",
	})
	rules.Insert()
	rules.Reset()
	rules.Load()

	rule := rules.Get(proxysql.Rule{ID: 100})

	assert.Equal(t, rule.Active, 1)
	assert.Equal(t, rule.Apply, 1)
	assert.Equal(t, rule.HostgroupID, 11)
	assert.Equal(t, rule.MatchDigest, `^SELECT.*WHERE.* IN \(.*$`)
	assert.Equal(t, rule.Username, "foo")
	assert.NoError(t, rule.Delete())
}
