// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type DevUsage struct {
	_tab flatbuffers.Table
}

func GetRootAsDevUsage(buf []byte, offset flatbuffers.UOffsetT) *DevUsage {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DevUsage{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *DevUsage) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DevUsage) Timestamp() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DevUsage) TimeDelta() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DevUsage) Devices(obj *Device, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
	if obj == nil {
		obj = new(Device)
	}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DevUsage) DevicesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DevUsageStart(builder *flatbuffers.Builder) { builder.StartObject(3) }
func DevUsageAddTimestamp(builder *flatbuffers.Builder, Timestamp int64) { builder.PrependInt64Slot(0, Timestamp, 0) }
func DevUsageAddTimeDelta(builder *flatbuffers.Builder, TimeDelta int64) { builder.PrependInt64Slot(1, TimeDelta, 0) }
func DevUsageAddDevices(builder *flatbuffers.Builder, Devices flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(Devices), 0) }
func DevUsageStartDevicesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func DevUsageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
