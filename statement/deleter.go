package statement

import (
	"fmt"
	"strings"
)

type Deleter struct {
	URD
	table string
}

func Delete() Deleter {
	return Deleter{}
}

func (d Deleter) From(table string) Deleter {
	d.table = table
	return d
}

// Implements CURDInterface::Where
func (d Deleter) Where(conditions ...string) Deleter {
	d.where = append(d.where, Express(conditions...))
	return d
}

// Implements CURDInterface::And
func (d Deleter) And(conditions ...string) Deleter {
	d.where = append(d.where, "AND")
	return d.Where(conditions...)
}

// Implements CURDInterface::Or
func (d Deleter) Or(conditions ...string) Deleter {
	d.where = append(d.where, "OR")
	return d.Where(conditions...)
}

func (d Deleter) String() string {
	conds := strings.Join(d.where, " ")
	if len(strings.Fields(conds)) > 0 {
		conds = fmt.Sprintf("WHERE %v", conds)
	}
	return fmt.Sprintf("DELETE FROM %v %v", d.table, conds)
}
