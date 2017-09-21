// automatically generated by the FlatBuffers compiler, do not modify

package structs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type CPUs struct {
	_tab flatbuffers.Table
}

func GetRootAsCPUs(buf []byte, offset flatbuffers.UOffsetT) *CPUs {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CPUs{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *CPUs) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CPUs) Sockets() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPUs) Possible() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPUs) Online() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPUs) CPU(obj *CPU, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
	if obj == nil {
		obj = new(CPU)
	}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *CPUs) CPULength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func CPUsStart(builder *flatbuffers.Builder) { builder.StartObject(4) }
func CPUsAddSockets(builder *flatbuffers.Builder, Sockets int32) { builder.PrependInt32Slot(0, Sockets, 0) }
func CPUsAddPossible(builder *flatbuffers.Builder, Possible flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Possible), 0) }
func CPUsAddOnline(builder *flatbuffers.Builder, Online flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(Online), 0) }
func CPUsAddCPU(builder *flatbuffers.Builder, CPU flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(CPU), 0) }
func CPUsStartCPUVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func CPUsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
