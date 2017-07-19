// automatically generated by the FlatBuffers compiler, do not modify

package structs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type CPU struct {
	_tab flatbuffers.Table
}

func (rcv *CPU) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CPU) Processor() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) VendorID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) CPUFamily() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) Model() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) ModelName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) Stepping() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) Microcode() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) CPUMHz() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *CPU) CacheSize() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) PhysicalID() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) Siblings() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) CoreID() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) CPUCores() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) APICID() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) InitialAPICID() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) FPU() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) FPUException() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) CPUIDLevel() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) WP() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) Flags(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j * 4))
	}
	return nil
}

func (rcv *CPU) FlagsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CPU) Bugs(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j * 4))
	}
	return nil
}

func (rcv *CPU) BugsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CPU) BogoMIPS() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *CPU) CLFlushSize() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) CacheAlignment() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) AddressSizes(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j * 4))
	}
	return nil
}

func (rcv *CPU) AddressSizesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CPU) PowerManagement(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j * 4))
	}
	return nil
}

func (rcv *CPU) PowerManagementLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CPU) TLBSize() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func CPUStart(builder *flatbuffers.Builder) { builder.StartObject(27) }
func CPUAddProcessor(builder *flatbuffers.Builder, Processor int16) { builder.PrependInt16Slot(0, Processor, 0) }
func CPUAddVendorID(builder *flatbuffers.Builder, VendorID flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(VendorID), 0) }
func CPUAddCPUFamily(builder *flatbuffers.Builder, CPUFamily flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(CPUFamily), 0) }
func CPUAddModel(builder *flatbuffers.Builder, Model flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(Model), 0) }
func CPUAddModelName(builder *flatbuffers.Builder, ModelName flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(ModelName), 0) }
func CPUAddStepping(builder *flatbuffers.Builder, Stepping flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(Stepping), 0) }
func CPUAddMicrocode(builder *flatbuffers.Builder, Microcode flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(Microcode), 0) }
func CPUAddCPUMHz(builder *flatbuffers.Builder, CPUMHz float32) { builder.PrependFloat32Slot(7, CPUMHz, 0.0) }
func CPUAddCacheSize(builder *flatbuffers.Builder, CacheSize flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(CacheSize), 0) }
func CPUAddPhysicalID(builder *flatbuffers.Builder, PhysicalID int16) { builder.PrependInt16Slot(9, PhysicalID, 0) }
func CPUAddSiblings(builder *flatbuffers.Builder, Siblings int16) { builder.PrependInt16Slot(10, Siblings, 0) }
func CPUAddCoreID(builder *flatbuffers.Builder, CoreID int16) { builder.PrependInt16Slot(11, CoreID, 0) }
func CPUAddCPUCores(builder *flatbuffers.Builder, CPUCores int16) { builder.PrependInt16Slot(12, CPUCores, 0) }
func CPUAddAPICID(builder *flatbuffers.Builder, APICID int16) { builder.PrependInt16Slot(13, APICID, 0) }
func CPUAddInitialAPICID(builder *flatbuffers.Builder, InitialAPICID int16) { builder.PrependInt16Slot(14, InitialAPICID, 0) }
func CPUAddFPU(builder *flatbuffers.Builder, FPU flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(15, flatbuffers.UOffsetT(FPU), 0) }
func CPUAddFPUException(builder *flatbuffers.Builder, FPUException flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(16, flatbuffers.UOffsetT(FPUException), 0) }
func CPUAddCPUIDLevel(builder *flatbuffers.Builder, CPUIDLevel flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(17, flatbuffers.UOffsetT(CPUIDLevel), 0) }
func CPUAddWP(builder *flatbuffers.Builder, WP flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(18, flatbuffers.UOffsetT(WP), 0) }
func CPUAddFlags(builder *flatbuffers.Builder, Flags flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(19, flatbuffers.UOffsetT(Flags), 0) }
func CPUStartFlagsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func CPUAddBugs(builder *flatbuffers.Builder, Bugs flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(20, flatbuffers.UOffsetT(Bugs), 0) }
func CPUStartBugsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func CPUAddBogoMIPS(builder *flatbuffers.Builder, BogoMIPS float32) { builder.PrependFloat32Slot(21, BogoMIPS, 0.0) }
func CPUAddCLFlushSize(builder *flatbuffers.Builder, CLFlushSize flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(22, flatbuffers.UOffsetT(CLFlushSize), 0) }
func CPUAddCacheAlignment(builder *flatbuffers.Builder, CacheAlignment flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(23, flatbuffers.UOffsetT(CacheAlignment), 0) }
func CPUAddAddressSizes(builder *flatbuffers.Builder, AddressSizes flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(24, flatbuffers.UOffsetT(AddressSizes), 0) }
func CPUStartAddressSizesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func CPUAddPowerManagement(builder *flatbuffers.Builder, PowerManagement flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(25, flatbuffers.UOffsetT(PowerManagement), 0) }
func CPUStartPowerManagementVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func CPUAddTLBSize(builder *flatbuffers.Builder, TLBSize flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(26, flatbuffers.UOffsetT(TLBSize), 0) }
func CPUEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
