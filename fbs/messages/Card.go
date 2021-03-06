// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package messages

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Card struct {
	_tab flatbuffers.Table
}

func GetRootAsCard(buf []byte, offset flatbuffers.UOffsetT) *Card {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Card{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Card) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Card) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Card) Cnum() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Card) MutateCnum(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *Card) Pid() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Card) MutatePid(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func CardStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func CardAddCnum(builder *flatbuffers.Builder, Cnum int32) {
	builder.PrependInt32Slot(0, Cnum, 0)
}
func CardAddPid(builder *flatbuffers.Builder, Pid int32) {
	builder.PrependInt32Slot(1, Pid, 0)
}
func CardEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
