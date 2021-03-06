// Code generated by protoc-gen-go. DO NOT EDIT.
// source: capabilities.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GPUVendorType int32

const (
	GPUVendorType_GPU_UNKNOWN GPUVendorType = 0
	GPUVendorType_NVIDIA      GPUVendorType = 1
	GPUVendorType_RADEON      GPUVendorType = 2
	GPUVendorType_FAKE        GPUVendorType = 99
)

var GPUVendorType_name = map[int32]string{
	0:  "GPU_UNKNOWN",
	1:  "NVIDIA",
	2:  "RADEON",
	99: "FAKE",
}
var GPUVendorType_value = map[string]int32{
	"GPU_UNKNOWN": 0,
	"NVIDIA":      1,
	"RADEON":      2,
	"FAKE":        99,
}

func (x GPUVendorType) String() string {
	return proto.EnumName(GPUVendorType_name, int32(x))
}
func (GPUVendorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type CPUDevice struct {
	// ModelName describes full model name.
	// For example "Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz".
	ModelName string `protobuf:"bytes,1,opt,name=modelName" json:"modelName,omitempty"`
	// Cores describes number of cores on a CPU device.
	Cores uint32 `protobuf:"varint,2,opt,name=cores" json:"cores,omitempty"`
	// Sockets describes number of CPU sockets on a host system.
	Sockets uint32 `protobuf:"varint,3,opt,name=sockets" json:"sockets,omitempty"`
}

func (m *CPUDevice) Reset()                    { *m = CPUDevice{} }
func (m *CPUDevice) String() string            { return proto.CompactTextString(m) }
func (*CPUDevice) ProtoMessage()               {}
func (*CPUDevice) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *CPUDevice) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *CPUDevice) GetCores() uint32 {
	if m != nil {
		return m.Cores
	}
	return 0
}

func (m *CPUDevice) GetSockets() uint32 {
	if m != nil {
		return m.Sockets
	}
	return 0
}

type CPU struct {
	Device     *CPUDevice            `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	Benchmarks map[uint64]*Benchmark `protobuf:"bytes,2,rep,name=benchmarks" json:"benchmarks,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CPU) Reset()                    { *m = CPU{} }
func (m *CPU) String() string            { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()               {}
func (*CPU) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *CPU) GetDevice() *CPUDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *CPU) GetBenchmarks() map[uint64]*Benchmark {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

type RAMDevice struct {
	Total     uint64 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
	Available uint64 `protobuf:"varint,2,opt,name=available" json:"available,omitempty"`
}

func (m *RAMDevice) Reset()                    { *m = RAMDevice{} }
func (m *RAMDevice) String() string            { return proto.CompactTextString(m) }
func (*RAMDevice) ProtoMessage()               {}
func (*RAMDevice) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *RAMDevice) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RAMDevice) GetAvailable() uint64 {
	if m != nil {
		return m.Available
	}
	return 0
}

type RAM struct {
	Device     *RAMDevice            `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	Benchmarks map[uint64]*Benchmark `protobuf:"bytes,2,rep,name=benchmarks" json:"benchmarks,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RAM) Reset()                    { *m = RAM{} }
func (m *RAM) String() string            { return proto.CompactTextString(m) }
func (*RAM) ProtoMessage()               {}
func (*RAM) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *RAM) GetDevice() *RAMDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *RAM) GetBenchmarks() map[uint64]*Benchmark {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

type GPUDevice struct {
	// ID returns unique device ID on workers machine,
	// typically PCI bus ID
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	// VendorID returns an unique device vendor identifier
	VendorID uint64 `protobuf:"varint,2,opt,name=vendorID" json:"vendorID,omitempty"`
	// VendorName returns GPU vendor name.
	VendorName string `protobuf:"bytes,3,opt,name=vendorName" json:"vendorName,omitempty"`
	// DeviceID returns device ID (e.g.: NVidia)
	DeviceID uint64 `protobuf:"varint,5,opt,name=deviceID" json:"deviceID,omitempty"`
	// DeviceName returns device name, (e.g.: 1080Ti)
	DeviceName string `protobuf:"bytes,6,opt,name=deviceName" json:"deviceName,omitempty"`
	// MajorNumber returns device's major number
	MajorNumber uint64 `protobuf:"varint,7,opt,name=majorNumber" json:"majorNumber,omitempty"`
	// MinorNumber returns device's minor number
	MinorNumber uint64 `protobuf:"varint,8,opt,name=minorNumber" json:"minorNumber,omitempty"`
	// Memory is amount of vmem for device, in bytes
	Memory uint64 `protobuf:"varint,9,opt,name=Memory" json:"Memory,omitempty"`
	// Hash string built from device parameters
	Hash string `protobuf:"bytes,10,opt,name=hash" json:"hash,omitempty"`
}

func (m *GPUDevice) Reset()                    { *m = GPUDevice{} }
func (m *GPUDevice) String() string            { return proto.CompactTextString(m) }
func (*GPUDevice) ProtoMessage()               {}
func (*GPUDevice) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *GPUDevice) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GPUDevice) GetVendorID() uint64 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *GPUDevice) GetVendorName() string {
	if m != nil {
		return m.VendorName
	}
	return ""
}

