package util

import (
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/tidb/expression"
	"github.com/pingcap/tidb/sessionctx"
	"github.com/pingcap/tidb/types"
	"github.com/pingcap/tidb/util/chunk"
	"github.com/pingcap/tidb/util/mock"
)

// col0 * 0.8 + col1
func BuildExpr(ctx sessionctx.Context) expression.Expression {
	tpDouble := types.NewFieldType(mysql.TypeDouble)
	constant := &expression.Constant{
		RetType: tpDouble,
		Value:   types.NewDatum(0.8),
	}
	col0 := &expression.Column{
		RetType: tpDouble,
		Index:   0,
	}
	col1 := &expression.Column{
		RetType: tpDouble,
		Index:   1,
	}
	mul, err := expression.NewFunction(ctx, ast.Mul, tpDouble, col0, constant)
	if err != nil {
		panic(err)
	}
	plus, err := expression.NewFunction(ctx, ast.Plus, tpDouble, mul, col1)
	if err != nil {
		panic(err)
	}
	return plus
}

func GenData() *chunk.Chunk {
	tpDouble := types.NewFieldType(mysql.TypeDouble)
	chk := chunk.New([]*types.FieldType{tpDouble, tpDouble}, 1024, 1024)
	for i := 0; i < 1024; i ++ {
		chk.AppendFloat64(0, float64(i))
		chk.AppendFloat64(1, float64(i))
	}
	return chk
}

func RowBasedEval(ctx sessionctx.Context, expr expression.Expression, data *chunk.Chunk) {
	it := chunk.NewIterator4Chunk(data)
	for i := 0; i < 10000; i++ {
		for row := it.Begin(); row != it.End(); row = it.Next() {
			if _, err := expr.Eval(row); err != nil {
				panic(err)
			}
		}
	}
}

func VectorizedEval(ctx sessionctx.Context, expr expression.Expression, data *chunk.Chunk) {
	buf := chunk.NewColumn(types.NewFieldType(mysql.TypeDouble), 1024)
	for i := 0; i < 10000; i++ {
		if err := expr.VecEvalReal(ctx, data, buf); err != nil {
			panic(err)
		}
	}
}

func MockCtx() sessionctx.Context {
	return mock.NewContext()
}
