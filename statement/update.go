package statement

import (
	"fmt"
	"strings"
)

type Updater struct {
	URD
	table   string
	mapping map[string]interface{}
}

func Update(table string) Updater {
	return Updater{
		table:   table,
		mapping: map[string]interface{}{},
	}
}

func (u Updater) Set(mapping map[string]interface{}) Updater {
	for k, v := range mapping {
		u.mapping[k] = v
	}
	return u
}

// Implements CURDInterface::Where
func (u Updater) Where(conditions ...string) Updater {
	u.where = append(u.where, Express(conditions...))
	return u
}

// Implements CURDInterface::And
func (u Updater) And(conditions ...string) Updater {
	u.where = append(u.where, "AND")
	return u.Where(conditions...)
}

// Implements CURDInterface::Or
func (u Updater) Or(conditions ...string) Updater {
	u.where = append(u.where, "OR")
	return u.Where(conditions...)
}
func (u Updater) String() string {
	exps := []string{}
	for k, v := range u.mapping {
		exps = append(exps, fmt.Sprintf("%v=%v", k, v))
	}
	expect := strings.Join(exps, ",")
	conds := strings.Join(u.where, " ")
	if len(strings.Fields(conds)) > 0 {
		conds = fmt.Sprintf("WHERE %v", conds)
	}
	return fmt.Sprintf("UPDATE %v SET %v %v", u.table, expect, conds)
}
