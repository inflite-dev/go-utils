package postgres

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomDBName generates Postgres database names with a common prefix
// in order to avoid collisions and race conditions in parallel tests
func RandomDBName(dbNamePrefix string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf(
		"%s_%d",
		dbNamePrefix,
		rand.Int(),
	)
}
