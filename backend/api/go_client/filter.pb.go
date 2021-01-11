// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend/api/filter.proto

package go_client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Op is the operation to apply.
type Predicate_Op int32

const (
	Predicate_UNKNOWN Predicate_Op = 0
	// Operators on scalar values. Only applies to one of |int_value|,
	// |long_value|, |string_value| or |timestamp_value|.
	Predicate_EQUALS              Predicate_Op = 1
	Predicate_NOT_EQUALS          Predicate_Op = 2
	Predicate_GREATER_THAN        Predicate_Op = 3
	Predicate_GREATER_THAN_EQUALS Predicate_Op = 5
	Predicate_LESS_THAN           Predicate_Op = 6
	Predicate_LESS_THAN_EQUALS    Predicate_Op = 7
	// Checks if the value is a member of a given array, which should be one of
	// |int_values|, |long_values| or |string_values|.
	Predicate_IN Predicate_Op = 8
	// Checks if the value contains |string_value| as a substring match. Only
	// applies to |string_value|.
	Predicate_IS_SUBSTRING Predicate_Op = 9
)

var Predicate_Op_name = map[int32]string{
	0: "UNKNOWN",
	1: "EQUALS",
	2: "NOT_EQUALS",
	3: "GREATER_THAN",
	5: "GREATER_THAN_EQUALS",
	6: "LESS_THAN",
	7: "LESS_THAN_EQUALS",
	8: "IN",
	9: "IS_SUBSTRING",
}

var Predicate_Op_value = map[string]int32{
	"UNKNOWN":             0,
	"EQUALS":              1,
	"NOT_EQUALS":          2,
	"GREATER_THAN":        3,
	"GREATER_THAN_EQUALS": 5,
	"LESS_THAN":           6,
	"LESS_THAN_EQUALS":    7,
	"IN":                  8,
	"IS_SUBSTRING":        9,
}

func (x Predicate_Op) String() string {
	return proto.EnumName(Predicate_Op_name, int32(x))
}

func (Predicate_Op) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aab96529e99c2762, []int{0, 0}
}

// Predicate captures individual conditions that must be true for a resource
// being filtered.
type Predicate struct {
	Op  Predicate_Op `protobuf:"varint,1,opt,name=op,proto3,enum=api.Predicate_Op" json:"op,omitempty"`
	Key string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Predicate_IntValue
	//	*Predicate_LongValue
	//	*Predicate_StringValue
	//	*Predicate_TimestampValue
	//	*Predicate_IntValues
	//	*Predicate_LongValues
	//	*Predicate_StringValues
	Value                isPredicate_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Predicate) Reset()         { *m = Predicate{} }
func (m *Predicate) String() string { return proto.CompactTextString(m) }
func (*Predicate) ProtoMessage()    {}
func (*Predicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_aab96529e99c2762, []int{0}
}

func (m *Predicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Predicate.Unmarshal(m, b)
}
func (m *Predicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Predicate.Marshal(b, m, deterministic)
}
func (m *Predicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Predicate.Merge(m, src)
}
func (m *Predicate) XXX_Size() int {
	return xxx_messageInfo_Predicate.Size(m)
}
func (m *Predicate) XXX_DiscardUnknown() {
	xxx_messageInfo_Predicate.DiscardUnknown(m)
}

var xxx_messageInfo_Predicate proto.InternalMessageInfo

func (m *Predicate) GetOp() Predicate_Op {
	if m != nil {
		return m.Op
	}
	return Predicate_UNKNOWN
}

func (m *Predicate) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type isPredicate_Value interface {
	isPredicate_Value()
}

type Predicate_IntValue struct {
	IntValue int32 `protobuf:"varint,3,opt,name=int_value,json=intValue,proto3,oneof"`
}

type Predicate_LongValue struct {
	LongValue int64 `protobuf:"varint,4,opt,name=long_value,json=longValue,proto3,oneof"`
}

type Predicate_StringValue struct {
	StringValue string `protobuf:"bytes,5,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Predicate_TimestampValue struct {
	TimestampValue *timestamp.Timestamp `protobuf:"bytes,6,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type Predicate_IntValues struct {
	IntValues *IntValues `protobuf:"bytes,7,opt,name=int_values,json=intValues,proto3,oneof"`
}

type Predicate_LongValues struct {
	LongValues *LongValues `protobuf:"bytes,8,opt,name=long_values,json=longValues,proto3,oneof"`
}

type Predicate_StringValues struct {
	StringValues *StringValues `protobuf:"bytes,9,opt,name=string_values,json=stringValues,proto3,oneof"`
}

