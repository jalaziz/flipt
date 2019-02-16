// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/keyword_plan_competition_level.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Competition level of a keyword.
type KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel int32

const (
	// Not specified.
	KeywordPlanCompetitionLevelEnum_UNSPECIFIED KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel = 0
	// The value is unknown in this version.
	KeywordPlanCompetitionLevelEnum_UNKNOWN KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel = 1
	// Low competition.
	KeywordPlanCompetitionLevelEnum_LOW KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel = 2
	// Medium competition.
	KeywordPlanCompetitionLevelEnum_MEDIUM KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel = 3
	// High competition.
	KeywordPlanCompetitionLevelEnum_HIGH KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel = 4
)

var KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "LOW",
	3: "MEDIUM",
	4: "HIGH",
}
var KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"LOW":         2,
	"MEDIUM":      3,
	"HIGH":        4,
}

func (x KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel) String() string {
	return proto.EnumName(KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel_name, int32(x))
}
func (KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_competition_level_253209b650169a13, []int{0, 0}
}

// Container for enumeration of keyword competition levels. The competition
// level indicates how competitive ad placement is for a keyword and
// is determined by the number of advertisers bidding on that keyword relative
// to all keywords across Google. The competition level can depend on the
// location and Search Network targeting options you've selected.
type KeywordPlanCompetitionLevelEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanCompetitionLevelEnum) Reset()         { *m = KeywordPlanCompetitionLevelEnum{} }
func (m *KeywordPlanCompetitionLevelEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanCompetitionLevelEnum) ProtoMessage()    {}
func (*KeywordPlanCompetitionLevelEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_competition_level_253209b650169a13, []int{0}
}
func (m *KeywordPlanCompetitionLevelEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanCompetitionLevelEnum.Unmarshal(m, b)
}
func (m *KeywordPlanCompetitionLevelEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanCompetitionLevelEnum.Marshal(b, m, deterministic)
}
func (dst *KeywordPlanCompetitionLevelEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanCompetitionLevelEnum.Merge(dst, src)
}
func (m *KeywordPlanCompetitionLevelEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanCompetitionLevelEnum.Size(m)
}
func (m *KeywordPlanCompetitionLevelEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanCompetitionLevelEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanCompetitionLevelEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*KeywordPlanCompetitionLevelEnum)(nil), "google.ads.googleads.v0.enums.KeywordPlanCompetitionLevelEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel", KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel_name, KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/keyword_plan_competition_level.proto", fileDescriptor_keyword_plan_competition_level_253209b650169a13)
}

var fileDescriptor_keyword_plan_competition_level_253209b650169a13 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x6a, 0x83, 0x30,
	0x18, 0xc7, 0xa7, 0x96, 0x76, 0xa4, 0x87, 0x49, 0xae, 0xa3, 0x6c, 0xeb, 0x03, 0x44, 0x61, 0xc7,
	0x9d, 0xb4, 0x75, 0x56, 0xda, 0x5a, 0x61, 0xd8, 0x42, 0x11, 0xc4, 0xd5, 0x10, 0x64, 0x31, 0x9f,
	0x18, 0xeb, 0xd8, 0x61, 0x2f, 0xb3, 0xe3, 0x1e, 0x65, 0xb7, 0xbd, 0xd1, 0x50, 0xd7, 0xee, 0x34,
	0x2f, 0xe1, 0x0f, 0xff, 0xe4, 0xf7, 0xe5, 0xfb, 0x21, 0x9b, 0x01, 0x30, 0x4e, 0x8d, 0x24, 0x95,
	0x46, 0x17, 0x9b, 0x54, 0x9b, 0x06, 0x15, 0xc7, 0x5c, 0x1a, 0x2f, 0xf4, 0xed, 0x15, 0xca, 0x34,
	0x2e, 0x78, 0x22, 0xe2, 0x03, 0xe4, 0x05, 0xad, 0xb2, 0x2a, 0x03, 0x11, 0x73, 0x5a, 0x53, 0x4e,
	0x8a, 0x12, 0x2a, 0xc0, 0x93, 0xee, 0x21, 0x49, 0x52, 0x49, 0xce, 0x0c, 0x52, 0x9b, 0xa4, 0x65,
	0x4c, 0xdf, 0xd1, 0xcd, 0xb2, 0xc3, 0x04, 0x3c, 0x11, 0xb3, 0x3f, 0xc8, 0xaa, 0x61, 0x38, 0xe2,
	0x98, 0x4f, 0xf7, 0xe8, 0xba, 0xe7, 0x0a, 0xbe, 0x42, 0xe3, 0xd0, 0x7f, 0x0a, 0x9c, 0x99, 0xf7,
	0xe8, 0x39, 0x73, 0xfd, 0x02, 0x8f, 0xd1, 0x28, 0xf4, 0x97, 0xfe, 0x66, 0xe7, 0xeb, 0x0a, 0x1e,
	0x21, 0x6d, 0xb5, 0xd9, 0xe9, 0x2a, 0x46, 0x68, 0xb8, 0x76, 0xe6, 0x5e, 0xb8, 0xd6, 0x35, 0x7c,
	0x89, 0x06, 0x0b, 0xcf, 0x5d, 0xe8, 0x03, 0xfb, 0x5b, 0x41, 0x77, 0x07, 0xc8, 0x49, 0xef, 0x27,
	0xed, 0xdb, 0x9e, 0xf9, 0x41, 0xb3, 0x65, 0xa0, 0xec, 0x7f, 0x5d, 0x11, 0x06, 0x3c, 0x11, 0x8c,
	0x40, 0xc9, 0x0c, 0x46, 0x45, 0xeb, 0xe0, 0xe4, 0xae, 0xc8, 0xe4, 0x3f, 0x2a, 0x1f, 0xda, 0xf3,
	0x43, 0xd5, 0x5c, 0xcb, 0xfa, 0x54, 0x27, 0x6e, 0x87, 0xb2, 0x52, 0x49, 0xba, 0xd8, 0xa4, 0xad,
	0x49, 0x1a, 0x1d, 0xf2, 0xeb, 0xd4, 0x47, 0x56, 0x2a, 0xa3, 0x73, 0x1f, 0x6d, 0xcd, 0xa8, 0xed,
	0x9f, 0x87, 0xed, 0xd0, 0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x98, 0xeb, 0x0f, 0xbe,
	0x01, 0x00, 0x00,
}
