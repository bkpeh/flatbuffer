// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package messages

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DOB struct {
	_tab flatbuffers.Struct
}

func (rcv *DOB) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DOB) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *DOB) Year() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *DOB) MutateYear(n int32) bool {
	return rcv._tab.MutateInt32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *DOB) Month() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *DOB) MutateMonth(n int32) bool {
	return rcv._tab.MutateInt32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func (rcv *DOB) Day() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *DOB) MutateDay(n int32) bool {
	return rcv._tab.MutateInt32(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func CreateDOB(builder *flatbuffers.Builder, year int32, month int32, day int32) flatbuffers.UOffsetT {
	builder.Prep(4, 12)
	builder.PrependInt32(day)
	builder.PrependInt32(month)
	builder.PrependInt32(year)
	return builder.Offset()
}
