// Code generated by protoc-gen-go.
// source: limit_disk.proto
// DO NOT EDIT!

package warden

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type LimitDiskRequest struct {
	Handle           *string `protobuf:"bytes,1,req,name=handle" json:"handle,omitempty"`
	BlockLimit       *uint64 `protobuf:"varint,10,opt,name=block_limit" json:"block_limit,omitempty"`
	Block            *uint64 `protobuf:"varint,11,opt,name=block" json:"block,omitempty"`
	BlockSoft        *uint64 `protobuf:"varint,12,opt,name=block_soft" json:"block_soft,omitempty"`
	BlockHard        *uint64 `protobuf:"varint,13,opt,name=block_hard" json:"block_hard,omitempty"`
	InodeLimit       *uint64 `protobuf:"varint,20,opt,name=inode_limit" json:"inode_limit,omitempty"`
	Inode            *uint64 `protobuf:"varint,21,opt,name=inode" json:"inode,omitempty"`
	InodeSoft        *uint64 `protobuf:"varint,22,opt,name=inode_soft" json:"inode_soft,omitempty"`
	InodeHard        *uint64 `protobuf:"varint,23,opt,name=inode_hard" json:"inode_hard,omitempty"`
	ByteLimit        *uint64 `protobuf:"varint,30,opt,name=byte_limit" json:"byte_limit,omitempty"`
	Byte             *uint64 `protobuf:"varint,31,opt,name=byte" json:"byte,omitempty"`
	ByteSoft         *uint64 `protobuf:"varint,32,opt,name=byte_soft" json:"byte_soft,omitempty"`
	ByteHard         *uint64 `protobuf:"varint,33,opt,name=byte_hard" json:"byte_hard,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LimitDiskRequest) Reset()         { *m = LimitDiskRequest{} }
func (m *LimitDiskRequest) String() string { return proto.CompactTextString(m) }
func (*LimitDiskRequest) ProtoMessage()    {}

func (m *LimitDiskRequest) GetHandle() string {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return ""
}

func (m *LimitDiskRequest) GetBlockLimit() uint64 {
	if m != nil && m.BlockLimit != nil {
		return *m.BlockLimit
	}
	return 0
}

func (m *LimitDiskRequest) GetBlock() uint64 {
	if m != nil && m.Block != nil {
		return *m.Block
	}
	return 0
}

func (m *LimitDiskRequest) GetBlockSoft() uint64 {
	if m != nil && m.BlockSoft != nil {
		return *m.BlockSoft
	}
	return 0
}

func (m *LimitDiskRequest) GetBlockHard() uint64 {
	if m != nil && m.BlockHard != nil {
		return *m.BlockHard
	}
	return 0
}

func (m *LimitDiskRequest) GetInodeLimit() uint64 {
	if m != nil && m.InodeLimit != nil {
		return *m.InodeLimit
	}
	return 0
}

func (m *LimitDiskRequest) GetInode() uint64 {
	if m != nil && m.Inode != nil {
		return *m.Inode
	}
	return 0
}

func (m *LimitDiskRequest) GetInodeSoft() uint64 {
	if m != nil && m.InodeSoft != nil {
		return *m.InodeSoft
	}
	return 0
}

func (m *LimitDiskRequest) GetInodeHard() uint64 {
	if m != nil && m.InodeHard != nil {
		return *m.InodeHard
	}
	return 0
}

func (m *LimitDiskRequest) GetByteLimit() uint64 {
	if m != nil && m.ByteLimit != nil {
		return *m.ByteLimit
	}
	return 0
}

func (m *LimitDiskRequest) GetByte() uint64 {
	if m != nil && m.Byte != nil {
		return *m.Byte
	}
	return 0
}

func (m *LimitDiskRequest) GetByteSoft() uint64 {
	if m != nil && m.ByteSoft != nil {
		return *m.ByteSoft
	}
	return 0
}

func (m *LimitDiskRequest) GetByteHard() uint64 {
	if m != nil && m.ByteHard != nil {
		return *m.ByteHard
	}
	return 0
}

type LimitDiskResponse struct {
	BlockLimit       *uint64 `protobuf:"varint,10,opt,name=block_limit" json:"block_limit,omitempty"`
	Block            *uint64 `protobuf:"varint,11,opt,name=block" json:"block,omitempty"`
	BlockSoft        *uint64 `protobuf:"varint,12,opt,name=block_soft" json:"block_soft,omitempty"`
	BlockHard        *uint64 `protobuf:"varint,13,opt,name=block_hard" json:"block_hard,omitempty"`
	InodeLimit       *uint64 `protobuf:"varint,20,opt,name=inode_limit" json:"inode_limit,omitempty"`
	Inode            *uint64 `protobuf:"varint,21,opt,name=inode" json:"inode,omitempty"`
	InodeSoft        *uint64 `protobuf:"varint,22,opt,name=inode_soft" json:"inode_soft,omitempty"`
	InodeHard        *uint64 `protobuf:"varint,23,opt,name=inode_hard" json:"inode_hard,omitempty"`
	ByteLimit        *uint64 `protobuf:"varint,30,opt,name=byte_limit" json:"byte_limit,omitempty"`
	Byte             *uint64 `protobuf:"varint,31,opt,name=byte" json:"byte,omitempty"`
	ByteSoft         *uint64 `protobuf:"varint,32,opt,name=byte_soft" json:"byte_soft,omitempty"`
	ByteHard         *uint64 `protobuf:"varint,33,opt,name=byte_hard" json:"byte_hard,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LimitDiskResponse) Reset()         { *m = LimitDiskResponse{} }
func (m *LimitDiskResponse) String() string { return proto.CompactTextString(m) }
func (*LimitDiskResponse) ProtoMessage()    {}

func (m *LimitDiskResponse) GetBlockLimit() uint64 {
	if m != nil && m.BlockLimit != nil {
		return *m.BlockLimit
	}
	return 0
}

func (m *LimitDiskResponse) GetBlock() uint64 {
	if m != nil && m.Block != nil {
		return *m.Block
	}
	return 0
}

func (m *LimitDiskResponse) GetBlockSoft() uint64 {
	if m != nil && m.BlockSoft != nil {
		return *m.BlockSoft
	}
	return 0
}

func (m *LimitDiskResponse) GetBlockHard() uint64 {
	if m != nil && m.BlockHard != nil {
		return *m.BlockHard
	}
	return 0
}

func (m *LimitDiskResponse) GetInodeLimit() uint64 {
	if m != nil && m.InodeLimit != nil {
		return *m.InodeLimit
	}
	return 0
}

func (m *LimitDiskResponse) GetInode() uint64 {
	if m != nil && m.Inode != nil {
		return *m.Inode
	}
	return 0
}

func (m *LimitDiskResponse) GetInodeSoft() uint64 {
	if m != nil && m.InodeSoft != nil {
		return *m.InodeSoft
	}
	return 0
}

func (m *LimitDiskResponse) GetInodeHard() uint64 {
	if m != nil && m.InodeHard != nil {
		return *m.InodeHard
	}
	return 0
}

func (m *LimitDiskResponse) GetByteLimit() uint64 {
	if m != nil && m.ByteLimit != nil {
		return *m.ByteLimit
	}
	return 0
}

func (m *LimitDiskResponse) GetByte() uint64 {
	if m != nil && m.Byte != nil {
		return *m.Byte
	}
	return 0
}

func (m *LimitDiskResponse) GetByteSoft() uint64 {
	if m != nil && m.ByteSoft != nil {
		return *m.ByteSoft
	}
	return 0
}

func (m *LimitDiskResponse) GetByteHard() uint64 {
	if m != nil && m.ByteHard != nil {
		return *m.ByteHard
	}
	return 0
}

func init() {
}
