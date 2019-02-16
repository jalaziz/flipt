// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/call_placeholder_field.proto

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

// Possible values for Call placeholder fields.
type CallPlaceholderFieldEnum_CallPlaceholderField int32

const (
	// Not specified.
	CallPlaceholderFieldEnum_UNSPECIFIED CallPlaceholderFieldEnum_CallPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	CallPlaceholderFieldEnum_UNKNOWN CallPlaceholderFieldEnum_CallPlaceholderField = 1
	// Data Type: STRING. The advertiser's phone number to append to the ad.
	CallPlaceholderFieldEnum_PHONE_NUMBER CallPlaceholderFieldEnum_CallPlaceholderField = 2
	// Data Type: STRING. Uppercase two-letter country code of the advertiser's
	// phone number.
	CallPlaceholderFieldEnum_COUNTRY_CODE CallPlaceholderFieldEnum_CallPlaceholderField = 3
	// Data Type: BOOLEAN. Indicates whether call tracking is enabled. Default:
	// true.
	CallPlaceholderFieldEnum_TRACKED CallPlaceholderFieldEnum_CallPlaceholderField = 4
	// Data Type: INT64. The ID of an AdCallMetricsConversion object. This
	// object contains the phoneCallDurationfield which is the minimum duration
	// (in seconds) of a call to be considered a conversion.
	CallPlaceholderFieldEnum_CONVERSION_TYPE_ID CallPlaceholderFieldEnum_CallPlaceholderField = 5
	// Data Type: STRING. Indicates whether this call extension uses its own
	// call conversion setting or follows the account level setting.
	// Valid values are: USE_ACCOUNT_LEVEL_CALL_CONVERSION_ACTION and
	// USE_RESOURCE_LEVEL_CALL_CONVERSION_ACTION.
	CallPlaceholderFieldEnum_CONVERSION_REPORTING_STATE CallPlaceholderFieldEnum_CallPlaceholderField = 6
)

var CallPlaceholderFieldEnum_CallPlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PHONE_NUMBER",
	3: "COUNTRY_CODE",
	4: "TRACKED",
	5: "CONVERSION_TYPE_ID",
	6: "CONVERSION_REPORTING_STATE",
}
var CallPlaceholderFieldEnum_CallPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"PHONE_NUMBER":               2,
	"COUNTRY_CODE":               3,
	"TRACKED":                    4,
	"CONVERSION_TYPE_ID":         5,
	"CONVERSION_REPORTING_STATE": 6,
}

func (x CallPlaceholderFieldEnum_CallPlaceholderField) String() string {
	return proto.EnumName(CallPlaceholderFieldEnum_CallPlaceholderField_name, int32(x))
}
func (CallPlaceholderFieldEnum_CallPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_call_placeholder_field_1efb02a4392b5d2d, []int{0, 0}
}

// Values for Call placeholder fields.
type CallPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallPlaceholderFieldEnum) Reset()         { *m = CallPlaceholderFieldEnum{} }
func (m *CallPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*CallPlaceholderFieldEnum) ProtoMessage()    {}
func (*CallPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_call_placeholder_field_1efb02a4392b5d2d, []int{0}
}
func (m *CallPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *CallPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (dst *CallPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallPlaceholderFieldEnum.Merge(dst, src)
}
func (m *CallPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_CallPlaceholderFieldEnum.Size(m)
}
func (m *CallPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CallPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CallPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CallPlaceholderFieldEnum)(nil), "google.ads.googleads.v0.enums.CallPlaceholderFieldEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.CallPlaceholderFieldEnum_CallPlaceholderField", CallPlaceholderFieldEnum_CallPlaceholderField_name, CallPlaceholderFieldEnum_CallPlaceholderField_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/call_placeholder_field.proto", fileDescriptor_call_placeholder_field_1efb02a4392b5d2d)
}

