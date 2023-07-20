package statement

import (
	"fmt"
	"strings"
)

type Inserter struct {
	table   string
	columns []string
	values  []Recordable
}

func Insert() Inserter {
	return Inserter{}
}
func (i Inserter) Into(table string, columns []string) Inserter {
	i.table = table
	i.columns = columns
	return i
}

func (i Inserter) Values(values []Recordable) Inserter {
	i.values = append(i.values, values...)
	return i
}

func (i Inserter) String() string {
	if len(i.values) <= 0 {
		return ""
	}
	values := []string{}
	for _, val := range i.values {
		values = append(values, fmt.Sprintf("(%v)", val.Recordizate(i.columns)))
	}
	return fmt.Sprintf("INSERT INTO %v(%v) VALUES %v", i.table, strings.Join(i.columns, ","), strings.Join(values, ","))
}
