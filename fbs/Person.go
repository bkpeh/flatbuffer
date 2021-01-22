// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Person struct {
	_tab flatbuffers.Table
}

func GetRootAsPerson(buf []byte, offset flatbuffers.UOffsetT) *Person {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Person{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Person) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Person) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Person) Fname() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Person) Lname() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Person) Age(obj *DOB) *DOB {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(DOB)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Person) Id() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Person) MutateId(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *Person) Dp() Dept {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return Dept(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 2
}

func (rcv *Person) MutateDp(n Dept) bool {
	return rcv._tab.MutateInt8Slot(12, int8(n))
}

func (rcv *Person) Location(obj *Address) *Address {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Address)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func PersonStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func PersonAddFname(builder *flatbuffers.Builder, fname flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(fname), 0)
}
func PersonAddLname(builder *flatbuffers.Builder, lname flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(lname), 0)
}
func PersonAddAge(builder *flatbuffers.Builder, age flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(age), 0)
}
func PersonAddId(builder *flatbuffers.Builder, id int32) {
	builder.PrependInt32Slot(3, id, 0)
}
func PersonAddDp(builder *flatbuffers.Builder, dp Dept) {
	builder.PrependInt8Slot(4, int8(dp), 2)
}
func PersonAddLocation(builder *flatbuffers.Builder, location flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(location), 0)
}
func PersonEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}