var fileDescriptor_call_placeholder_field_1efb02a4392b5d2d = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0xed, 0xa6, 0x13, 0x32, 0xc1, 0x50, 0x44, 0x54, 0x98, 0xe0, 0x1e, 0x20, 0x2d, 0x78,
	0xa7, 0x57, 0xfd, 0x93, 0xcd, 0x32, 0x4c, 0x42, 0xd7, 0x4e, 0x26, 0x85, 0x50, 0xd7, 0x5a, 0x07,
	0x59, 0x33, 0x16, 0xb7, 0xb7, 0xf1, 0xc6, 0x4b, 0x1f, 0xc1, 0x47, 0xd8, 0x53, 0x49, 0x5a, 0x37,
	0xbd, 0x98, 0xde, 0x84, 0x8f, 0xf3, 0xf1, 0xcb, 0xe1, 0xfc, 0xc0, 0x4d, 0x21, 0x65, 0x21, 0x72,
	0x2b, 0xcd, 0x94, 0x55, 0x47, 0x9d, 0x56, 0xb6, 0x95, 0x97, 0xcb, 0x99, 0xb2, 0x26, 0xa9, 0x10,
	0x7c, 0x2e, 0xd2, 0x49, 0xfe, 0x22, 0x45, 0x96, 0x2f, 0xf8, 0xf3, 0x34, 0x17, 0x19, 0x9a, 0x2f,
	0xe4, 0xab, 0x34, 0x3b, 0x35, 0x80, 0xd2, 0x4c, 0xa1, 0x2d, 0x8b, 0x56, 0x36, 0xaa, 0xd8, 0xee,
	0xa7, 0x01, 0xce, 0xbc, 0x54, 0x08, 0xf6, 0x83, 0xf7, 0x34, 0x8d, 0xcb, 0xe5, 0xac, 0xfb, 0x66,
	0x80, 0x93, 0x5d, 0xa5, 0x79, 0x0c, 0xda, 0x31, 0x19, 0x32, 0xec, 0x05, 0xbd, 0x00, 0xfb, 0x70,
	0xcf, 0x6c, 0x83, 0xc3, 0x98, 0x0c, 0x08, 0x7d, 0x20, 0xd0, 0x30, 0x21, 0x38, 0x62, 0x77, 0x94,
	0x60, 0x4e, 0xe2, 0x7b, 0x17, 0x87, 0xb0, 0xa1, 0x27, 0x1e, 0x8d, 0x49, 0x14, 0x8e, 0xb9, 0x47,
	0x7d, 0x0c, 0x9b, 0x1a, 0x88, 0x42, 0xc7, 0x1b, 0x60, 0x1f, 0xee, 0x9b, 0xa7, 0xc0, 0xf4, 0x28,
	0x19, 0xe1, 0x70, 0x18, 0x50, 0xc2, 0xa3, 0x31, 0xc3, 0x3c, 0xf0, 0xe1, 0x81, 0x79, 0x09, 0x2e,
	0x7e, 0xcd, 0x43, 0xcc, 0x68, 0x18, 0x05, 0xa4, 0xcf, 0x87, 0x91, 0x13, 0x61, 0xd8, 0x72, 0xd7,
	0x06, 0xb8, 0x9a, 0xc8, 0x19, 0xfa, 0xf7, 0x44, 0xf7, 0x7c, 0xd7, 0x09, 0x4c, 0xcb, 0x61, 0xc6,
	0xa3, 0xfb, 0xcd, 0x16, 0x52, 0xa4, 0x65, 0x81, 0xe4, 0xa2, 0xb0, 0x8a, 0xbc, 0xac, 0xd4, 0x6d,
	0x54, 0xcf, 0xa7, 0xea, 0x0f, 0xf3, 0xb7, 0xd5, 0xfb, 0xde, 0x68, 0xf6, 0x1d, 0xe7, 0xa3, 0xd1,
	0xe9, 0xd7, 0x5f, 0x39, 0x99, 0x42, 0x75, 0xd4, 0x69, 0x64, 0x23, 0xed, 0x52, 0xad, 0x37, 0x7d,
	0xe2, 0x64, 0x2a, 0xd9, 0xf6, 0xc9, 0xc8, 0x4e, 0xaa, 0xfe, 0xa9, 0x55, 0x2d, 0xbd, 0xfe, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0xfb, 0xf5, 0x10, 0x42, 0xed, 0x01, 0x00, 0x00,
}
