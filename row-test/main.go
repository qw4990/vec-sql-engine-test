package main

import (
	"github.com/pingcap/tidb/util/chunk"
	"github.com/qw4990/vec-sql-engine-test/util"
)

func main() {
	ctx := util.MockCtx()
	expr := util.BuildExpr(ctx)
	data := util.GenData()
	it := chunk.NewIterator4Chunk(data)
	for i := 0; i < 10000; i++ {
		for row := it.Begin(); row != it.End(); row = it.Next() {
			if _, _, err := expr.EvalReal(ctx, row); err != nil {
				panic(err)
			}
		}
	}
}
