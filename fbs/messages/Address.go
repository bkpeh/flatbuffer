// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package messages

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Address struct {
	_tab flatbuffers.Table
}

func GetRootAsAddress(buf []byte, offset flatbuffers.UOffsetT) *Address {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Address{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Address) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Address) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Address) Add() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Address) Phone() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Address) MutatePhone(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func AddressStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AddressAddAdd(builder *flatbuffers.Builder, add flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(add), 0)
}
func AddressAddPhone(builder *flatbuffers.Builder, phone int32) {
	builder.PrependInt32Slot(1, phone, 0)
}
func AddressEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}