func (m *GPUDevice) GetDeviceID() uint64 {
	if m != nil {
		return m.DeviceID
	}
	return 0
}

func (m *GPUDevice) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *GPUDevice) GetMajorNumber() uint64 {
	if m != nil {
		return m.MajorNumber
	}
	return 0
}

func (m *GPUDevice) GetMinorNumber() uint64 {
	if m != nil {
		return m.MinorNumber
	}
	return 0
}

func (m *GPUDevice) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *GPUDevice) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GPU struct {
	Device     *GPUDevice            `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	Benchmarks map[uint64]*Benchmark `protobuf:"bytes,2,rep,name=benchmarks" json:"benchmarks,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GPU) Reset()                    { *m = GPU{} }
func (m *GPU) String() string            { return proto.CompactTextString(m) }
func (*GPU) ProtoMessage()               {}
func (*GPU) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *GPU) GetDevice() *GPUDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *GPU) GetBenchmarks() map[uint64]*Benchmark {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

type NetFlags struct {
	Flags uint64 `protobuf:"varint,1,opt,name=flags" json:"flags,omitempty"`
}

func (m *NetFlags) Reset()                    { *m = NetFlags{} }
func (m *NetFlags) String() string            { return proto.CompactTextString(m) }
func (*NetFlags) ProtoMessage()               {}
func (*NetFlags) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *NetFlags) GetFlags() uint64 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type Network struct {
	In            uint64                `protobuf:"varint,1,opt,name=in" json:"in,omitempty"`
	Out           uint64                `protobuf:"varint,2,opt,name=out" json:"out,omitempty"`
	NetFlags      *NetFlags             `protobuf:"bytes,3,opt,name=netFlags" json:"netFlags,omitempty"`
	BenchmarksIn  map[uint64]*Benchmark `protobuf:"bytes,4,rep,name=benchmarksIn" json:"benchmarksIn,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BenchmarksOut map[uint64]*Benchmark `protobuf:"bytes,5,rep,name=benchmarksOut" json:"benchmarksOut,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Network) Reset()                    { *m = Network{} }
func (m *Network) String() string            { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()               {}
func (*Network) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *Network) GetIn() uint64 {
	if m != nil {
		return m.In
	}
	return 0
}

func (m *Network) GetOut() uint64 {
	if m != nil {
		return m.Out
	}
	return 0
}

func (m *Network) GetNetFlags() *NetFlags {
	if m != nil {
		return m.NetFlags
	}
	return nil
}

func (m *Network) GetBenchmarksIn() map[uint64]*Benchmark {
	if m != nil {
		return m.BenchmarksIn
	}
	return nil
}

func (m *Network) GetBenchmarksOut() map[uint64]*Benchmark {
	if m != nil {
		return m.BenchmarksOut
	}
	return nil
}

type StorageDevice struct {
	BytesAvailable uint64 `protobuf:"varint,1,opt,name=bytesAvailable" json:"bytesAvailable,omitempty"`
}

func (m *StorageDevice) Reset()                    { *m = StorageDevice{} }
func (m *StorageDevice) String() string            { return proto.CompactTextString(m) }
func (*StorageDevice) ProtoMessage()               {}
func (*StorageDevice) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *StorageDevice) GetBytesAvailable() uint64 {
	if m != nil {
		return m.BytesAvailable
	}
	return 0
}

