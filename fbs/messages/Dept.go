// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package messages

import "strconv"

type Dept int8

const (
	DeptD1 Dept = 0
	DeptD2 Dept = 1
	DeptD3 Dept = 2
)

var EnumNamesDept = map[Dept]string{
	DeptD1: "D1",
	DeptD2: "D2",
	DeptD3: "D3",
}

var EnumValuesDept = map[string]Dept{
	"D1": DeptD1,
	"D2": DeptD2,
	"D3": DeptD3,
}

func (v Dept) String() string {
	if s, ok := EnumNamesDept[v]; ok {
		return s
	}
	return "Dept(" + strconv.FormatInt(int64(v), 10) + ")"
}
