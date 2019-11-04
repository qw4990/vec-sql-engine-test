package main

import (
	"testing"

	"github.com/pingcap/tidb/util/chunk"
	"github.com/qw4990/vec-sql-engine-test/util"
)

func BenchmarkRow(b *testing.B) {
	ctx := util.MockCtx()
	expr := util.BuildExpr(ctx)
	data := util.GenData()
	it := chunk.NewIterator4Chunk(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for row := it.Begin(); row != it.End(); row = it.Next() {
			if _, _, err := expr.EvalReal(ctx, row); err != nil {
				panic(err)
			}
		}
	}
}
