package proxysql

const (
	ServerStatus string = ""
	ONLINE              = "ONLINE"
	OFFLINE_SOFT        = "OFFLINE_SOFT"
	OFFLINE_HARD        = "OFFLINE_HARD"
	SHUNNED             = "SHUNNED"
)

var instance *ProxySQL

type ProxySQL struct {
	connection *Connection
	servers    Servers
	rules      Rules
}

func Instance() *ProxySQL {
	if instance == nil {
		instance = &ProxySQL{}
	}
	return instance
}

func (p *ProxySQL) Connect(c *Connection) error {
	if err := c.Connect(); err != nil {
		return err
	}

	p.connection = c
	p.servers.New(p.connection)
	p.rules.New(p.connection)
	return nil
}

func (p *ProxySQL) Servers() *Servers {
	return &(*p).servers
}

func (p *ProxySQL) Rules() *Rules {
	return &(*p).rules
}
