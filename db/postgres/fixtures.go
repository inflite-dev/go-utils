package postgres

import (
	"fmt"
	"math/rand"
)

// RandomDBName generates a Postgres database name with a provided prefix and random suffix
// in order to avoid collisions and race conditions in parallel tests.
//
// This uses a pseudorandom integer rather than a UUID to avoid hyphens in the database name.
// Postgres database names with hyphens do not autocomplete on the psql command line,
// which is a pain in case you get stuck manually deleting a bunch of these.
func RandomDBName(prefix string) string {
	return fmt.Sprintf("%s_%d", prefix, rand.Int())
}
