package statement

import (
	"strings"
)

func Express(parts ...string) string {
	return strings.Join(parts, "")
}

// Upate & Read & Delete
type URDInterface interface {
	Where(conditions ...string) URDInterface
	And(conditions ...string) URDInterface
	Or(conditions ...string) URDInterface
}
type URD struct {
	where []string
	URDInterface
}

// Insertable object
type Recordable interface {
	Recordizate(columns []string) string
}