type Storage struct {
	Device     *StorageDevice        `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	Benchmarks map[uint64]*Benchmark `protobuf:"bytes,2,rep,name=benchmarks" json:"benchmarks,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Storage) Reset()                    { *m = Storage{} }
func (m *Storage) String() string            { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()               {}
func (*Storage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *Storage) GetDevice() *StorageDevice {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *Storage) GetBenchmarks() map[uint64]*Benchmark {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

func init() {
	proto.RegisterType((*CPUDevice)(nil), "sonm.CPUDevice")
	proto.RegisterType((*CPU)(nil), "sonm.CPU")
	proto.RegisterType((*RAMDevice)(nil), "sonm.RAMDevice")
	proto.RegisterType((*RAM)(nil), "sonm.RAM")
	proto.RegisterType((*GPUDevice)(nil), "sonm.GPUDevice")
	proto.RegisterType((*GPU)(nil), "sonm.GPU")
	proto.RegisterType((*NetFlags)(nil), "sonm.NetFlags")
	proto.RegisterType((*Network)(nil), "sonm.Network")
	proto.RegisterType((*StorageDevice)(nil), "sonm.StorageDevice")
	proto.RegisterType((*Storage)(nil), "sonm.Storage")
	proto.RegisterEnum("sonm.GPUVendorType", GPUVendorType_name, GPUVendorType_value)
}

func init() { proto.RegisterFile("capabilities.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xc6, 0xce, 0x7b, 0x42, 0xda, 0xb0, 0x54, 0xc8, 0x44, 0x3c, 0xa2, 0x48, 0x40, 0x55, 0xa4,
	0x1c, 0xca, 0x81, 0x87, 0x84, 0x90, 0x5b, 0xb7, 0x91, 0x55, 0xc5, 0x09, 0x0b, 0x29, 0xe2, 0x84,
	0x36, 0xe9, 0xd2, 0x9a, 0xd8, 0xde, 0xca, 0xde, 0x04, 0xe5, 0xcc, 0x89, 0x9f, 0xc4, 0x0f, 0xe0,
	0x7f, 0xa1, 0x7d, 0x64, 0x9d, 0xb4, 0x54, 0x02, 0x55, 0xea, 0x6d, 0x5e, 0xdf, 0x37, 0x3b, 0x33,
	0x5f, 0x1c, 0x40, 0x13, 0x72, 0x4e, 0xc6, 0x61, 0x14, 0xf2, 0x90, 0x66, 0xdd, 0xf3, 0x94, 0x71,
	0x86, 0x8a, 0x19, 0x4b, 0xe2, 0x56, 0x73, 0x4c, 0x93, 0xc9, 0x59, 0x4c, 0xd2, 0xa9, 0x8e, 0x77,
	0x3e, 0x43, 0x6d, 0x7f, 0x38, 0xf2, 0xe8, 0x3c, 0x9c, 0x50, 0xf4, 0x00, 0x6a, 0x31, 0x3b, 0xa1,
	0x51, 0x40, 0x62, 0xea, 0x58, 0x6d, 0x6b, 0xbb, 0x86, 0xf3, 0x00, 0xda, 0x82, 0xd2, 0x84, 0xa5,
	0x34, 0x73, 0xec, 0xb6, 0xb5, 0xdd, 0xc0, 0xca, 0x41, 0x0e, 0x54, 0x32, 0x36, 0x99, 0x52, 0x9e,
	0x39, 0x05, 0x19, 0x5f, 0xba, 0x9d, 0x5f, 0x16, 0x14, 0xf6, 0x87, 0x23, 0xf4, 0x0c, 0xca, 0x27,
	0x92, 0x5f, 0x52, 0xd6, 0x77, 0x37, 0xbb, 0xe2, 0x2d, 0x5d, 0xd3, 0x16, 0xeb, 0x34, 0x7a, 0x0d,
	0x90, 0xbf, 0xcf, 0xb1, 0xdb, 0x85, 0xed, 0xfa, 0xee, 0x7d, 0x53, 0xdc, 0xdd, 0x33, 0xb9, 0x83,
	0x84, 0xa7, 0x0b, 0xbc, 0x52, 0xdc, 0x0a, 0x60, 0xf3, 0x42, 0x1a, 0x35, 0xa1, 0x30, 0xa5, 0x0b,
	0xd9, 0xb3, 0x88, 0x85, 0x89, 0x9e, 0x40, 0x69, 0x4e, 0xa2, 0x19, 0x95, 0x03, 0x98, 0x77, 0x18,
	0x1c, 0x56, 0xd9, 0x37, 0xf6, 0x2b, 0xab, 0xf3, 0x0e, 0x6a, 0xd8, 0xed, 0xeb, 0xb5, 0x6c, 0x41,
	0x89, 0x33, 0x4e, 0x22, 0xcd, 0xa5, 0x1c, 0xb1, 0x2c, 0x32, 0x27, 0x61, 0x44, 0xc6, 0x91, 0x62,
	0x2c, 0xe2, 0x3c, 0x20, 0x87, 0xc7, 0x6e, 0xff, 0xaa, 0xe1, 0x0d, 0xf9, 0xbf, 0x0c, 0x8f, 0xdd,
	0xfe, 0x8d, 0x0e, 0xff, 0xc3, 0x86, 0x5a, 0xcf, 0x88, 0x62, 0x03, 0x6c, 0xdf, 0xd3, 0x6a, 0xb0,
	0x7d, 0x0f, 0xb5, 0xa0, 0x3a, 0xa7, 0xc9, 0x09, 0x4b, 0x7d, 0x4f, 0x8f, 0x6d, 0x7c, 0xf4, 0x08,
	0x40, 0xd9, 0x52, 0x41, 0x05, 0x89, 0x59, 0x89, 0x08, 0xac, 0x1a, 0xd7, 0xf7, 0x9c, 0x92, 0xc2,
	0x2e, 0x7d, 0x81, 0x55, 0xb6, 0xc4, 0x96, 0x15, 0x36, 0x8f, 0xa0, 0x36, 0xd4, 0x63, 0xf2, 0x8d,
	0xa5, 0xc1, 0x2c, 0x1e, 0xd3, 0xd4, 0xa9, 0x48, 0xf8, 0x6a, 0x48, 0x56, 0x84, 0x89, 0xa9, 0xa8,
	0xea, 0x8a, 0x3c, 0x84, 0xee, 0x41, 0xb9, 0x4f, 0x63, 0x96, 0x2e, 0x9c, 0x9a, 0x4c, 0x6a, 0x0f,
	0x21, 0x28, 0x9e, 0x91, 0xec, 0xcc, 0x01, 0xd9, 0x55, 0xda, 0xf2, 0x82, 0xbd, 0xab, 0xe5, 0xdb,
	0xfb, 0x1f, 0xf9, 0xf6, 0x6e, 0x58, 0xbe, 0x6d, 0xa8, 0x06, 0x94, 0x1f, 0x46, 0xe4, 0x34, 0x13,
	0xea, 0xfd, 0x2a, 0x8c, 0xa5, 0x7a, 0xa5, 0xd3, 0xf9, 0x59, 0x80, 0x4a, 0x40, 0xf9, 0x77, 0x96,
	0x4e, 0xc5, 0x85, 0xc3, 0x44, 0xa7, 0xed, 0x30, 0x11, 0xad, 0xd9, 0x8c, 0xeb, 0xe3, 0x0a, 0x13,
	0xed, 0x40, 0x35, 0xd1, 0x7c, 0xf2, 0xaa, 0xf5, 0xdd, 0x0d, 0xd5, 0x7d, 0xd9, 0x05, 0x9b, 0x3c,
	0xda, 0x87, 0xdb, 0xf9, 0x64, 0x7e, 0xe2, 0x14, 0xe5, 0x22, 0x1e, 0x9b, 0x7a, 0xd1, 0x72, 0x65,
	0x19, 0x7e, 0xa2, 0xd6, 0xb1, 0x06, 0x42, 0x87, 0xd0, 0xc8, 0xfd, 0xc1, 0x8c, 0x3b, 0x25, 0xc9,
	0xd2, 0xbe, 0x8a, 0x65, 0x30, 0xe3, 0x8a, 0x66, 0x1d, 0xd6, 0x1a, 0xc2, 0x9d, 0x4b, 0xad, 0xae,
	0xb5, 0xda, 0xd6, 0x7b, 0x40, 0x97, 0xdb, 0x5e, 0xef, 0x5a, 0x2f, 0xa1, 0xf1, 0x81, 0xb3, 0x94,
	0x9c, 0x52, 0xfd, 0x93, 0x7b, 0x0a, 0x1b, 0xe3, 0x05, 0xa7, 0x99, 0x6b, 0xbe, 0x2f, 0x8a, 0xf8,
	0x42, 0xb4, 0xf3, 0xdb, 0x82, 0x8a, 0x46, 0xa2, 0xe7, 0x17, 0x64, 0x7a, 0x57, 0x35, 0x5c, 0x23,
	0x36, 0x52, 0x7d, 0xfb, 0x17, 0xa9, 0x3e, 0x5c, 0x03, 0xdc, 0xa4, 0x5c, 0x77, 0xf6, 0xa0, 0xd1,
	0x1b, 0x8e, 0x8e, 0xe5, 0x77, 0xe2, 0xe3, 0xe2, 0x9c, 0xa2, 0x4d, 0xa8, 0xf7, 0x86, 0xa3, 0x2f,
	0xa3, 0xe0, 0x28, 0x18, 0x7c, 0x0a, 0x9a, 0xb7, 0x10, 0x40, 0x39, 0x38, 0xf6, 0x3d, 0xdf, 0x6d,
	0x5a, 0xc2, 0xc6, 0xae, 0x77, 0x30, 0x08, 0x9a, 0x36, 0xaa, 0x42, 0xf1, 0xd0, 0x3d, 0x3a, 0x68,
	0x4e, 0xc6, 0x65, 0xf9, 0x7f, 0xf6, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xd4, 0x5e,
	0x77, 0xfd, 0x06, 0x00, 0x00,
}
