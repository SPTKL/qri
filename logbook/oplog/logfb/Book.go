// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package logfb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Book struct {
	_tab flatbuffers.Table
}

func GetRootAsBook(buf []byte, offset flatbuffers.UOffsetT) *Book {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Book{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Book) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Book) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Book) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Book) Identifier() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Book) Logs(obj *Log, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Book) LogsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func BookStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func BookAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func BookAddIdentifier(builder *flatbuffers.Builder, identifier flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(identifier), 0)
}
func BookAddLogs(builder *flatbuffers.Builder, logs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(logs), 0)
}
func BookStartLogsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BookEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}