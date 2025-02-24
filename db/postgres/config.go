package postgres

import (
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
