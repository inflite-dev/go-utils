package postgres

import (
	"flag"
	"fmt"
)

// Config defines Postgres sql connection parameters
type Config struct {
	Host                  string
	Port                  int
	Username              string
	Password              string
	Database              string
	ApplicationName       string
	ConnectTimeoutSeconds int
	SSLMode               string
}

func (c *Config) RegisterFlags(f *flag.FlagSet) {
	f.StringVar(&c.Host, "postgres.host", "localhost", "hostname to connect to")
	f.IntVar(&c.Port, "postgres.port", 5432, "port to connect to")
	f.StringVar(&c.Username, "postgres.username", "", "username to connect with")
	f.StringVar(&c.Password, "postgres.password", "", "password to connect with")
	f.StringVar(&c.Database, "postgres.database", "", "database to connect to")
	f.StringVar(&c.ApplicationName, "postgres.application-name", "", "application name to connect with")
	f.IntVar(&c.ConnectTimeoutSeconds, "postgres.connect-timeout-seconds", 5, "connection timeout in seconds")
	f.StringVar(&c.SSLMode, "postgres.sslmode", "disable", "sslmode")
}

// BuildConnectionURI builds a connection string for lib/pq from Config.
// If a missing or invalid field is provided, an error is returned.
func BuildConnectionURI(cc Config) string {
	auth := ""
	if cc.Username != "" || cc.Password != "" {
		auth = fmt.Sprintf("%s:%s@", cc.Username, cc.Password)
	}
	url := fmt.Sprintf(
		"postgres://%s%s:%d/%s?application_name=%s&connect_timeout=%d&sslmode=%s",
		auth,
		cc.Host,
		cc.Port,
		cc.Database,
		cc.ApplicationName,
		cc.ConnectTimeoutSeconds,
		cc.SSLMode,
	)
	return url
}
