// automatically generated by the FlatBuffers compiler, do not modify

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Timestamp struct {
	_tab flatbuffers.Struct
}

func (rcv *Timestamp) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Timestamp) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Timestamp) Unix() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Timestamp) MutateUnix(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Timestamp) Ns() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Timestamp) MutateNs(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func (rcv *Timestamp) Offset() int8 {
	return rcv._tab.GetInt8(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *Timestamp) MutateOffset(n int8) bool {
	return rcv._tab.MutateInt8(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func (rcv *Timestamp) Dst() byte {
	return rcv._tab.GetByte(rcv._tab.Pos + flatbuffers.UOffsetT(9))
}
func (rcv *Timestamp) MutateDst(n byte) bool {
	return rcv._tab.MutateByte(rcv._tab.Pos+flatbuffers.UOffsetT(9), n)
}

func CreateTimestamp(builder *flatbuffers.Builder, unix uint32, ns uint32, offset int8, dst byte) flatbuffers.UOffsetT {
	builder.Prep(4, 12)
	builder.Pad(2)
	builder.PrependByte(dst)
	builder.PrependInt8(offset)
	builder.PrependUint32(ns)
	builder.PrependUint32(unix)
	return builder.Offset()
}
