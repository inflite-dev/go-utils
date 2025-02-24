package postgres

import (
	"regexp"

	"github.com/lib/pq"
)

const ErrCodeUniqueViolation = "23505"

var pqErrorDetailRegex = regexp.MustCompile(`Key\s\((?P<Key>[a-zA-Z0-9_]*)\)=\((?P<Value>.*)\).*`)

func GetErrCodeUniqueViolationKV(err *pq.Error) (string, string) {
	matches := pqErrorDetailRegex.FindStringSubmatch(err.Detail)
	return matches[1], matches[2]
}
