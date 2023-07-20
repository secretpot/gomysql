package statement

import (
	"fmt"
	"strings"
)

type Selector struct {
	URD
	table   string
	columns []string
}

func Select(columns ...string) Selector {
	return Selector{
		columns: columns,
	}
}

func (s Selector) From(table string) Selector {
	s.table = table
	return s
}

// Implements CURDInterface::Where
func (s Selector) Where(conditions ...string) Selector {
	s.where = append(s.where, Express(conditions...))
	return s
}

// Implements CURDInterface::And
func (s Selector) And(conditions ...string) Selector {
	s.where = append(s.where, "AND")
	return s.Where(conditions...)
}

// Implements CURDInterface::Or
func (s Selector) Or(conditions ...string) Selector {
	s.where = append(s.where, "OR")
	return s.Where(conditions...)
}

func (s Selector) String() string {
	wanna := strings.Join(s.columns, ",")
	conds := strings.Join(s.where, " ")
	if len(strings.Fields(conds)) > 0 {
		conds = fmt.Sprintf("WHERE %v", conds)
	}
	return fmt.Sprintf("SELECT %v FROM %v %v", wanna, s.table, conds)
}