func (*Predicate_IntValue) isPredicate_Value() {}

func (*Predicate_LongValue) isPredicate_Value() {}

func (*Predicate_StringValue) isPredicate_Value() {}

func (*Predicate_TimestampValue) isPredicate_Value() {}

func (*Predicate_IntValues) isPredicate_Value() {}

func (*Predicate_LongValues) isPredicate_Value() {}

func (*Predicate_StringValues) isPredicate_Value() {}

func (m *Predicate) GetValue() isPredicate_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Predicate) GetIntValue() int32 {
	if x, ok := m.GetValue().(*Predicate_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *Predicate) GetLongValue() int64 {
	if x, ok := m.GetValue().(*Predicate_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (m *Predicate) GetStringValue() string {
	if x, ok := m.GetValue().(*Predicate_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Predicate) GetTimestampValue() *timestamp.Timestamp {
	if x, ok := m.GetValue().(*Predicate_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (m *Predicate) GetIntValues() *IntValues {
	if x, ok := m.GetValue().(*Predicate_IntValues); ok {
		return x.IntValues
	}
	return nil
}

func (m *Predicate) GetLongValues() *LongValues {
	if x, ok := m.GetValue().(*Predicate_LongValues); ok {
		return x.LongValues
	}
	return nil
}

func (m *Predicate) GetStringValues() *StringValues {
	if x, ok := m.GetValue().(*Predicate_StringValues); ok {
		return x.StringValues
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Predicate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Predicate_IntValue)(nil),
		(*Predicate_LongValue)(nil),
		(*Predicate_StringValue)(nil),
		(*Predicate_TimestampValue)(nil),
		(*Predicate_IntValues)(nil),
		(*Predicate_LongValues)(nil),
		(*Predicate_StringValues)(nil),
	}
}

type IntValues struct {
	Values               []int32  `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntValues) Reset()         { *m = IntValues{} }
func (m *IntValues) String() string { return proto.CompactTextString(m) }
func (*IntValues) ProtoMessage()    {}
func (*IntValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_aab96529e99c2762, []int{1}
}

func (m *IntValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntValues.Unmarshal(m, b)
}
func (m *IntValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntValues.Marshal(b, m, deterministic)
}
func (m *IntValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntValues.Merge(m, src)
}
func (m *IntValues) XXX_Size() int {
	return xxx_messageInfo_IntValues.Size(m)
}
func (m *IntValues) XXX_DiscardUnknown() {
	xxx_messageInfo_IntValues.DiscardUnknown(m)
}

var xxx_messageInfo_IntValues proto.InternalMessageInfo

func (m *IntValues) GetValues() []int32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type StringValues struct {
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringValues) Reset()         { *m = StringValues{} }
func (m *StringValues) String() string { return proto.CompactTextString(m) }
func (*StringValues) ProtoMessage()    {}
func (*StringValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_aab96529e99c2762, []int{2}
}

func (m *StringValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringValues.Unmarshal(m, b)
}
func (m *StringValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringValues.Marshal(b, m, deterministic)
}
func (m *StringValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringValues.Merge(m, src)
}
func (m *StringValues) XXX_Size() int {
	return xxx_messageInfo_StringValues.Size(m)
}
func (m *StringValues) XXX_DiscardUnknown() {
	xxx_messageInfo_StringValues.DiscardUnknown(m)
}

var xxx_messageInfo_StringValues proto.InternalMessageInfo

func (m *StringValues) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type LongValues struct {
	Values               []int64  `protobuf:"varint,3,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongValues) Reset()         { *m = LongValues{} }
func (m *LongValues) String() string { return proto.CompactTextString(m) }
func (*LongValues) ProtoMessage()    {}
func (*LongValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_aab96529e99c2762, []int{3}
}

func (m *LongValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongValues.Unmarshal(m, b)
}
func (m *LongValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongValues.Marshal(b, m, deterministic)
}
func (m *LongValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongValues.Merge(m, src)
}
func (m *LongValues) XXX_Size() int {
	return xxx_messageInfo_LongValues.Size(m)
}
func (m *LongValues) XXX_DiscardUnknown() {
	xxx_messageInfo_LongValues.DiscardUnknown(m)
}

var xxx_messageInfo_LongValues proto.InternalMessageInfo

func (m *LongValues) GetValues() []int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

// Filter is used to filter resources returned from a ListXXX request.
//
// Example filters:
// 1) Filter runs with status = 'Running'
// filter {
//   predicate {
//     key: "status"
//     op: EQUALS
//     string_value: "Running"
//   }
// }
//
// 2) Filter runs that succeeded since Dec 1, 2018
// filter {
//   predicate {
//     key: "status"
//     op: EQUALS
//     string_value: "Succeeded"
//   }
//   predicate {
//     key: "created_at"
//     op: GREATER_THAN
//     timestamp_value {
//       seconds: 1543651200
//     }
//   }
// }
//
// 3) Filter runs with one of labels 'label_1' or 'label_2'
//
// filter {
//   predicate {
//     key: "label"
//     op: IN
//     string_values {
//       value: 'label_1'
//       value: 'label_2'
//     }
//   }
// }
type Filter struct {
	// All predicates are AND-ed when this filter is applied.
	Predicates           []*Predicate `protobuf:"bytes,1,rep,name=predicates,proto3" json:"predicates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_aab96529e99c2762, []int{4}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetPredicates() []*Predicate {
	if m != nil {
		return m.Predicates
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.Predicate_Op", Predicate_Op_name, Predicate_Op_value)
	proto.RegisterType((*Predicate)(nil), "api.Predicate")
	proto.RegisterType((*IntValues)(nil), "api.IntValues")
	proto.RegisterType((*StringValues)(nil), "api.StringValues")
	proto.RegisterType((*LongValues)(nil), "api.LongValues")
	proto.RegisterType((*Filter)(nil), "api.Filter")
}

func init() { proto.RegisterFile("backend/api/filter.proto", fileDescriptor_aab96529e99c2762) }

var fileDescriptor_aab96529e99c2762 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x93, 0xdf, 0x8f, 0xd2, 0x40,
	0x10, 0xc7, 0xfb, 0xe3, 0x28, 0xd7, 0x81, 0xe3, 0xea, 0x6a, 0xb4, 0x21, 0x9a, 0xab, 0x9c, 0xd1,
	0x3e, 0xb5, 0x09, 0x17, 0x93, 0x7b, 0xf1, 0x01, 0x22, 0x02, 0x91, 0x14, 0x6d, 0x41, 0x13, 0x5f,
	0x48, 0xe1, 0x16, 0xdc, 0x50, 0xba, 0x1b, 0xba, 0x9c, 0xb9, 0xbf, 0xc4, 0xff, 0xc2, 0xbf, 0xd1,
	0xb4, 0xdb, 0x2e, 0x7d, 0xeb, 0xcc, 0x7c, 0xbe, 0xd3, 0xf9, 0x4e, 0xa7, 0x60, 0xaf, 0xe3, 0xcd,
	0x1e, 0xa7, 0x0f, 0x7e, 0xcc, 0x88, 0xbf, 0x25, 0x09, 0xc7, 0x47, 0x8f, 0x1d, 0x29, 0xa7, 0x48,
	0x8f, 0x19, 0xe9, 0xbe, 0xde, 0x51, 0xba, 0x4b, 0x70, 0x51, 0x8d, 0xd3, 0x94, 0xf2, 0x98, 0x13,
	0x9a, 0x66, 0x02, 0xe9, 0xde, 0x94, 0xd5, 0x22, 0x5a, 0x9f, 0xb6, 0x3e, 0x27, 0x07, 0x9c, 0xf1,
	0xf8, 0xc0, 0x04, 0xd0, 0xfb, 0x77, 0x01, 0xe6, 0xb7, 0x23, 0x7e, 0x20, 0x9b, 0x98, 0x63, 0xf4,
	0x16, 0x34, 0xca, 0x6c, 0xd5, 0x51, 0xdd, 0x4e, 0xff, 0x99, 0x17, 0x33, 0xe2, 0xc9, 0x9a, 0x37,
	0x67, 0xa1, 0x46, 0x19, 0xb2, 0x40, 0xdf, 0xe3, 0x27, 0x5b, 0x73, 0x54, 0xd7, 0x0c, 0xf3, 0x47,
	0xf4, 0x06, 0x4c, 0x92, 0xf2, 0xd5, 0x63, 0x9c, 0x9c, 0xb0, 0xad, 0x3b, 0xaa, 0xdb, 0x98, 0x28,
	0xe1, 0x25, 0x49, 0xf9, 0x8f, 0x3c, 0x83, 0x6e, 0x00, 0x12, 0x9a, 0xee, 0xca, 0xfa, 0x85, 0xa3,
	0xba, 0xfa, 0x44, 0x09, 0xcd, 0x3c, 0x27, 0x80, 0x5b, 0x68, 0x67, 0xfc, 0x48, 0x24, 0xd2, 0xc8,
	0x5b, 0x4f, 0x94, 0xb0, 0x25, 0xb2, 0x02, 0x1a, 0xc1, 0xb5, 0x1c, 0xbd, 0xe4, 0x0c, 0x47, 0x75,
	0x5b, 0xfd, 0xae, 0x27, 0x2c, 0x7a, 0x95, 0x45, 0x6f, 0x51, 0x71, 0x13, 0x25, 0xec, 0x48, 0x91,
	0x68, 0xe3, 0x03, 0xc8, 0x59, 0x33, 0xbb, 0x59, 0x74, 0xe8, 0x14, 0x46, 0xa7, 0xe5, 0xbc, 0x59,
	0x3e, 0x5c, 0x35, 0x7c, 0x86, 0xfa, 0xd0, 0x3a, 0x4f, 0x9f, 0xd9, 0x97, 0x85, 0xe2, 0xba, 0x50,
	0xcc, 0x2a, 0x07, 0xb9, 0x04, 0xa4, 0x9f, 0x0c, 0xdd, 0xc3, 0x55, 0xdd, 0x50, 0x66, 0x9b, 0x85,
	0x4a, 0x2c, 0x34, 0x3a, 0x9b, 0xca, 0x75, 0xed, 0x9a, 0xc9, 0xac, 0xf7, 0x57, 0x05, 0x6d, 0xce,
	0x50, 0x0b, 0x9a, 0xcb, 0xe0, 0x6b, 0x30, 0xff, 0x19, 0x58, 0x0a, 0x02, 0x30, 0x46, 0xdf, 0x97,
	0x83, 0x59, 0x64, 0xa9, 0xa8, 0x03, 0x10, 0xcc, 0x17, 0xab, 0x32, 0xd6, 0x90, 0x05, 0xed, 0x71,
	0x38, 0x1a, 0x2c, 0x46, 0xe1, 0x6a, 0x31, 0x19, 0x04, 0x96, 0x8e, 0x5e, 0xc1, 0xf3, 0x7a, 0xa6,
	0x42, 0x1b, 0xe8, 0x0a, 0xcc, 0xd9, 0x28, 0x8a, 0x04, 0x67, 0xa0, 0x17, 0x60, 0xc9, 0xb0, 0x82,
	0x9a, 0xc8, 0x00, 0x6d, 0x1a, 0x58, 0x97, 0x79, 0xdf, 0x69, 0xb4, 0x8a, 0x96, 0xc3, 0x68, 0x11,
	0x4e, 0x83, 0xb1, 0x65, 0x0e, 0x9b, 0xd0, 0x28, 0xcc, 0xf4, 0x6e, 0xc1, 0x94, 0xab, 0x42, 0x2f,
	0xc1, 0x28, 0x2d, 0xaa, 0x8e, 0xee, 0x36, 0xc2, 0x32, 0xea, 0xbd, 0x87, 0x76, 0xdd, 0x67, 0x8d,
	0xd3, 0x1c, 0xdd, 0x35, 0x25, 0xf7, 0x0e, 0xe0, 0xbc, 0xc5, 0x1a, 0xa5, 0x3b, 0xba, 0xab, 0x4b,
	0xea, 0x1e, 0x8c, 0x2f, 0xc5, 0xdd, 0x23, 0x0f, 0x80, 0x55, 0x07, 0x29, 0xde, 0x59, 0x7d, 0x3e,
	0x79, 0xa7, 0x61, 0x8d, 0xe8, 0x7f, 0x02, 0xf4, 0xf9, 0x74, 0x38, 0x3c, 0x09, 0x79, 0x84, 0x8f,
	0x8f, 0x64, 0x83, 0xd1, 0x07, 0x30, 0xc7, 0x98, 0x97, 0x2d, 0x5b, 0x85, 0x5c, 0x04, 0xdd, 0x7a,
	0xd0, 0x53, 0x86, 0x1f, 0x7f, 0xdd, 0xed, 0x08, 0xff, 0x7d, 0x5a, 0x7b, 0x1b, 0x7a, 0xf0, 0xf7,
	0xa7, 0x35, 0xde, 0x26, 0xf4, 0x8f, 0xcf, 0x08, 0xc3, 0x09, 0x49, 0x71, 0xe6, 0xd7, 0x7f, 0xcd,
	0x1d, 0x5d, 0x6d, 0x12, 0x82, 0x53, 0xbe, 0x36, 0x8a, 0x53, 0xbc, 0xfb, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0xf8, 0x38, 0x00, 0xb0, 0xba, 0x03, 0x00, 0x00,
}
