// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cron_fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DatasetOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsDatasetOptions(buf []byte, offset flatbuffers.UOffsetT) *DatasetOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DatasetOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DatasetOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DatasetOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DatasetOptions) Title() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DatasetOptions) Message() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DatasetOptions) BodyPath() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DatasetOptions) FilePaths(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *DatasetOptions) FilePathsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DatasetOptions) Recall() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DatasetOptions) Publish() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *DatasetOptions) MutatePublish(n bool) bool {
	return rcv._tab.MutateBoolSlot(14, n)
}

func (rcv *DatasetOptions) Strict() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *DatasetOptions) MutateStrict(n bool) bool {
	return rcv._tab.MutateBoolSlot(16, n)
}

func (rcv *DatasetOptions) Force() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *DatasetOptions) MutateForce(n bool) bool {
	return rcv._tab.MutateBoolSlot(18, n)
}

func (rcv *DatasetOptions) ConvertFormatToPrev() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *DatasetOptions) MutateConvertFormatToPrev(n bool) bool {
	return rcv._tab.MutateBoolSlot(20, n)
}

func (rcv *DatasetOptions) ShouldRender() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *DatasetOptions) MutateShouldRender(n bool) bool {
	return rcv._tab.MutateBoolSlot(22, n)
}

func (rcv *DatasetOptions) Config(obj *StringMapVal, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DatasetOptions) ConfigLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DatasetOptions) Secrets(obj *StringMapVal, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DatasetOptions) SecretsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DatasetOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(12)
}
func DatasetOptionsAddTitle(builder *flatbuffers.Builder, title flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(title), 0)
}
func DatasetOptionsAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(message), 0)
}
func DatasetOptionsAddBodyPath(builder *flatbuffers.Builder, bodyPath flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(bodyPath), 0)
}
func DatasetOptionsAddFilePaths(builder *flatbuffers.Builder, filePaths flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(filePaths), 0)
}
func DatasetOptionsStartFilePathsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DatasetOptionsAddRecall(builder *flatbuffers.Builder, recall flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(recall), 0)
}
func DatasetOptionsAddPublish(builder *flatbuffers.Builder, publish bool) {
	builder.PrependBoolSlot(5, publish, false)
}
func DatasetOptionsAddStrict(builder *flatbuffers.Builder, strict bool) {
	builder.PrependBoolSlot(6, strict, false)
}
func DatasetOptionsAddForce(builder *flatbuffers.Builder, force bool) {
	builder.PrependBoolSlot(7, force, false)
}
func DatasetOptionsAddConvertFormatToPrev(builder *flatbuffers.Builder, convertFormatToPrev bool) {
	builder.PrependBoolSlot(8, convertFormatToPrev, false)
}
func DatasetOptionsAddShouldRender(builder *flatbuffers.Builder, shouldRender bool) {
	builder.PrependBoolSlot(9, shouldRender, false)
}
func DatasetOptionsAddConfig(builder *flatbuffers.Builder, config flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(config), 0)
}
func DatasetOptionsStartConfigVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DatasetOptionsAddSecrets(builder *flatbuffers.Builder, secrets flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(secrets), 0)
}
func DatasetOptionsStartSecretsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DatasetOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
