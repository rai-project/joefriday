// automatically generated, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type LoadAvg struct {
	_tab flatbuffers.Table
}

func GetRootAsLoadAvg(buf []byte, offset flatbuffers.UOffsetT) *LoadAvg {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LoadAvg{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *LoadAvg) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LoadAvg) Timestamp() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadAvg) Minute() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadAvg) Five() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadAvg) Fifteen() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadAvg) Running() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadAvg) Total() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadAvg) PID() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func LoadAvgStart(builder *flatbuffers.Builder) { builder.StartObject(7) }
func LoadAvgAddTimestamp(builder *flatbuffers.Builder, Timestamp int64) { builder.PrependInt64Slot(0, Timestamp, 0) }
func LoadAvgAddMinute(builder *flatbuffers.Builder, Minute float32) { builder.PrependFloat32Slot(1, Minute, 0) }
func LoadAvgAddFive(builder *flatbuffers.Builder, Five float32) { builder.PrependFloat32Slot(2, Five, 0) }
func LoadAvgAddFifteen(builder *flatbuffers.Builder, Fifteen float32) { builder.PrependFloat32Slot(3, Fifteen, 0) }
func LoadAvgAddRunning(builder *flatbuffers.Builder, Running int32) { builder.PrependInt32Slot(4, Running, 0) }
func LoadAvgAddTotal(builder *flatbuffers.Builder, Total int32) { builder.PrependInt32Slot(5, Total, 0) }
func LoadAvgAddPID(builder *flatbuffers.Builder, PID int32) { builder.PrependInt32Slot(6, PID, 0) }
func LoadAvgEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
