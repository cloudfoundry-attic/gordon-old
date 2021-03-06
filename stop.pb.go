// Code generated by protoc-gen-go.
// source: stop.proto
// DO NOT EDIT!

package warden

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type StopRequest struct {
	Handle           *string `protobuf:"bytes,1,req,name=handle" json:"handle,omitempty"`
	Background       *bool   `protobuf:"varint,10,opt,name=background,def=0" json:"background,omitempty"`
	Kill             *bool   `protobuf:"varint,20,opt,name=kill,def=0" json:"kill,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}

const Default_StopRequest_Background bool = false
const Default_StopRequest_Kill bool = false

func (m *StopRequest) GetHandle() string {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return ""
}

func (m *StopRequest) GetBackground() bool {
	if m != nil && m.Background != nil {
		return *m.Background
	}
	return Default_StopRequest_Background
}

func (m *StopRequest) GetKill() bool {
	if m != nil && m.Kill != nil {
		return *m.Kill
	}
	return Default_StopRequest_Kill
}

type StopResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}

func init() {
}
