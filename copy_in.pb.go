// Code generated by protoc-gen-go.
// source: copy_in.proto
// DO NOT EDIT!

package warden

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CopyInRequest struct {
	Handle           *string `protobuf:"bytes,1,req,name=handle" json:"handle,omitempty"`
	SrcPath          *string `protobuf:"bytes,2,req,name=src_path" json:"src_path,omitempty"`
	DstPath          *string `protobuf:"bytes,3,req,name=dst_path" json:"dst_path,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CopyInRequest) Reset()         { *m = CopyInRequest{} }
func (m *CopyInRequest) String() string { return proto.CompactTextString(m) }
func (*CopyInRequest) ProtoMessage()    {}

func (m *CopyInRequest) GetHandle() string {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return ""
}

func (m *CopyInRequest) GetSrcPath() string {
	if m != nil && m.SrcPath != nil {
		return *m.SrcPath
	}
	return ""
}

func (m *CopyInRequest) GetDstPath() string {
	if m != nil && m.DstPath != nil {
		return *m.DstPath
	}
	return ""
}

type CopyInResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CopyInResponse) Reset()         { *m = CopyInResponse{} }
func (m *CopyInResponse) String() string { return proto.CompactTextString(m) }
func (*CopyInResponse) ProtoMessage()    {}

func init() {
}
