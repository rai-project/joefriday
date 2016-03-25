// automatically generated, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Info struct {
	_tab flatbuffers.Table
}

func GetRootAsInfo(buf []byte, offset flatbuffers.UOffsetT) *Info {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Info{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Info) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Info) Timestamp() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) MemTotal() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) MemFree() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) MemAvailable() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Buffers() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Cached() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) SwapCached() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Active() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Inactive() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) SwapTotal() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) SwapFree() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func InfoStart(builder *flatbuffers.Builder) { builder.StartObject(11) }
func InfoAddTimestamp(builder *flatbuffers.Builder, Timestamp int64) { builder.PrependInt64Slot(0, Timestamp, 0) }
func InfoAddMemTotal(builder *flatbuffers.Builder, MemTotal int64) { builder.PrependInt64Slot(1, MemTotal, 0) }
func InfoAddMemFree(builder *flatbuffers.Builder, MemFree int64) { builder.PrependInt64Slot(2, MemFree, 0) }
func InfoAddMemAvailable(builder *flatbuffers.Builder, MemAvailable int64) { builder.PrependInt64Slot(3, MemAvailable, 0) }
func InfoAddBuffers(builder *flatbuffers.Builder, Buffers int64) { builder.PrependInt64Slot(4, Buffers, 0) }
func InfoAddCached(builder *flatbuffers.Builder, Cached int64) { builder.PrependInt64Slot(5, Cached, 0) }
func InfoAddSwapCached(builder *flatbuffers.Builder, SwapCached int64) { builder.PrependInt64Slot(6, SwapCached, 0) }
func InfoAddActive(builder *flatbuffers.Builder, Active int64) { builder.PrependInt64Slot(7, Active, 0) }
func InfoAddInactive(builder *flatbuffers.Builder, Inactive int64) { builder.PrependInt64Slot(8, Inactive, 0) }
func InfoAddSwapTotal(builder *flatbuffers.Builder, SwapTotal int64) { builder.PrependInt64Slot(9, SwapTotal, 0) }
func InfoAddSwapFree(builder *flatbuffers.Builder, SwapFree int64) { builder.PrependInt64Slot(10, SwapFree, 0) }
func InfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
