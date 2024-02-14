package proxysql

type Rules struct {
	connection *Connection
	items      *[]Rule
}

func (r *Rules) New(c *Connection) {
	r.connection = c
	r.items = &[]Rule{}
}

func (r *Rules) Reset() {
	*r.items = (*r.items)[:0]
}

func (r *Rules) Add(in Rule) {
	in.New(r.connection)

	*r.items = append(*r.items, in)
}

func (r *Rules) Count() int {
	return len(*r.items)
}

func (r *Rules) Get(in Rule) *Rule {
	for index, rule := range *r.items {
		if rule.Equal(in) {
			return &(*r.items)[index]
		}
	}
	return &Rule{}
}

func (r *Rules) Disable() {
	for index, _ := range *r.items {
		(*r.items)[index].Active = 0
	}
}

func (r *Rules) Enable() {
	for index, _ := range *r.items {
		(*r.items)[index].Active = 1
	}
}

func (r *Rules) Load() error {
	results, err := r.connection.Instance.Query(r.QuerySelect())
	if err != nil {
		return err
	}

	r.Reset()

	for results.Next() {
		rule := Rule{}

		err = results.Scan(
			&rule.ID,
			&rule.Active,
			&rule.Apply,
			&rule.HostgroupID,
			&rule.Username,
			&rule.MatchDigest,
		)
		if err != nil {
			return err
		}

		r.Add(rule)
	}

	return nil
}

func (r *Rules) Insert() {
	for index, _ := range *r.items {
		(*r.items)[index].Insert()
	}
}

func (r *Rules) Update() {
	for index, _ := range *r.items {
		(*r.items)[index].Update()
	}
}

func (r *Rules) LoadToRunTime() {
	r.connection.Instance.Query("LOAD MYSQL QUERY RULES TO RUNTIME;")
}

func (r *Rules) SaveToDisk() {
	r.connection.Instance.Query("SAVE MYSQL QUERY RULES TO DISK;")
}

func (r *Rules) QuerySelect() string {
	return "SELECT rule_id, active, apply, destination_hostgroup, username, match_digest FROM mysql_query_rules;"
}
