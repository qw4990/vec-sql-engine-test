package main

type Row interface {
	GetReal(colIdx int) (float64, bool)
}

type Node interface {
	evalReal(row Row) (val float64, isNull bool)
}

type plusRealNode struct {
	leftChild  Node
	rightChild Node
}

func (node *plusRealNode) evalReal(row Row) (float64, bool) {
	v1, null1 := node.leftChild.evalReal(row)
	v2, null2 := node.rightChild.evalReal(row)
	return v1 + v2, null1 || null2
}

type multiplyRealNode struct {
	leftChild  Node
	rightChild Node
}

func (node *multiplyRealNode) evalReal(row Row) (float64, bool) {
	v1, null1 := node.leftChild.evalReal(row)
	v2, null2 := node.rightChild.evalReal(row)
	return v1 * v2, null1 || null2
}

type constantNode struct {
	someConstantValue float64
	isNull            bool
}

func (node *constantNode) evalReal(row Row) (float64, bool) {
	return node.someConstantValue, node.isNull // 0.8 in this case
}

type columnNode struct {
	colIdx int
}

func (node *columnNode) evalReal(row Row) (float64, bool) {
	return row.GetReal(node.colIdx)
}

func main() {

}
