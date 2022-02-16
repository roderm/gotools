package squirrel

import "github.com/Masterminds/squirrel"

type EqCall map[string]func() interface{}

func (eqc EqCall) ToSql() (sql string, args []interface{}, err error) {
	eq := squirrel.Eq{}
	for k, fn := range eqc {
		eq[k] = fn()
	}
	return eq.ToSql()
}
