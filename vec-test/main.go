package main

import (
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/tidb/types"
	"github.com/pingcap/tidb/util/chunk"
	"github.com/qw4990/vec-sql-engine-test/util"
)

func main() {
	ctx := util.MockCtx()
	expr := util.BuildExpr(ctx)
	data := util.GenData()
	buf := chunk.NewColumn(types.NewFieldType(mysql.TypeDouble), 1024)
	for i := 0; i < 10000; i++ {
		if err := expr.VecEvalReal(ctx, data, buf); err != nil {
			panic(err)
		}
	}
}
