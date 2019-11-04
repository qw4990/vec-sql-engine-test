package main

import (
	"testing"

	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/tidb/types"
	"github.com/pingcap/tidb/util/chunk"
	"github.com/qw4990/vec-sql-engine-test/util"
)

func BenchmarkVec(b *testing.B) {
	ctx := util.MockCtx()
	expr := util.BuildExpr(ctx)
	data := util.GenData()
	buf := chunk.NewColumn(types.NewFieldType(mysql.TypeDouble), 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := expr.VecEvalReal(ctx, data, buf); err != nil {
			b.Fatal(err)
		}
	}
}
