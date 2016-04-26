// automatically generated, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Uptime struct {
	_tab flatbuffers.Table
}

func GetRootAsUptime(buf []byte, offset flatbuffers.UOffsetT) *Uptime {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Uptime{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Uptime) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Uptime) Timestamp() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Uptime) Total() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Uptime) Idle() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0
}

func UptimeStart(builder *flatbuffers.Builder) { builder.StartObject(3) }
func UptimeAddTimestamp(builder *flatbuffers.Builder, Timestamp int64) { builder.PrependInt64Slot(0, Timestamp, 0) }
func UptimeAddTotal(builder *flatbuffers.Builder, Total float64) { builder.PrependFloat64Slot(1, Total, 0) }
func UptimeAddIdle(builder *flatbuffers.Builder, Idle float64) { builder.PrependFloat64Slot(2, Idle, 0) }
func UptimeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
