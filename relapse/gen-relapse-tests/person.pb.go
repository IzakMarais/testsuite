// Code generated by protoc-gen-gogo.
// source: person.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	person.proto

It has these top-level messages:
	Person
	Address
*/
package main

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_protoc_gen_gogo_descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import compress_gzip "compress/gzip"
import bytes "bytes"
import io_ioutil "io/ioutil"

import strings "strings"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Person struct {
	Name             *string    `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Addresses        []*Address `protobuf:"bytes,2,rep,name=Addresses" json:"Addresses,omitempty"`
	Telephone        *string    `protobuf:"bytes,3,opt,name=Telephone" json:"Telephone,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptorPerson, []int{0} }

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Person) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Person) GetTelephone() string {
	if m != nil && m.Telephone != nil {
		return *m.Telephone
	}
	return ""
}

type Address struct {
	Number           *int64  `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
	Street           *string `protobuf:"bytes,2,opt,name=Street" json:"Street,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptorPerson, []int{1} }

func (m *Address) GetNumber() int64 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *Address) GetStreet() string {
	if m != nil && m.Street != nil {
		return *m.Street
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "main.Person")
	proto.RegisterType((*Address)(nil), "main.Address")
}
func (this *Person) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return PersonDescription()
}
func PersonDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3452 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x5a, 0x5b, 0x6c, 0x1c, 0x55,
		0x9a, 0xa6, 0xaf, 0xee, 0xfe, 0xbb, 0xdd, 0x2e, 0x1f, 0x1b, 0xa7, 0x63, 0x20, 0x71, 0x4c, 0x58,
		0x0c, 0x59, 0x1c, 0x14, 0x92, 0x90, 0x74, 0x16, 0xa2, 0xb6, 0xdd, 0x31, 0x8e, 0x7c, 0xe9, 0xad,
		0xb6, 0x21, 0xc0, 0x43, 0xe9, 0xb8, 0xea, 0x74, 0xbb, 0x92, 0xea, 0xaa, 0xde, 0xaa, 0xea, 0x24,
		0xce, 0x53, 0x10, 0x7b, 0x11, 0x42, 0x7b, 0x5f, 0x69, 0xb9, 0xc3, 0x22, 0xed, 0xb2, 0xcb, 0xde,
		0x86, 0xb9, 0x3d, 0xcc, 0xd3, 0xbc, 0x30, 0xf3, 0x36, 0x12, 0xef, 0xf3, 0x82, 0x84, 0x34, 0x37,
		0x66, 0x86, 0x99, 0x89, 0xc4, 0x48, 0x79, 0x19, 0x9d, 0x5b, 0x75, 0xf5, 0xc5, 0xa9, 0x36, 0x12,
		0x30, 0x4f, 0xf6, 0xf9, 0xcf, 0xff, 0x7d, 0x75, 0xce, 0xff, 0xff, 0xe7, 0xff, 0xff, 0x3a, 0xd5,
		0xf0, 0xc2, 0x49, 0x98, 0x69, 0x38, 0x4e, 0xc3, 0x22, 0xc7, 0x5b, 0xae, 0xe3, 0x3b, 0xdb, 0xed,
		0xfa, 0x71, 0x83, 0x78, 0xba, 0x6b, 0xb6, 0x7c, 0xc7, 0x9d, 0x67, 0x32, 0x34, 0xc6, 0x35, 0xe6,
		0xa5, 0xc6, 0xec, 0x1a, 0x8c, 0x5f, 0x30, 0x2d, 0xb2, 0x14, 0x28, 0xd6, 0x88, 0x8f, 0xce, 0x40,
		0xb2, 0x6e, 0x5a, 0xa4, 0x18, 0x9b, 0x49, 0xcc, 0xe5, 0x4e, 0x1c, 0x9d, 0xef, 0x01, 0xcd, 0x77,
		0x23, 0xaa, 0x54, 0xac, 0x32, 0xc4, 0xec, 0x27, 0x49, 0x98, 0x18, 0x30, 0x8b, 0x10, 0x24, 0x6d,
		0xdc, 0xa4, 0x8c, 0xb1, 0xb9, 0xac, 0xca, 0xfe, 0x47, 0x45, 0x18, 0x69, 0x61, 0xfd, 0x0a, 0x6e,
		0x90, 0x62, 0x9c, 0x89, 0xe5, 0x10, 0x1d, 0x02, 0x30, 0x48, 0x8b, 0xd8, 0x06, 0xb1, 0xf5, 0xdd,
		0x62, 0x62, 0x26, 0x31, 0x97, 0x55, 0x43, 0x12, 0x74, 0x0c, 0xc6, 0x5b, 0xed, 0x6d, 0xcb, 0xd4,
		0xb5, 0x90, 0x1a, 0xcc, 0x24, 0xe6, 0x52, 0xaa, 0xc2, 0x27, 0x96, 0x3a, 0xca, 0x0f, 0xc2, 0xd8,
		0x35, 0x82, 0xaf, 0x84, 0x55, 0x73, 0x4c, 0xb5, 0x40, 0xc5, 0x21, 0xc5, 0x45, 0xc8, 0x37, 0x89,
		0xe7, 0xe1, 0x06, 0xd1, 0xfc, 0xdd, 0x16, 0x29, 0x26, 0xd9, 0xee, 0x67, 0xfa, 0x76, 0xdf, 0xbb,
		0xf3, 0x9c, 0x40, 0x6d, 0xee, 0xb6, 0x08, 0x2a, 0x43, 0x96, 0xd8, 0xed, 0x26, 0x67, 0x48, 0xed,
		0x61, 0xbf, 0x8a, 0xdd, 0x6e, 0xf6, 0xb2, 0x64, 0x28, 0x4c, 0x50, 0x8c, 0x78, 0xc4, 0xbd, 0x6a,
		0xea, 0xa4, 0x98, 0x66, 0x04, 0x0f, 0xf6, 0x11, 0xd4, 0xf8, 0x7c, 0x2f, 0x87, 0xc4, 0xa1, 0x45,
		0xc8, 0x92, 0xeb, 0x3e, 0xb1, 0x3d, 0xd3, 0xb1, 0x8b, 0x23, 0x8c, 0xe4, 0x81, 0x01, 0x5e, 0x24,
		0x96, 0xd1, 0x4b, 0xd1, 0xc1, 0xa1, 0xd3, 0x30, 0xe2, 0xb4, 0x7c, 0xd3, 0xb1, 0xbd, 0x62, 0x66,
		0x26, 0x36, 0x97, 0x3b, 0x71, 0xef, 0xc0, 0x40, 0xd8, 0xe0, 0x3a, 0xaa, 0x54, 0x46, 0x2b, 0xa0,
		0x78, 0x4e, 0xdb, 0xd5, 0x89, 0xa6, 0x3b, 0x06, 0xd1, 0x4c, 0xbb, 0xee, 0x14, 0xb3, 0x8c, 0xe0,
		0x70, 0xff, 0x46, 0x98, 0xe2, 0xa2, 0x63, 0x90, 0x15, 0xbb, 0xee, 0xa8, 0x05, 0xaf, 0x6b, 0x8c,
		0xa6, 0x20, 0xed, 0xed, 0xda, 0x3e, 0xbe, 0x5e, 0xcc, 0xb3, 0x08, 0x11, 0xa3, 0xd9, 0xcf, 0x53,
		0x30, 0x36, 0x4c, 0x88, 0x9d, 0x83, 0x54, 0x9d, 0xee, 0xb2, 0x18, 0xdf, 0x8f, 0x0d, 0x38, 0xa6,
		0xdb, 0x88, 0xe9, 0x2f, 0x68, 0xc4, 0x32, 0xe4, 0x6c, 0xe2, 0xf9, 0xc4, 0xe0, 0x11, 0x91, 0x18,
		0x32, 0xa6, 0x80, 0x83, 0xfa, 0x43, 0x2a, 0xf9, 0x85, 0x42, 0xea, 0x12, 0x8c, 0x05, 0x4b, 0xd2,
		0x5c, 0x6c, 0x37, 0x64, 0x6c, 0x1e, 0x8f, 0x5a, 0xc9, 0x7c, 0x45, 0xe2, 0x54, 0x0a, 0x53, 0x0b,
		0xa4, 0x6b, 0x8c, 0x96, 0x00, 0x1c, 0x9b, 0x38, 0x75, 0xcd, 0x20, 0xba, 0x55, 0xcc, 0xec, 0x61,
		0xa5, 0x0d, 0xaa, 0xd2, 0x67, 0x25, 0x87, 0x4b, 0x75, 0x0b, 0x9d, 0xed, 0x84, 0xda, 0xc8, 0x1e,
		0x91, 0xb2, 0xc6, 0x0f, 0x59, 0x5f, 0xb4, 0x6d, 0x41, 0xc1, 0x25, 0x34, 0xee, 0x89, 0x21, 0x76,
		0x96, 0x65, 0x8b, 0x98, 0x8f, 0xdc, 0x99, 0x2a, 0x60, 0x7c, 0x63, 0xa3, 0x6e, 0x78, 0x88, 0xee,
		0x87, 0x40, 0xa0, 0xb1, 0xb0, 0x02, 0x96, 0x85, 0xf2, 0x52, 0xb8, 0x8e, 0x9b, 0x64, 0xfa, 0x0c,
		0x14, 0xba, 0xcd, 0x83, 0x26, 0x21, 0xe5, 0xf9, 0xd8, 0xf5, 0x59, 0x14, 0xa6, 0x54, 0x3e, 0x40,
		0x0a, 0x24, 0x88, 0x6d, 0xb0, 0x2c, 0x97, 0x52, 0xe9, 0xbf, 0xd3, 0x8f, 0xc3, 0x68, 0xd7, 0xe3,
		0x87, 0x05, 0xce, 0xbe, 0x92, 0x86, 0xc9, 0x41, 0x31, 0x37, 0x30, 0xfc, 0xa7, 0x20, 0x6d, 0xb7,
		0x9b, 0xdb, 0xc4, 0x2d, 0x26, 0x18, 0x83, 0x18, 0xa1, 0x32, 0xa4, 0x2c, 0xbc, 0x4d, 0xac, 0x62,
		0x72, 0x26, 0x36, 0x57, 0x38, 0x71, 0x6c, 0xa8, 0xa8, 0x9e, 0x5f, 0xa5, 0x10, 0x95, 0x23, 0xd1,
		0x93, 0x90, 0x14, 0x29, 0x8e, 0x32, 0x3c, 0x3c, 0x1c, 0x03, 0x8d, 0x45, 0x95, 0xe1, 0xd0, 0x3d,
		0x90, 0xa5, 0x7f, 0xb9, 0x6d, 0xd3, 0x6c, 0xcd, 0x19, 0x2a, 0xa0, 0x76, 0x45, 0xd3, 0x90, 0x61,
		0x61, 0x66, 0x10, 0x59, 0x1a, 0x82, 0x31, 0x75, 0x8c, 0x41, 0xea, 0xb8, 0x6d, 0xf9, 0xda, 0x55,
		0x6c, 0xb5, 0x09, 0x0b, 0x98, 0xac, 0x9a, 0x17, 0xc2, 0xa7, 0xa9, 0x0c, 0x1d, 0x86, 0x1c, 0x8f,
		0x4a, 0xd3, 0x36, 0xc8, 0x75, 0x96, 0x7d, 0x52, 0x2a, 0x0f, 0xd4, 0x15, 0x2a, 0xa1, 0x8f, 0xbf,
		0xec, 0x39, 0xb6, 0x74, 0x2d, 0x7b, 0x04, 0x15, 0xb0, 0xc7, 0x3f, 0xde, 0x9b, 0xf8, 0xee, 0x1b,
		0xbc, 0xbd, 0xde, 0x58, 0x9c, 0xfd, 0x6e, 0x1c, 0x92, 0xec, 0xbc, 0x8d, 0x41, 0x6e, 0xf3, 0xd9,
		0x6a, 0x45, 0x5b, 0xda, 0xd8, 0x5a, 0x58, 0xad, 0x28, 0x31, 0x54, 0x00, 0x60, 0x82, 0x0b, 0xab,
		0x1b, 0xe5, 0x4d, 0x25, 0x1e, 0x8c, 0x57, 0xd6, 0x37, 0x4f, 0x9f, 0x54, 0x12, 0x01, 0x60, 0x8b,
		0x0b, 0x92, 0x61, 0x85, 0xc7, 0x4e, 0x28, 0x29, 0xa4, 0x40, 0x9e, 0x13, 0xac, 0x5c, 0xaa, 0x2c,
		0x9d, 0x3e, 0xa9, 0xa4, 0xbb, 0x25, 0x8f, 0x9d, 0x50, 0x46, 0xd0, 0x28, 0x64, 0x99, 0x64, 0x61,
		0x63, 0x63, 0x55, 0xc9, 0x04, 0x9c, 0xb5, 0x4d, 0x75, 0x65, 0x7d, 0x59, 0xc9, 0x06, 0x9c, 0xcb,
		0xea, 0xc6, 0x56, 0x55, 0x81, 0x80, 0x61, 0xad, 0x52, 0xab, 0x95, 0x97, 0x2b, 0x4a, 0x2e, 0xd0,
		0x58, 0x78, 0x76, 0xb3, 0x52, 0x53, 0xf2, 0x5d, 0xcb, 0x7a, 0xec, 0x84, 0x32, 0x1a, 0x3c, 0xa2,
		0xb2, 0xbe, 0xb5, 0xa6, 0x14, 0xd0, 0x38, 0x8c, 0xf2, 0x47, 0xc8, 0x45, 0x8c, 0xf5, 0x88, 0x4e,
		0x9f, 0x54, 0x94, 0xce, 0x42, 0x38, 0xcb, 0x78, 0x97, 0xe0, 0xf4, 0x49, 0x05, 0xcd, 0x2e, 0x42,
		0x8a, 0x45, 0x17, 0x42, 0x50, 0x58, 0x2d, 0x2f, 0x54, 0x56, 0xb5, 0x8d, 0xea, 0xe6, 0xca, 0xc6,
		0x7a, 0x79, 0x55, 0x89, 0x75, 0x64, 0x6a, 0xe5, 0xcf, 0xb7, 0x56, 0xd4, 0xca, 0x92, 0x12, 0x0f,
		0xcb, 0xaa, 0x95, 0xf2, 0x66, 0x65, 0x49, 0x49, 0xcc, 0xea, 0x30, 0x39, 0x28, 0xcf, 0x0c, 0x3c,
		0x19, 0x21, 0x17, 0xc7, 0xf7, 0x70, 0x31, 0xe3, 0xea, 0x73, 0xf1, 0xbb, 0x31, 0x98, 0x18, 0x90,
		0x6b, 0x07, 0x3e, 0xe4, 0x3c, 0xa4, 0x78, 0x88, 0xf2, 0xea, 0xf3, 0xd0, 0xc0, 0xa4, 0xcd, 0x02,
		0xb6, 0xaf, 0x02, 0x31, 0x5c, 0xb8, 0x02, 0x27, 0xf6, 0xa8, 0xc0, 0x94, 0xa2, 0x6f, 0x91, 0x2f,
		0xc6, 0xa0, 0xb8, 0x17, 0x77, 0x44, 0xa2, 0x88, 0x77, 0x25, 0x8a, 0x73, 0xbd, 0x0b, 0x38, 0xb2,
		0xf7, 0x1e, 0xfa, 0x56, 0xf1, 0x5e, 0x0c, 0xa6, 0x06, 0x37, 0x2a, 0x03, 0xd7, 0xf0, 0x24, 0xa4,
		0x9b, 0xc4, 0xdf, 0x71, 0x64, 0xb1, 0xfe, 0x93, 0x01, 0x25, 0x80, 0x4e, 0xf7, 0xda, 0x4a, 0xa0,
		0xc2, 0x35, 0x24, 0xb1, 0x57, 0xb7, 0xc1, 0x57, 0xd3, 0xb7, 0xd2, 0x97, 0xe2, 0x70, 0xf7, 0x40,
		0xf2, 0x81, 0x0b, 0xbd, 0x0f, 0xc0, 0xb4, 0x5b, 0x6d, 0x9f, 0x17, 0x64, 0x9e, 0x9f, 0xb2, 0x4c,
		0xc2, 0xce, 0x3e, 0xcd, 0x3d, 0x6d, 0x3f, 0x98, 0x4f, 0xb0, 0x79, 0xe0, 0x22, 0xa6, 0x70, 0xa6,
		0xb3, 0xd0, 0x24, 0x5b, 0xe8, 0xa1, 0x3d, 0x76, 0xda, 0x57, 0xeb, 0x1e, 0x05, 0x45, 0xb7, 0x4c,
		0x62, 0xfb, 0x9a, 0xe7, 0xbb, 0x04, 0x37, 0x4d, 0xbb, 0xc1, 0x12, 0x70, 0xa6, 0x94, 0xaa, 0x63,
		0xcb, 0x23, 0xea, 0x18, 0x9f, 0xae, 0xc9, 0x59, 0x8a, 0x60, 0x55, 0xc6, 0x0d, 0x21, 0xd2, 0x5d,
		0x08, 0x3e, 0x1d, 0x20, 0x66, 0x5f, 0x1e, 0x81, 0x5c, 0xa8, 0xad, 0x43, 0x47, 0x20, 0x7f, 0x19,
		0x5f, 0xc5, 0x9a, 0x6c, 0xd5, 0xb9, 0x25, 0x72, 0x54, 0x56, 0x15, 0xed, 0xfa, 0xa3, 0x30, 0xc9,
		0x54, 0x9c, 0xb6, 0x4f, 0x5c, 0x4d, 0xb7, 0xb0, 0xe7, 0x31, 0xa3, 0x65, 0x98, 0x2a, 0xa2, 0x73,
		0x1b, 0x74, 0x6a, 0x51, 0xce, 0xa0, 0x53, 0x30, 0xc1, 0x10, 0xcd, 0xb6, 0xe5, 0x9b, 0x2d, 0x8b,
		0x68, 0xf4, 0xe5, 0xc1, 0x63, 0x89, 0x38, 0x58, 0xd9, 0x38, 0xd5, 0x58, 0x13, 0x0a, 0x74, 0x45,
		0x1e, 0x5a, 0x86, 0xfb, 0x18, 0xac, 0x41, 0x6c, 0xe2, 0x62, 0x9f, 0x68, 0xe4, 0x2f, 0xda, 0xd8,
		0xf2, 0x34, 0x6c, 0x1b, 0xda, 0x0e, 0xf6, 0x76, 0x8a, 0x93, 0x61, 0x82, 0x83, 0x54, 0x77, 0x59,
		0xa8, 0x56, 0x98, 0x66, 0xd9, 0x36, 0x9e, 0xc2, 0xde, 0x0e, 0x2a, 0xc1, 0x14, 0x23, 0xf2, 0x7c,
		0xd7, 0xb4, 0x1b, 0x9a, 0xbe, 0x43, 0xf4, 0x2b, 0x5a, 0xdb, 0xaf, 0x9f, 0x29, 0xde, 0x13, 0x66,
		0x60, 0x8b, 0xac, 0x31, 0x9d, 0x45, 0xaa, 0xb2, 0xe5, 0xd7, 0xcf, 0xa0, 0x1a, 0xe4, 0xa9, 0x3f,
		0x9a, 0xe6, 0x0d, 0xa2, 0xd5, 0x1d, 0x97, 0x15, 0x97, 0xc2, 0x80, 0xc3, 0x1d, 0x32, 0xe2, 0xfc,
		0x86, 0x00, 0xac, 0x39, 0x06, 0x29, 0xa5, 0x6a, 0xd5, 0x4a, 0x65, 0x49, 0xcd, 0x49, 0x96, 0x0b,
		0x8e, 0x4b, 0x63, 0xaa, 0xe1, 0x04, 0x36, 0xce, 0xf1, 0x98, 0x6a, 0x38, 0xd2, 0xc2, 0xa7, 0x60,
		0x42, 0xd7, 0xf9, 0xb6, 0x4d, 0x5d, 0x13, 0x5d, 0xbe, 0x57, 0x54, 0xba, 0xec, 0xa5, 0xeb, 0xcb,
		0x5c, 0x41, 0x84, 0xb9, 0x87, 0xce, 0xc2, 0xdd, 0x1d, 0x7b, 0x85, 0x81, 0xe3, 0x7d, 0xbb, 0xec,
		0x85, 0x9e, 0x82, 0x89, 0xd6, 0x6e, 0x3f, 0x10, 0x75, 0x3d, 0xb1, 0xb5, 0xdb, 0x0b, 0x7b, 0x80,
		0xbd, 0xb9, 0xb9, 0x44, 0xc7, 0x3e, 0x31, 0x8a, 0x07, 0xc2, 0xda, 0xa1, 0x09, 0x74, 0x1c, 0x14,
		0x5d, 0xd7, 0x88, 0x8d, 0xb7, 0x2d, 0xa2, 0x61, 0x97, 0xd8, 0xd8, 0x2b, 0x1e, 0x0e, 0x2b, 0x17,
		0x74, 0xbd, 0xc2, 0x66, 0xcb, 0x6c, 0x12, 0x3d, 0x0c, 0xe3, 0xce, 0xf6, 0x65, 0x9d, 0x07, 0x97,
		0xd6, 0x72, 0x49, 0xdd, 0xbc, 0x5e, 0x3c, 0xca, 0xcc, 0x34, 0x46, 0x27, 0x58, 0x68, 0x55, 0x99,
		0x18, 0x3d, 0x04, 0x8a, 0xee, 0xed, 0x60, 0xb7, 0xc5, 0xaa, 0xbb, 0xd7, 0xc2, 0x3a, 0x29, 0x3e,
		0xc0, 0x55, 0xb9, 0x7c, 0x5d, 0x8a, 0xd1, 0x25, 0x98, 0x6c, 0xdb, 0xa6, 0xed, 0x13, 0xb7, 0xe5,
		0x12, 0xda, 0xa4, 0xf3, 0x93, 0x56, 0xfc, 0xc9, 0xc8, 0x1e, 0x6d, 0xf6, 0x56, 0x58, 0x9b, 0x7b,
		0x57, 0x9d, 0x68, 0xf7, 0x0b, 0x67, 0x4b, 0x90, 0x0f, 0x3b, 0x1d, 0x65, 0x81, 0xbb, 0x5d, 0x89,
		0xd1, 0x1a, 0xba, 0xb8, 0xb1, 0x44, 0xab, 0xdf, 0x73, 0x15, 0x25, 0x4e, 0xab, 0xf0, 0xea, 0xca,
		0x66, 0x45, 0x53, 0xb7, 0xd6, 0x37, 0x57, 0xd6, 0x2a, 0x4a, 0xe2, 0xe1, 0x6c, 0xe6, 0xa7, 0x23,
		0xca, 0xcd, 0x9b, 0x37, 0x6f, 0xc6, 0x67, 0x3f, 0x8c, 0x43, 0xa1, 0xbb, 0xf3, 0x45, 0x7f, 0x06,
		0x07, 0xe4, 0x6b, 0xaa, 0x47, 0x7c, 0xed, 0x9a, 0xe9, 0xb2, 0x38, 0x6c, 0x62, 0xde, 0x3b, 0x06,
		0x26, 0x9c, 0x14, 0x5a, 0x35, 0xe2, 0x3f, 0x63, 0xba, 0x34, 0xca, 0x9a, 0xd8, 0x47, 0xab, 0x70,
		0xd8, 0x76, 0x34, 0xcf, 0xc7, 0xb6, 0x81, 0x5d, 0x43, 0xeb, 0x5c, 0x10, 0x68, 0x58, 0xd7, 0x89,
		0xe7, 0x39, 0xbc, 0x04, 0x04, 0x2c, 0xf7, 0xda, 0x4e, 0x4d, 0x28, 0x77, 0x72, 0x63, 0x59, 0xa8,
		0xf6, 0xb8, 0x3b, 0xb1, 0x97, 0xbb, 0xef, 0x81, 0x6c, 0x13, 0xb7, 0x34, 0x62, 0xfb, 0xee, 0x2e,
		0xeb, 0xd7, 0x32, 0x6a, 0xa6, 0x89, 0x5b, 0x15, 0x3a, 0xfe, 0xf2, 0x7c, 0x10, 0xb6, 0xe3, 0x8f,
		0x13, 0x90, 0x0f, 0xf7, 0x6c, 0xb4, 0x05, 0xd6, 0x59, 0x7e, 0x8e, 0xb1, 0xe3, 0x7b, 0xff, 0x1d,
		0x3b, 0xbc, 0xf9, 0x45, 0x9a, 0xb8, 0x4b, 0x69, 0xde, 0x49, 0xa9, 0x1c, 0x49, 0x8b, 0x26, 0x3d,
		0xb0, 0x84, 0xf7, 0xe7, 0x19, 0x55, 0x8c, 0xd0, 0x32, 0xa4, 0x2f, 0x7b, 0x8c, 0x3b, 0xcd, 0xb8,
		0x8f, 0xde, 0x99, 0xfb, 0x62, 0x8d, 0x91, 0x67, 0x2f, 0xd6, 0xb4, 0xf5, 0x0d, 0x75, 0xad, 0xbc,
		0xaa, 0x0a, 0x38, 0x3a, 0x08, 0x49, 0x0b, 0xdf, 0xd8, 0xed, 0x4e, 0xf1, 0x4c, 0x34, 0xac, 0xe1,
		0x0f, 0x42, 0xf2, 0x1a, 0xc1, 0x57, 0xba, 0x13, 0x2b, 0x13, 0x7d, 0x89, 0xa1, 0x7f, 0x1c, 0x52,
		0xcc, 0x5e, 0x08, 0x40, 0x58, 0x4c, 0xb9, 0x0b, 0x65, 0x20, 0xb9, 0xb8, 0xa1, 0xd2, 0xf0, 0x57,
		0x20, 0xcf, 0xa5, 0x5a, 0x75, 0xa5, 0xb2, 0x58, 0x51, 0xe2, 0xb3, 0xa7, 0x20, 0xcd, 0x8d, 0x40,
		0x8f, 0x46, 0x60, 0x06, 0xe5, 0x2e, 0x31, 0x14, 0x1c, 0x31, 0x39, 0xbb, 0xb5, 0xb6, 0x50, 0x51,
		0x95, 0x78, 0xd8, 0xbd, 0x1e, 0xe4, 0xc3, 0xed, 0xda, 0x57, 0x13, 0x53, 0xdf, 0x8b, 0x41, 0x2e,
		0xd4, 0x7e, 0xd1, 0xc2, 0x8f, 0x2d, 0xcb, 0xb9, 0xa6, 0x61, 0xcb, 0xc4, 0x9e, 0x08, 0x0a, 0x60,
		0xa2, 0x32, 0x95, 0x0c, 0xeb, 0xb4, 0xaf, 0x64, 0xf1, 0x6f, 0xc5, 0x40, 0xe9, 0x6d, 0xdd, 0x7a,
		0x16, 0x18, 0xfb, 0x5a, 0x17, 0xf8, 0x46, 0x0c, 0x0a, 0xdd, 0xfd, 0x5a, 0xcf, 0xf2, 0x8e, 0x7c,
		0xad, 0xcb, 0x7b, 0x3d, 0x06, 0xa3, 0x5d, 0x5d, 0xda, 0x1f, 0xd5, 0xea, 0x5e, 0x4b, 0xc0, 0xc4,
		0x00, 0x1c, 0x2a, 0x8b, 0x76, 0x96, 0x77, 0xd8, 0x8f, 0x0c, 0xf3, 0xac, 0x79, 0x5a, 0x2d, 0xab,
		0xd8, 0xf5, 0x45, 0xf7, 0xfb, 0x10, 0x28, 0xa6, 0x41, 0x6c, 0xdf, 0xac, 0x9b, 0xc4, 0x15, 0xaf,
		0xe0, 0xbc, 0xc7, 0x1d, 0xeb, 0xc8, 0xf9, 0x5b, 0xf8, 0x9f, 0x02, 0x6a, 0x39, 0x9e, 0xe9, 0x9b,
		0x57, 0x89, 0x66, 0xda, 0xf2, 0x7d, 0x9d, 0xf6, 0xbc, 0x49, 0x55, 0x91, 0x33, 0x2b, 0xb6, 0x1f,
		0x68, 0xdb, 0xa4, 0x81, 0x7b, 0xb4, 0x69, 0xee, 0x4b, 0xa8, 0x8a, 0x9c, 0x09, 0xb4, 0x8f, 0x40,
		0xde, 0x70, 0xda, 0xb4, 0x7d, 0xe0, 0x7a, 0x34, 0xd5, 0xc6, 0xd4, 0x1c, 0x97, 0x05, 0x2a, 0xa2,
		0xbf, 0xeb, 0x5c, 0x14, 0xe4, 0xd5, 0x1c, 0x97, 0x71, 0x95, 0x07, 0x61, 0x0c, 0x37, 0x1a, 0x2e,
		0x25, 0x97, 0x44, 0xbc, 0x69, 0x2d, 0x04, 0x62, 0xa6, 0x38, 0x7d, 0x11, 0x32, 0xd2, 0x0e, 0xb4,
		0x9a, 0x51, 0x4b, 0x68, 0x2d, 0x7e, 0x5d, 0x13, 0x9f, 0xcb, 0xaa, 0x19, 0x5b, 0x4e, 0x1e, 0x81,
		0xbc, 0xe9, 0x69, 0x9d, 0x7b, 0xc3, 0xf8, 0x4c, 0x7c, 0x2e, 0xa3, 0xe6, 0x4c, 0x2f, 0xb8, 0x28,
		0x9a, 0x7d, 0x2f, 0x0e, 0x85, 0xee, 0x7b, 0x4f, 0xb4, 0x04, 0x19, 0xcb, 0xd1, 0x31, 0x0b, 0x04,
		0x7e, 0xe9, 0x3e, 0x17, 0x71, 0x55, 0x3a, 0xbf, 0x2a, 0xf4, 0xd5, 0x00, 0x39, 0xfd, 0xa3, 0x18,
		0x64, 0xa4, 0x18, 0x4d, 0x41, 0xb2, 0x85, 0xfd, 0x1d, 0x46, 0x97, 0x5a, 0x88, 0x2b, 0x31, 0x95,
		0x8d, 0xa9, 0xdc, 0x6b, 0x61, 0x9b, 0x85, 0x80, 0x90, 0xd3, 0x31, 0xf5, 0xab, 0x45, 0xb0, 0xc1,
		0xda, 0x61, 0xa7, 0xd9, 0x24, 0xb6, 0xef, 0x49, 0xbf, 0x0a, 0xf9, 0xa2, 0x10, 0xa3, 0x63, 0x30,
		0xee, 0xbb, 0xd8, 0xb4, 0xba, 0x74, 0x93, 0x4c, 0x57, 0x91, 0x13, 0x81, 0x72, 0x09, 0x0e, 0x4a,
		0x5e, 0x83, 0xf8, 0x58, 0xdf, 0x21, 0x46, 0x07, 0x94, 0x66, 0x97, 0x6a, 0x07, 0x84, 0xc2, 0x92,
		0x98, 0x97, 0xd8, 0xd9, 0x8f, 0x62, 0x30, 0x2e, 0x1b, 0x78, 0x23, 0x30, 0xd6, 0x1a, 0x00, 0xb6,
		0x6d, 0xc7, 0x0f, 0x9b, 0xab, 0x3f, 0x94, 0xfb, 0x70, 0xf3, 0xe5, 0x00, 0xa4, 0x86, 0x08, 0xa6,
		0x9b, 0x00, 0x9d, 0x99, 0x3d, 0xcd, 0x76, 0x18, 0x72, 0xe2, 0x52, 0x9b, 0x7d, 0x19, 0xe1, 0x6f,
		0x7d, 0xc0, 0x45, 0xb4, 0xd3, 0x47, 0x93, 0x90, 0xda, 0x26, 0x0d, 0xd3, 0x16, 0x57, 0x6d, 0x7c,
		0x20, 0x2f, 0xf0, 0x92, 0xc1, 0x05, 0xde, 0xc2, 0xf3, 0x30, 0xa1, 0x3b, 0xcd, 0xde, 0xe5, 0x2e,
		0x28, 0x3d, 0x6f, 0x9e, 0xde, 0x53, 0xb1, 0xe7, 0xa0, 0xd3, 0x9d, 0xbd, 0x13, 0x8b, 0xbd, 0x1b,
		0x4f, 0x2c, 0x57, 0x17, 0xde, 0x8f, 0x4f, 0x2f, 0x73, 0x68, 0x55, 0xee, 0x54, 0x25, 0x75, 0x8b,
		0xe8, 0x74, 0xf5, 0xf0, 0xf6, 0x51, 0x78, 0xa4, 0x61, 0xfa, 0x3b, 0xed, 0xed, 0x79, 0xdd, 0x69,
		0x1e, 0x6f, 0x38, 0x0d, 0xa7, 0xf3, 0x31, 0x88, 0x8e, 0xd8, 0x80, 0xfd, 0x27, 0x3e, 0x08, 0x65,
		0x03, 0xe9, 0x74, 0xe4, 0xd7, 0xa3, 0xd2, 0x3a, 0x4c, 0x08, 0x65, 0x8d, 0xdd, 0x48, 0xf3, 0x3e,
		0x1c, 0xdd, 0xf1, 0x56, 0xa2, 0xf8, 0xc1, 0x27, 0xac, 0xd2, 0xa9, 0xe3, 0x02, 0x4a, 0xe7, 0x78,
		0xa7, 0x5e, 0x52, 0xe1, 0xee, 0x2e, 0x3e, 0x7e, 0x34, 0x89, 0x1b, 0xc1, 0xf8, 0xa1, 0x60, 0x9c,
		0x08, 0x31, 0xd6, 0x04, 0xb4, 0xb4, 0x08, 0xa3, 0xfb, 0xe1, 0xfa, 0x81, 0xe0, 0xca, 0x93, 0x30,
		0xc9, 0x32, 0x8c, 0x31, 0x12, 0xbd, 0xed, 0xf9, 0x4e, 0x93, 0xe5, 0xbd, 0x3b, 0xd3, 0xfc, 0xf0,
		0x13, 0x7e, 0x56, 0x0a, 0x14, 0xb6, 0x18, 0xa0, 0x4a, 0x4f, 0xc3, 0x24, 0x95, 0xb0, 0xd4, 0x12,
		0x66, 0x8b, 0xbe, 0x47, 0x29, 0x7e, 0xf4, 0x22, 0x3f, 0x52, 0x13, 0x01, 0x41, 0x88, 0x37, 0xe4,
		0x89, 0x06, 0xf1, 0x7d, 0xe2, 0x7a, 0x1a, 0xb6, 0x2c, 0x74, 0xc7, 0x2f, 0x34, 0xc5, 0x57, 0x3f,
		0xed, 0xf6, 0xc4, 0x32, 0x47, 0x96, 0x2d, 0xab, 0xb4, 0x05, 0x07, 0x06, 0x78, 0x76, 0x08, 0xce,
		0xd7, 0x04, 0xe7, 0x64, 0x9f, 0x77, 0x29, 0x6d, 0x15, 0xa4, 0x3c, 0xf0, 0xc7, 0x10, 0x9c, 0xaf,
		0x0b, 0x4e, 0x24, 0xb0, 0xd2, 0x2d, 0x94, 0xf1, 0x22, 0x8c, 0x5f, 0x25, 0xee, 0xb6, 0xe3, 0x89,
		0x97, 0xff, 0x21, 0xe8, 0xde, 0x10, 0x74, 0x63, 0x02, 0xc8, 0xae, 0x02, 0x28, 0xd7, 0x59, 0xc8,
		0xd4, 0xb1, 0x4e, 0x86, 0xa0, 0x78, 0x53, 0x50, 0x8c, 0x50, 0x7d, 0x0a, 0x2d, 0x43, 0xbe, 0xe1,
		0x88, 0xea, 0x12, 0x0d, 0x7f, 0x4b, 0xc0, 0x73, 0x12, 0x23, 0x28, 0x5a, 0x4e, 0xab, 0x6d, 0xd1,
		0xd2, 0x13, 0x4d, 0xf1, 0xb6, 0xa4, 0x90, 0x18, 0x41, 0xb1, 0x0f, 0xb3, 0xbe, 0x23, 0x29, 0xbc,
		0x90, 0x3d, 0xcf, 0x43, 0xce, 0xb1, 0xad, 0x5d, 0xc7, 0x1e, 0x66, 0x11, 0xff, 0x26, 0x18, 0x40,
		0x40, 0x28, 0xc1, 0x39, 0xc8, 0x0e, 0xeb, 0x88, 0x7f, 0x17, 0xf0, 0x0c, 0x91, 0x1e, 0x58, 0x86,
		0x31, 0x99, 0x64, 0x4c, 0xc7, 0x1e, 0x82, 0xe2, 0x3f, 0x04, 0x45, 0x21, 0x04, 0x13, 0xdb, 0xf0,
		0x89, 0xe7, 0x37, 0xc8, 0x30, 0x24, 0xef, 0xc9, 0x6d, 0x08, 0x88, 0x30, 0xe5, 0x36, 0xb1, 0xf5,
		0x9d, 0xe1, 0x18, 0xfe, 0x53, 0x9a, 0x52, 0x62, 0x28, 0xc5, 0x22, 0x8c, 0x36, 0xb1, 0xeb, 0xed,
		0x60, 0x6b, 0x28, 0x77, 0xfc, 0x97, 0xe0, 0xc8, 0x07, 0x20, 0x61, 0x91, 0xb6, 0xbd, 0x1f, 0x9a,
		0xf7, 0xa5, 0x45, 0x42, 0x30, 0x71, 0xf4, 0x3c, 0x9f, 0xdd, 0xaf, 0xec, 0x87, 0xed, 0xbf, 0xe5,
		0xd1, 0xe3, 0xd8, 0xb5, 0x30, 0xe3, 0x39, 0xc8, 0x7a, 0xe6, 0x8d, 0xa1, 0x68, 0xfe, 0x47, 0x7a,
		0x9a, 0x01, 0x28, 0xf8, 0x59, 0x38, 0x38, 0x30, 0xd5, 0x0f, 0x41, 0xf6, 0xbf, 0x82, 0x6c, 0x6a,
		0x40, 0xba, 0x17, 0x29, 0x61, 0xbf, 0x94, 0xff, 0x27, 0x53, 0x02, 0xe9, 0xe1, 0xaa, 0xd2, 0xee,
		0xdc, 0xc3, 0xf5, 0xfd, 0x59, 0xed, 0xff, 0xa5, 0xd5, 0x38, 0xb6, 0xcb, 0x6a, 0x9b, 0x30, 0x25,
		0x18, 0xf7, 0xe7, 0xd7, 0x6f, 0xc8, 0xc4, 0xca, 0xd1, 0x5b, 0xdd, 0xde, 0x7d, 0x1e, 0xa6, 0x03,
		0x73, 0xca, 0xc6, 0xd2, 0xd3, 0x9a, 0xb8, 0x35, 0x04, 0xf3, 0x07, 0x82, 0x59, 0x66, 0xfc, 0xa0,
		0x33, 0xf5, 0xd6, 0x70, 0x8b, 0x92, 0x5f, 0x82, 0xa2, 0x24, 0x6f, 0xdb, 0x2e, 0xd1, 0x9d, 0x86,
		0x6d, 0xde, 0x20, 0xc6, 0x10, 0xd4, 0xdf, 0xec, 0x71, 0xd5, 0x56, 0x08, 0x4e, 0x99, 0x57, 0x40,
		0x09, 0xfa, 0x0d, 0xcd, 0x6c, 0xb6, 0x1c, 0xd7, 0x8f, 0x60, 0xfc, 0x96, 0xf4, 0x54, 0x80, 0x5b,
		0x61, 0xb0, 0x52, 0x05, 0x0a, 0x6c, 0x38, 0x6c, 0x48, 0x7e, 0x5b, 0x10, 0x8d, 0x76, 0x50, 0x22,
		0x71, 0xe8, 0x4e, 0xb3, 0x85, 0xdd, 0x61, 0xf2, 0xdf, 0x77, 0x64, 0xe2, 0x10, 0x10, 0x1e, 0x7d,
		0x63, 0x3d, 0x95, 0x18, 0x45, 0x7d, 0xbc, 0x2e, 0xbe, 0x70, 0x4b, 0x9c, 0xd9, 0xee, 0x42, 0x5c,
		0x5a, 0xa5, 0xe6, 0xe9, 0x2e, 0x97, 0xd1, 0x64, 0x2f, 0xde, 0x0a, 0x2c, 0xd4, 0x55, 0x2d, 0x4b,
		0x17, 0x60, 0xb4, 0xab, 0x54, 0x46, 0x53, 0xfd, 0xa5, 0xa0, 0xca, 0x87, 0x2b, 0x65, 0xe9, 0x14,
		0x24, 0x69, 0xd9, 0x8b, 0x86, 0xff, 0x95, 0x80, 0x33, 0xf5, 0xd2, 0x13, 0x90, 0x91, 0xe5, 0x2e,
		0x1a, 0xfa, 0xd7, 0x02, 0x1a, 0x40, 0x28, 0x5c, 0x96, 0xba, 0x68, 0xf8, 0xdf, 0x48, 0xb8, 0x84,
		0x50, 0xf8, 0xf0, 0x26, 0xfc, 0xfe, 0xcb, 0x49, 0x91, 0xae, 0xa4, 0xed, 0xce, 0xc1, 0x88, 0xa8,
		0x71, 0xd1, 0xe8, 0x97, 0xc4, 0xc3, 0x25, 0xa2, 0xf4, 0x38, 0xa4, 0x86, 0x34, 0xf8, 0xdf, 0x0a,
		0x28, 0xd7, 0x2f, 0x2d, 0x42, 0x2e, 0x54, 0xd7, 0xa2, 0xe1, 0x7f, 0x27, 0xe0, 0x61, 0x14, 0x5d,
		0xba, 0xa8, 0x6b, 0xd1, 0x04, 0x7f, 0x2f, 0x97, 0x2e, 0x10, 0xd4, 0x6c, 0xb2, 0xa4, 0x45, 0xa3,
		0xff, 0x41, 0x5a, 0x5d, 0x42, 0x4a, 0xe7, 0x21, 0x1b, 0xa4, 0xa9, 0x68, 0xfc, 0x3f, 0x0a, 0x7c,
		0x07, 0x43, 0x2d, 0x10, 0x4a, 0x93, 0xd1, 0x14, 0xff, 0x24, 0x2d, 0x10, 0x42, 0xd1, 0x63, 0xd4,
		0x5b, 0xfa, 0xa2, 0x99, 0xfe, 0x59, 0x1e, 0xa3, 0x9e, 0xca, 0x47, 0xbd, 0xc9, 0xb2, 0x45, 0x34,
		0xc5, 0xbf, 0x48, 0x6f, 0x32, 0x7d, 0xba, 0x8c, 0xde, 0x5a, 0x12, 0xcd, 0xf1, 0xaf, 0x72, 0x19,
		0x3d, 0xa5, 0xa4, 0x54, 0x05, 0xd4, 0x5f, 0x47, 0xa2, 0xf9, 0x5e, 0x11, 0x7c, 0xe3, 0x7d, 0x65,
		0xa4, 0xf4, 0x0c, 0x4c, 0x0d, 0xae, 0x21, 0xd1, 0xac, 0xaf, 0xde, 0xea, 0xe9, 0xfa, 0xc3, 0x25,
		0xa4, 0xb4, 0xd9, 0xe9, 0xfa, 0xc3, 0xf5, 0x23, 0x9a, 0xf6, 0xb5, 0x5b, 0xdd, 0x2f, 0x76, 0xe1,
		0xf2, 0x51, 0x2a, 0x03, 0x74, 0x52, 0x77, 0x34, 0xd7, 0x1b, 0x82, 0x2b, 0x04, 0xa2, 0x47, 0x43,
		0x64, 0xee, 0x68, 0xfc, 0x9b, 0xf2, 0x68, 0x08, 0x44, 0xe9, 0x1c, 0x64, 0xec, 0xb6, 0x65, 0xd1,
		0xe0, 0x40, 0x77, 0xfe, 0x41, 0x48, 0xf1, 0x67, 0xb7, 0xc5, 0xc1, 0x90, 0x80, 0xd2, 0x29, 0x48,
		0x91, 0xe6, 0x36, 0x31, 0xa2, 0x90, 0x3f, 0xbf, 0x2d, 0x13, 0x02, 0xd5, 0x2e, 0x9d, 0x07, 0xe0,
		0x2f, 0x8d, 0xec, 0x7b, 0x40, 0x04, 0xf6, 0x17, 0xb7, 0xc5, 0xb7, 0xe6, 0x0e, 0xa4, 0x43, 0xc0,
		0xbf, 0x5c, 0xdf, 0x99, 0xe0, 0xd3, 0x6e, 0x02, 0xf6, 0xa2, 0x79, 0x16, 0x46, 0x2e, 0x7b, 0x8e,
		0xed, 0xe3, 0x46, 0x14, 0xfa, 0x97, 0x02, 0x2d, 0xf5, 0xa9, 0xc1, 0x9a, 0x8e, 0x4b, 0x7c, 0xdc,
		0xf0, 0xa2, 0xb0, 0xbf, 0x12, 0xd8, 0x00, 0x40, 0xc1, 0x3a, 0xf6, 0xfc, 0x61, 0xf6, 0xfd, 0x6b,
		0x09, 0x96, 0x00, 0xba, 0x68, 0xfa, 0xff, 0x15, 0xb2, 0x1b, 0x85, 0xfd, 0x4c, 0x2e, 0x5a, 0xe8,
		0x97, 0x9e, 0x80, 0x2c, 0xfd, 0x97, 0xff, 0xfe, 0x22, 0x02, 0xfc, 0x1b, 0x01, 0xee, 0x20, 0xe8,
		0x93, 0x3d, 0xdf, 0xf0, 0xcd, 0x68, 0x63, 0xff, 0x56, 0x78, 0x5a, 0xea, 0x97, 0xca, 0x90, 0xf3,
		0x7c, 0xc3, 0x68, 0xbb, 0xfc, 0x22, 0x2a, 0x02, 0xfe, 0xbb, 0xdb, 0xc1, 0xcb, 0x5c, 0x80, 0x59,
		0x38, 0x32, 0xf8, 0x6e, 0x09, 0x96, 0x9d, 0x65, 0x87, 0xdf, 0x2a, 0xc1, 0xe7, 0x31, 0xc8, 0xb7,
		0x88, 0xeb, 0x39, 0xb6, 0xb8, 0x00, 0x4a, 0x36, 0xb1, 0x69, 0x4f, 0xef, 0xef, 0xd6, 0x68, 0xb6,
		0x09, 0xe9, 0x2a, 0x23, 0x41, 0x08, 0x92, 0xeb, 0xa1, 0x1f, 0x48, 0xb0, 0xdf, 0x4f, 0x1d, 0x83,
		0x6c, 0xd9, 0x30, 0x5c, 0xe2, 0x79, 0xc4, 0x13, 0x57, 0xcd, 0xa3, 0xf3, 0xf4, 0x31, 0xf3, 0x42,
		0xac, 0x76, 0xe6, 0xd1, 0xbd, 0x90, 0xdd, 0x24, 0x16, 0x69, 0xed, 0x38, 0xb6, 0xbc, 0x48, 0xee,
		0x08, 0x4a, 0xc9, 0xcf, 0xde, 0x39, 0x1c, 0x9b, 0x3d, 0x0b, 0x23, 0x02, 0x80, 0xa6, 0x20, 0xbd,
		0xce, 0x7f, 0xa9, 0x12, 0x63, 0x37, 0xc3, 0x62, 0x44, 0xe5, 0x35, 0xdf, 0x25, 0xc4, 0x17, 0x57,
		0x73, 0x62, 0xb4, 0x90, 0xf9, 0xec, 0xe3, 0x43, 0xb1, 0xdf, 0x7f, 0x7c, 0x28, 0xf6, 0x87, 0x00,
		0x00, 0x00, 0xff, 0xff, 0x8b, 0x1a, 0x76, 0x67, 0x20, 0x2d, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *Person) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&main.Person{")
	if this.Name != nil {
		s = append(s, "Name: "+valueToGoStringPerson(this.Name, "string")+",\n")
	}
	if this.Addresses != nil {
		s = append(s, "Addresses: "+fmt.Sprintf("%#v", this.Addresses)+",\n")
	}
	if this.Telephone != nil {
		s = append(s, "Telephone: "+valueToGoStringPerson(this.Telephone, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Address) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&main.Address{")
	if this.Number != nil {
		s = append(s, "Number: "+valueToGoStringPerson(this.Number, "int64")+",\n")
	}
	if this.Street != nil {
		s = append(s, "Street: "+valueToGoStringPerson(this.Street, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPerson(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedPerson(r randyPerson, easy bool) *Person {
	this := &Person{}
	if r.Intn(10) != 0 {
		v1 := string(randStringPerson(r))
		this.Name = &v1
	}
	if r.Intn(10) != 0 {
		v2 := r.Intn(5)
		this.Addresses = make([]*Address, v2)
		for i := 0; i < v2; i++ {
			this.Addresses[i] = NewPopulatedAddress(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		v3 := string(randStringPerson(r))
		this.Telephone = &v3
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedPerson(r, 4)
	}
	return this
}

func NewPopulatedAddress(r randyPerson, easy bool) *Address {
	this := &Address{}
	if r.Intn(10) != 0 {
		v4 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		this.Number = &v4
	}
	if r.Intn(10) != 0 {
		v5 := string(randStringPerson(r))
		this.Street = &v5
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedPerson(r, 3)
	}
	return this
}

type randyPerson interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RunePerson(r randyPerson) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringPerson(r randyPerson) string {
	v6 := r.Intn(100)
	tmps := make([]rune, v6)
	for i := 0; i < v6; i++ {
		tmps[i] = randUTF8RunePerson(r)
	}
	return string(tmps)
}
func randUnrecognizedPerson(r randyPerson, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldPerson(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldPerson(dAtA []byte, r randyPerson, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulatePerson(dAtA, uint64(key))
		v7 := r.Int63()
		if r.Intn(2) == 0 {
			v7 *= -1
		}
		dAtA = encodeVarintPopulatePerson(dAtA, uint64(v7))
	case 1:
		dAtA = encodeVarintPopulatePerson(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulatePerson(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulatePerson(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulatePerson(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulatePerson(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("person.proto", fileDescriptorPerson) }

var fileDescriptorPerson = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x93, 0xd2,
	0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7,
	0x07, 0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa4, 0x94, 0xcb, 0xc5,
	0x16, 0x00, 0x36, 0x44, 0x48, 0x88, 0x8b, 0xc5, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x08, 0xcc, 0x16, 0xd2, 0xe6, 0xe2, 0x74, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x4e,
	0x2d, 0x96, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0xe2, 0xd5, 0x03, 0x59, 0xa3, 0x07, 0x15, 0x0e,
	0x42, 0xc8, 0x0b, 0xc9, 0x70, 0x71, 0x86, 0xa4, 0xe6, 0xa4, 0x16, 0x64, 0xe4, 0xe7, 0xa5, 0x4a,
	0x30, 0x83, 0x4d, 0x41, 0x08, 0x58, 0xb1, 0x7c, 0x58, 0x20, 0xcf, 0xa8, 0x64, 0xc9, 0xc5, 0x0e,
	0xd5, 0x20, 0x24, 0xc6, 0xc5, 0xe6, 0x57, 0x9a, 0x9b, 0x94, 0x5a, 0x04, 0xb6, 0x91, 0x39, 0x08,
	0xca, 0x03, 0x89, 0x07, 0x97, 0x14, 0xa5, 0xa6, 0x96, 0x48, 0x30, 0x81, 0xcd, 0x80, 0xf2, 0x9c,
	0x38, 0x3e, 0x3c, 0x94, 0x63, 0xfc, 0xf1, 0x50, 0x8e, 0x11, 0x10, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0xe6, 0x3e, 0x1b, 0xf7, 0x00, 0x00, 0x00,
}
