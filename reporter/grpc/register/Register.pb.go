// Code generated by protoc-gen-go. DO NOT EDIT.
// source: register/Register.proto

package register

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/tetratelabs/go2sky/reporter/grpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Service register
type Services struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Services) Reset()         { *m = Services{} }
func (m *Services) String() string { return proto.CompactTextString(m) }
func (*Services) ProtoMessage()    {}
func (*Services) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{0}
}

func (m *Services) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Services.Unmarshal(m, b)
}
func (m *Services) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Services.Marshal(b, m, deterministic)
}
func (m *Services) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Services.Merge(m, src)
}
func (m *Services) XXX_Size() int {
	return xxx_messageInfo_Services.Size(m)
}
func (m *Services) XXX_DiscardUnknown() {
	xxx_messageInfo_Services.DiscardUnknown(m)
}

var xxx_messageInfo_Services proto.InternalMessageInfo

func (m *Services) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type Service struct {
	ServiceName          string                       `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Tags                 []*common.KeyStringValuePair `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Properties           []*common.KeyStringValuePair `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{1}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Service) GetTags() []*common.KeyStringValuePair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Service) GetProperties() []*common.KeyStringValuePair {
	if m != nil {
		return m.Properties
	}
	return nil
}

type ServiceRegisterMapping struct {
	Services             []*common.KeyIntValuePair `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ServiceRegisterMapping) Reset()         { *m = ServiceRegisterMapping{} }
func (m *ServiceRegisterMapping) String() string { return proto.CompactTextString(m) }
func (*ServiceRegisterMapping) ProtoMessage()    {}
func (*ServiceRegisterMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{2}
}

func (m *ServiceRegisterMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceRegisterMapping.Unmarshal(m, b)
}
func (m *ServiceRegisterMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceRegisterMapping.Marshal(b, m, deterministic)
}
func (m *ServiceRegisterMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceRegisterMapping.Merge(m, src)
}
func (m *ServiceRegisterMapping) XXX_Size() int {
	return xxx_messageInfo_ServiceRegisterMapping.Size(m)
}
func (m *ServiceRegisterMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceRegisterMapping.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceRegisterMapping proto.InternalMessageInfo

func (m *ServiceRegisterMapping) GetServices() []*common.KeyIntValuePair {
	if m != nil {
		return m.Services
	}
	return nil
}

// Service Instance register
type ServiceInstances struct {
	Instances            []*ServiceInstance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ServiceInstances) Reset()         { *m = ServiceInstances{} }
func (m *ServiceInstances) String() string { return proto.CompactTextString(m) }
func (*ServiceInstances) ProtoMessage()    {}
func (*ServiceInstances) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{3}
}

func (m *ServiceInstances) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInstances.Unmarshal(m, b)
}
func (m *ServiceInstances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInstances.Marshal(b, m, deterministic)
}
func (m *ServiceInstances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInstances.Merge(m, src)
}
func (m *ServiceInstances) XXX_Size() int {
	return xxx_messageInfo_ServiceInstances.Size(m)
}
func (m *ServiceInstances) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInstances.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInstances proto.InternalMessageInfo

func (m *ServiceInstances) GetInstances() []*ServiceInstance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type ServiceInstance struct {
	ServiceId            int32                        `protobuf:"varint,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	InstanceUUID         string                       `protobuf:"bytes,2,opt,name=instanceUUID,proto3" json:"instanceUUID,omitempty"`
	Time                 int64                        `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Tags                 []*common.KeyStringValuePair `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	Properties           []*common.KeyStringValuePair `protobuf:"bytes,5,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ServiceInstance) Reset()         { *m = ServiceInstance{} }
func (m *ServiceInstance) String() string { return proto.CompactTextString(m) }
func (*ServiceInstance) ProtoMessage()    {}
func (*ServiceInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{4}
}

func (m *ServiceInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInstance.Unmarshal(m, b)
}
func (m *ServiceInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInstance.Marshal(b, m, deterministic)
}
func (m *ServiceInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInstance.Merge(m, src)
}
func (m *ServiceInstance) XXX_Size() int {
	return xxx_messageInfo_ServiceInstance.Size(m)
}
func (m *ServiceInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInstance.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInstance proto.InternalMessageInfo

func (m *ServiceInstance) GetServiceId() int32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *ServiceInstance) GetInstanceUUID() string {
	if m != nil {
		return m.InstanceUUID
	}
	return ""
}

func (m *ServiceInstance) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *ServiceInstance) GetTags() []*common.KeyStringValuePair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ServiceInstance) GetProperties() []*common.KeyStringValuePair {
	if m != nil {
		return m.Properties
	}
	return nil
}

type ServiceInstanceRegisterMapping struct {
	ServiceInstances     []*common.KeyIntValuePair `protobuf:"bytes,1,rep,name=serviceInstances,proto3" json:"serviceInstances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ServiceInstanceRegisterMapping) Reset()         { *m = ServiceInstanceRegisterMapping{} }
func (m *ServiceInstanceRegisterMapping) String() string { return proto.CompactTextString(m) }
func (*ServiceInstanceRegisterMapping) ProtoMessage()    {}
func (*ServiceInstanceRegisterMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{5}
}

func (m *ServiceInstanceRegisterMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInstanceRegisterMapping.Unmarshal(m, b)
}
func (m *ServiceInstanceRegisterMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInstanceRegisterMapping.Marshal(b, m, deterministic)
}
func (m *ServiceInstanceRegisterMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInstanceRegisterMapping.Merge(m, src)
}
func (m *ServiceInstanceRegisterMapping) XXX_Size() int {
	return xxx_messageInfo_ServiceInstanceRegisterMapping.Size(m)
}
func (m *ServiceInstanceRegisterMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInstanceRegisterMapping.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInstanceRegisterMapping proto.InternalMessageInfo

func (m *ServiceInstanceRegisterMapping) GetServiceInstances() []*common.KeyIntValuePair {
	if m != nil {
		return m.ServiceInstances
	}
	return nil
}

// Only known use case is the language agent.
// Network address represents the ip/hostname:port, which is usually used at client side of RPC.
type NetAddresses struct {
	Addresses            []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetAddresses) Reset()         { *m = NetAddresses{} }
func (m *NetAddresses) String() string { return proto.CompactTextString(m) }
func (*NetAddresses) ProtoMessage()    {}
func (*NetAddresses) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{6}
}

func (m *NetAddresses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetAddresses.Unmarshal(m, b)
}
func (m *NetAddresses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetAddresses.Marshal(b, m, deterministic)
}
func (m *NetAddresses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetAddresses.Merge(m, src)
}
func (m *NetAddresses) XXX_Size() int {
	return xxx_messageInfo_NetAddresses.Size(m)
}
func (m *NetAddresses) XXX_DiscardUnknown() {
	xxx_messageInfo_NetAddresses.DiscardUnknown(m)
}

var xxx_messageInfo_NetAddresses proto.InternalMessageInfo

func (m *NetAddresses) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type NetAddressMapping struct {
	AddressIds           []*common.KeyIntValuePair `protobuf:"bytes,1,rep,name=addressIds,proto3" json:"addressIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *NetAddressMapping) Reset()         { *m = NetAddressMapping{} }
func (m *NetAddressMapping) String() string { return proto.CompactTextString(m) }
func (*NetAddressMapping) ProtoMessage()    {}
func (*NetAddressMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{7}
}

func (m *NetAddressMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetAddressMapping.Unmarshal(m, b)
}
func (m *NetAddressMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetAddressMapping.Marshal(b, m, deterministic)
}
func (m *NetAddressMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetAddressMapping.Merge(m, src)
}
func (m *NetAddressMapping) XXX_Size() int {
	return xxx_messageInfo_NetAddressMapping.Size(m)
}
func (m *NetAddressMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_NetAddressMapping.DiscardUnknown(m)
}

var xxx_messageInfo_NetAddressMapping proto.InternalMessageInfo

func (m *NetAddressMapping) GetAddressIds() []*common.KeyIntValuePair {
	if m != nil {
		return m.AddressIds
	}
	return nil
}

// Endpint register
type Enpoints struct {
	Endpoints            []*Endpoint `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Enpoints) Reset()         { *m = Enpoints{} }
func (m *Enpoints) String() string { return proto.CompactTextString(m) }
func (*Enpoints) ProtoMessage()    {}
func (*Enpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{8}
}

func (m *Enpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Enpoints.Unmarshal(m, b)
}
func (m *Enpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Enpoints.Marshal(b, m, deterministic)
}
func (m *Enpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Enpoints.Merge(m, src)
}
func (m *Enpoints) XXX_Size() int {
	return xxx_messageInfo_Enpoints.Size(m)
}
func (m *Enpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_Enpoints.DiscardUnknown(m)
}

var xxx_messageInfo_Enpoints proto.InternalMessageInfo

func (m *Enpoints) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type Endpoint struct {
	ServiceId    int32                        `protobuf:"varint,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	EndpointName string                       `protobuf:"bytes,2,opt,name=endpointName,proto3" json:"endpointName,omitempty"`
	Tags         []*common.KeyStringValuePair `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Properties   []*common.KeyStringValuePair `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty"`
	// For endpoint
	// from DetectPoint is either `client` or `server`. No chance to be `proxy`.
	From                 common.DetectPoint `protobuf:"varint,5,opt,name=from,proto3,enum=DetectPoint" json:"from,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{9}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetServiceId() int32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *Endpoint) GetEndpointName() string {
	if m != nil {
		return m.EndpointName
	}
	return ""
}

func (m *Endpoint) GetTags() []*common.KeyStringValuePair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Endpoint) GetProperties() []*common.KeyStringValuePair {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Endpoint) GetFrom() common.DetectPoint {
	if m != nil {
		return m.From
	}
	return common.DetectPoint_client
}

type EndpointMapping struct {
	Elements             []*EndpointMappingElement `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EndpointMapping) Reset()         { *m = EndpointMapping{} }
func (m *EndpointMapping) String() string { return proto.CompactTextString(m) }
func (*EndpointMapping) ProtoMessage()    {}
func (*EndpointMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{10}
}

func (m *EndpointMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointMapping.Unmarshal(m, b)
}
func (m *EndpointMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointMapping.Marshal(b, m, deterministic)
}
func (m *EndpointMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointMapping.Merge(m, src)
}
func (m *EndpointMapping) XXX_Size() int {
	return xxx_messageInfo_EndpointMapping.Size(m)
}
func (m *EndpointMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointMapping.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointMapping proto.InternalMessageInfo

func (m *EndpointMapping) GetElements() []*EndpointMappingElement {
	if m != nil {
		return m.Elements
	}
	return nil
}

type EndpointMappingElement struct {
	ServiceId            int32              `protobuf:"varint,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	EndpointName         string             `protobuf:"bytes,2,opt,name=endpointName,proto3" json:"endpointName,omitempty"`
	EndpointId           int32              `protobuf:"varint,3,opt,name=endpointId,proto3" json:"endpointId,omitempty"`
	From                 common.DetectPoint `protobuf:"varint,4,opt,name=from,proto3,enum=DetectPoint" json:"from,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EndpointMappingElement) Reset()         { *m = EndpointMappingElement{} }
func (m *EndpointMappingElement) String() string { return proto.CompactTextString(m) }
func (*EndpointMappingElement) ProtoMessage()    {}
func (*EndpointMappingElement) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{11}
}

func (m *EndpointMappingElement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointMappingElement.Unmarshal(m, b)
}
func (m *EndpointMappingElement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointMappingElement.Marshal(b, m, deterministic)
}
func (m *EndpointMappingElement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointMappingElement.Merge(m, src)
}
func (m *EndpointMappingElement) XXX_Size() int {
	return xxx_messageInfo_EndpointMappingElement.Size(m)
}
func (m *EndpointMappingElement) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointMappingElement.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointMappingElement proto.InternalMessageInfo

func (m *EndpointMappingElement) GetServiceId() int32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *EndpointMappingElement) GetEndpointName() string {
	if m != nil {
		return m.EndpointName
	}
	return ""
}

func (m *EndpointMappingElement) GetEndpointId() int32 {
	if m != nil {
		return m.EndpointId
	}
	return 0
}

func (m *EndpointMappingElement) GetFrom() common.DetectPoint {
	if m != nil {
		return m.From
	}
	return common.DetectPoint_client
}

type ServiceAndNetworkAddressMappings struct {
	Mappings             []*ServiceAndNetworkAddressMapping `protobuf:"bytes,1,rep,name=mappings,proto3" json:"mappings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ServiceAndNetworkAddressMappings) Reset()         { *m = ServiceAndNetworkAddressMappings{} }
func (m *ServiceAndNetworkAddressMappings) String() string { return proto.CompactTextString(m) }
func (*ServiceAndNetworkAddressMappings) ProtoMessage()    {}
func (*ServiceAndNetworkAddressMappings) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{12}
}

func (m *ServiceAndNetworkAddressMappings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAndNetworkAddressMappings.Unmarshal(m, b)
}
func (m *ServiceAndNetworkAddressMappings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAndNetworkAddressMappings.Marshal(b, m, deterministic)
}
func (m *ServiceAndNetworkAddressMappings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAndNetworkAddressMappings.Merge(m, src)
}
func (m *ServiceAndNetworkAddressMappings) XXX_Size() int {
	return xxx_messageInfo_ServiceAndNetworkAddressMappings.Size(m)
}
func (m *ServiceAndNetworkAddressMappings) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAndNetworkAddressMappings.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAndNetworkAddressMappings proto.InternalMessageInfo

func (m *ServiceAndNetworkAddressMappings) GetMappings() []*ServiceAndNetworkAddressMapping {
	if m != nil {
		return m.Mappings
	}
	return nil
}

type ServiceAndNetworkAddressMapping struct {
	ServiceId            int32    `protobuf:"varint,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	ServiceInstanceId    int32    `protobuf:"varint,2,opt,name=serviceInstanceId,proto3" json:"serviceInstanceId,omitempty"`
	NetworkAddress       string   `protobuf:"bytes,3,opt,name=networkAddress,proto3" json:"networkAddress,omitempty"`
	NetworkAddressId     int32    `protobuf:"varint,4,opt,name=networkAddressId,proto3" json:"networkAddressId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceAndNetworkAddressMapping) Reset()         { *m = ServiceAndNetworkAddressMapping{} }
func (m *ServiceAndNetworkAddressMapping) String() string { return proto.CompactTextString(m) }
func (*ServiceAndNetworkAddressMapping) ProtoMessage()    {}
func (*ServiceAndNetworkAddressMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_c38a4def3f450915, []int{13}
}

func (m *ServiceAndNetworkAddressMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAndNetworkAddressMapping.Unmarshal(m, b)
}
func (m *ServiceAndNetworkAddressMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAndNetworkAddressMapping.Marshal(b, m, deterministic)
}
func (m *ServiceAndNetworkAddressMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAndNetworkAddressMapping.Merge(m, src)
}
func (m *ServiceAndNetworkAddressMapping) XXX_Size() int {
	return xxx_messageInfo_ServiceAndNetworkAddressMapping.Size(m)
}
func (m *ServiceAndNetworkAddressMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAndNetworkAddressMapping.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAndNetworkAddressMapping proto.InternalMessageInfo

func (m *ServiceAndNetworkAddressMapping) GetServiceId() int32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *ServiceAndNetworkAddressMapping) GetServiceInstanceId() int32 {
	if m != nil {
		return m.ServiceInstanceId
	}
	return 0
}

func (m *ServiceAndNetworkAddressMapping) GetNetworkAddress() string {
	if m != nil {
		return m.NetworkAddress
	}
	return ""
}

func (m *ServiceAndNetworkAddressMapping) GetNetworkAddressId() int32 {
	if m != nil {
		return m.NetworkAddressId
	}
	return 0
}

func init() {
	proto.RegisterType((*Services)(nil), "Services")
	proto.RegisterType((*Service)(nil), "Service")
	proto.RegisterType((*ServiceRegisterMapping)(nil), "ServiceRegisterMapping")
	proto.RegisterType((*ServiceInstances)(nil), "ServiceInstances")
	proto.RegisterType((*ServiceInstance)(nil), "ServiceInstance")
	proto.RegisterType((*ServiceInstanceRegisterMapping)(nil), "ServiceInstanceRegisterMapping")
	proto.RegisterType((*NetAddresses)(nil), "NetAddresses")
	proto.RegisterType((*NetAddressMapping)(nil), "NetAddressMapping")
	proto.RegisterType((*Enpoints)(nil), "Enpoints")
	proto.RegisterType((*Endpoint)(nil), "Endpoint")
	proto.RegisterType((*EndpointMapping)(nil), "EndpointMapping")
	proto.RegisterType((*EndpointMappingElement)(nil), "EndpointMappingElement")
	proto.RegisterType((*ServiceAndNetworkAddressMappings)(nil), "ServiceAndNetworkAddressMappings")
	proto.RegisterType((*ServiceAndNetworkAddressMapping)(nil), "ServiceAndNetworkAddressMapping")
}

func init() { proto.RegisterFile("register/Register.proto", fileDescriptor_c38a4def3f450915) }

var fileDescriptor_c38a4def3f450915 = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0x6e, 0xd6, 0xf6, 0xff, 0x9b, 0xb3, 0xb1, 0xb5, 0x9e, 0xb4, 0x85, 0x0a, 0x6d, 0xc1, 0x42,
	0xac, 0xa0, 0xe1, 0x4e, 0x2d, 0x37, 0x48, 0x93, 0xd0, 0xc6, 0x3a, 0x29, 0x42, 0x4c, 0x25, 0xd3,
	0x40, 0x02, 0x09, 0x91, 0x35, 0x26, 0x8b, 0xda, 0xc4, 0x91, 0xed, 0x6d, 0xea, 0x3d, 0x37, 0xbc,
	0x03, 0x4f, 0xc0, 0x5b, 0x70, 0xc1, 0x1d, 0x6f, 0xc2, 0x4b, 0xa0, 0xa6, 0x4e, 0xd2, 0xa6, 0x2b,
	0x05, 0x09, 0xae, 0x62, 0x7f, 0xe7, 0x3b, 0x27, 0x3e, 0x9f, 0xcf, 0x39, 0x86, 0x4d, 0x4e, 0x3d,
	0x5f, 0x48, 0xca, 0x9b, 0xb6, 0x5a, 0x90, 0x88, 0x33, 0xc9, 0xea, 0xeb, 0x3d, 0x16, 0x04, 0x2c,
	0x6c, 0x8e, 0x3f, 0x63, 0x10, 0xef, 0x41, 0xe5, 0x94, 0xf2, 0x2b, 0xbf, 0x47, 0x05, 0xba, 0x07,
	0x15, 0xa1, 0xd6, 0x86, 0x66, 0x16, 0x1b, 0xcb, 0xad, 0x0a, 0x51, 0x46, 0x3b, 0xb5, 0xe0, 0x4f,
	0x1a, 0xfc, 0xaf, 0x50, 0x64, 0xc2, 0xb2, 0xc2, 0x4f, 0x9c, 0x80, 0x1a, 0x9a, 0xa9, 0x35, 0x74,
	0x7b, 0x12, 0x42, 0x3b, 0x50, 0x92, 0x8e, 0x27, 0x8c, 0x62, 0x1c, 0x6f, 0x9d, 0x3c, 0xa7, 0xc3,
	0x53, 0xc9, 0xfd, 0xd0, 0x7b, 0xe5, 0x0c, 0x2e, 0x69, 0xd7, 0xf1, 0xb9, 0x1d, 0x13, 0x50, 0x1b,
	0x20, 0xe2, 0x2c, 0xa2, 0x5c, 0xfa, 0x54, 0x18, 0xa5, 0xf9, 0xf4, 0x09, 0x1a, 0x3e, 0x86, 0x8d,
	0xe4, 0x80, 0x2a, 0xd7, 0x17, 0x4e, 0x14, 0xf9, 0xa1, 0x87, 0x76, 0x67, 0x72, 0xa9, 0x8e, 0x82,
	0x59, 0xa1, 0xcc, 0x22, 0x65, 0x39, 0x1d, 0x42, 0x55, 0xc5, 0xb1, 0x42, 0x21, 0x9d, 0x70, 0xa4,
	0x06, 0x01, 0xdd, 0x4f, 0x36, 0x69, 0x88, 0x1c, 0xcb, 0xce, 0x28, 0xf8, 0x9b, 0x06, 0x6b, 0x39,
	0x33, 0xba, 0x03, 0xba, 0xfa, 0x87, 0xe5, 0xc6, 0xea, 0x94, 0xed, 0x0c, 0x40, 0x18, 0x56, 0x12,
	0xf7, 0xb3, 0x33, 0xeb, 0xc8, 0x58, 0x8a, 0xe5, 0x9b, 0xc2, 0x10, 0x82, 0x92, 0xf4, 0x03, 0x6a,
	0x14, 0x4d, 0xad, 0x51, 0xb4, 0xe3, 0x75, 0xaa, 0x69, 0xe9, 0xcf, 0x34, 0x2d, 0xff, 0x9e, 0xa6,
	0xef, 0x60, 0x2b, 0x9f, 0x65, 0x4e, 0xdb, 0x7d, 0xa8, 0x8a, 0x9c, 0x5a, 0x73, 0x35, 0x9e, 0x61,
	0xe2, 0x5d, 0x58, 0x39, 0xa1, 0xf2, 0xc0, 0x75, 0x39, 0x15, 0x82, 0x8a, 0x91, 0x46, 0x4e, 0xb2,
	0x89, 0xc3, 0xe8, 0x76, 0x06, 0xe0, 0x0e, 0xd4, 0x32, 0x76, 0x72, 0x80, 0x3d, 0x00, 0xc5, 0xb0,
	0xdc, 0xf9, 0xbf, 0x9e, 0xe0, 0xe0, 0x36, 0x54, 0x3a, 0x61, 0xc4, 0xfc, 0x50, 0x0a, 0xb4, 0x03,
	0x3a, 0x0d, 0xdd, 0xf1, 0x46, 0x39, 0xeb, 0xa4, 0xa3, 0x10, 0x3b, 0xb3, 0xe1, 0xef, 0xda, 0xc8,
	0x6b, 0xbc, 0x5b, 0x7c, 0x95, 0x89, 0x5f, 0xdc, 0x09, 0xea, 0x2a, 0x27, 0xb1, 0x7f, 0xdb, 0x0a,
	0xc8, 0x84, 0xd2, 0x07, 0xce, 0x02, 0xa3, 0x6c, 0x6a, 0x8d, 0xd5, 0xd6, 0x0a, 0x39, 0xa2, 0x92,
	0xf6, 0x64, 0x37, 0xce, 0x29, 0xb6, 0xe0, 0x63, 0x58, 0x4b, 0xb2, 0x49, 0x84, 0x6c, 0x43, 0x85,
	0x0e, 0x68, 0x40, 0x33, 0x25, 0x36, 0x49, 0x8e, 0xd3, 0x19, 0xdb, 0xed, 0x94, 0x88, 0x3f, 0x6b,
	0xb0, 0x71, 0x33, 0xe9, 0x2f, 0x88, 0xb4, 0x05, 0x90, 0xec, 0x2d, 0x37, 0xae, 0xfa, 0xb2, 0x3d,
	0x81, 0xa4, 0x69, 0x96, 0xe6, 0xa6, 0xf9, 0x1e, 0x4c, 0x55, 0xbf, 0x07, 0xa1, 0x7b, 0x42, 0xe5,
	0x35, 0xe3, 0xfd, 0xe9, 0xfa, 0x11, 0x68, 0x1f, 0x2a, 0x81, 0x5a, 0xab, 0xbc, 0x4d, 0xb2, 0xc0,
	0xc9, 0x4e, 0x3d, 0xf0, 0x57, 0x0d, 0xb6, 0x17, 0xb0, 0x17, 0x28, 0xb1, 0x0b, 0xb5, 0x5c, 0x5f,
	0x58, 0x6e, 0x2c, 0x47, 0xd9, 0x9e, 0x35, 0xa0, 0xfb, 0xb0, 0x1a, 0x4e, 0xfd, 0x24, 0xd6, 0x45,
	0xb7, 0x73, 0x28, 0x7a, 0x08, 0xd5, 0x69, 0xc4, 0x72, 0x63, 0x9d, 0xca, 0xf6, 0x0c, 0xde, 0xfa,
	0xb1, 0x04, 0x95, 0xa4, 0xaf, 0xd1, 0x13, 0xa8, 0xb9, 0x2c, 0x37, 0x48, 0x91, 0x9e, 0x28, 0x22,
	0xea, 0x9b, 0xe4, 0xe6, 0x29, 0x8b, 0x0b, 0xe8, 0x25, 0xdc, 0x4e, 0x5d, 0xf3, 0xf3, 0x02, 0xd5,
	0xf2, 0xf3, 0x52, 0xd4, 0xb7, 0xc9, 0xaf, 0x87, 0x0b, 0x2e, 0xa0, 0x36, 0x20, 0x97, 0xa5, 0xfd,
	0x98, 0x1d, 0x27, 0x69, 0xe0, 0x7a, 0x35, 0x5f, 0xa3, 0xb8, 0x80, 0x9e, 0x82, 0xe1, 0xb2, 0xe9,
	0xab, 0x48, 0x5d, 0x6f, 0x91, 0xc9, 0x81, 0x53, 0x47, 0x64, 0x66, 0xa2, 0xe0, 0x02, 0x7a, 0x0b,
	0x0f, 0xd2, 0x44, 0xe6, 0xd6, 0x40, 0x12, 0xf1, 0xee, 0xa2, 0x6a, 0x11, 0x75, 0x9d, 0x3c, 0x63,
	0x41, 0xe0, 0x84, 0xae, 0xc0, 0x85, 0xc3, 0x8f, 0x1a, 0x3c, 0x62, 0xdc, 0x23, 0x4e, 0xe4, 0xf4,
	0x2e, 0x28, 0x11, 0xfd, 0xe1, 0xb5, 0x33, 0xe8, 0xfb, 0xe1, 0x08, 0x09, 0x88, 0xba, 0x1c, 0x92,
	0x3c, 0xdc, 0xe4, 0xaa, 0xd5, 0xd5, 0xde, 0x3c, 0xf6, 0x7c, 0x79, 0x71, 0x79, 0x4e, 0x7a, 0x2c,
	0x68, 0x4a, 0x2a, 0xb9, 0x23, 0xe9, 0xc0, 0x39, 0x17, 0x4d, 0x8f, 0xb5, 0x44, 0x7f, 0xd8, 0xe4,
	0x34, 0x62, 0x7c, 0xf4, 0xca, 0x7b, 0x3c, 0xea, 0x35, 0x13, 0xd7, 0x2f, 0x4b, 0xf5, 0xd3, 0xfe,
	0xf0, 0xb5, 0x0a, 0xae, 0x8e, 0xd6, 0x1d, 0xbd, 0xf4, 0x3d, 0x36, 0x38, 0xff, 0x2f, 0x7e, 0xf3,
	0xdb, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xc3, 0x2f, 0x8a, 0x23, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegisterClient is the client API for Register service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegisterClient interface {
	DoServiceRegister(ctx context.Context, in *Services, opts ...grpc.CallOption) (*ServiceRegisterMapping, error)
	DoServiceInstanceRegister(ctx context.Context, in *ServiceInstances, opts ...grpc.CallOption) (*ServiceInstanceRegisterMapping, error)
	DoEndpointRegister(ctx context.Context, in *Enpoints, opts ...grpc.CallOption) (*EndpointMapping, error)
	DoNetworkAddressRegister(ctx context.Context, in *NetAddresses, opts ...grpc.CallOption) (*NetAddressMapping, error)
	DoServiceAndNetworkAddressMappingRegister(ctx context.Context, in *ServiceAndNetworkAddressMappings, opts ...grpc.CallOption) (*common.Commands, error)
}

type registerClient struct {
	cc *grpc.ClientConn
}

func NewRegisterClient(cc *grpc.ClientConn) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) DoServiceRegister(ctx context.Context, in *Services, opts ...grpc.CallOption) (*ServiceRegisterMapping, error) {
	out := new(ServiceRegisterMapping)
	err := c.cc.Invoke(ctx, "/Register/doServiceRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) DoServiceInstanceRegister(ctx context.Context, in *ServiceInstances, opts ...grpc.CallOption) (*ServiceInstanceRegisterMapping, error) {
	out := new(ServiceInstanceRegisterMapping)
	err := c.cc.Invoke(ctx, "/Register/doServiceInstanceRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) DoEndpointRegister(ctx context.Context, in *Enpoints, opts ...grpc.CallOption) (*EndpointMapping, error) {
	out := new(EndpointMapping)
	err := c.cc.Invoke(ctx, "/Register/doEndpointRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) DoNetworkAddressRegister(ctx context.Context, in *NetAddresses, opts ...grpc.CallOption) (*NetAddressMapping, error) {
	out := new(NetAddressMapping)
	err := c.cc.Invoke(ctx, "/Register/doNetworkAddressRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) DoServiceAndNetworkAddressMappingRegister(ctx context.Context, in *ServiceAndNetworkAddressMappings, opts ...grpc.CallOption) (*common.Commands, error) {
	out := new(common.Commands)
	err := c.cc.Invoke(ctx, "/Register/doServiceAndNetworkAddressMappingRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServer is the server API for Register service.
type RegisterServer interface {
	DoServiceRegister(context.Context, *Services) (*ServiceRegisterMapping, error)
	DoServiceInstanceRegister(context.Context, *ServiceInstances) (*ServiceInstanceRegisterMapping, error)
	DoEndpointRegister(context.Context, *Enpoints) (*EndpointMapping, error)
	DoNetworkAddressRegister(context.Context, *NetAddresses) (*NetAddressMapping, error)
	DoServiceAndNetworkAddressMappingRegister(context.Context, *ServiceAndNetworkAddressMappings) (*common.Commands, error)
}

// UnimplementedRegisterServer can be embedded to have forward compatible implementations.
type UnimplementedRegisterServer struct {
}

func (*UnimplementedRegisterServer) DoServiceRegister(ctx context.Context, req *Services) (*ServiceRegisterMapping, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoServiceRegister not implemented")
}
func (*UnimplementedRegisterServer) DoServiceInstanceRegister(ctx context.Context, req *ServiceInstances) (*ServiceInstanceRegisterMapping, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoServiceInstanceRegister not implemented")
}
func (*UnimplementedRegisterServer) DoEndpointRegister(ctx context.Context, req *Enpoints) (*EndpointMapping, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoEndpointRegister not implemented")
}
func (*UnimplementedRegisterServer) DoNetworkAddressRegister(ctx context.Context, req *NetAddresses) (*NetAddressMapping, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoNetworkAddressRegister not implemented")
}
func (*UnimplementedRegisterServer) DoServiceAndNetworkAddressMappingRegister(ctx context.Context, req *ServiceAndNetworkAddressMappings) (*common.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoServiceAndNetworkAddressMappingRegister not implemented")
}

func RegisterRegisterServer(s *grpc.Server, srv RegisterServer) {
	s.RegisterService(&_Register_serviceDesc, srv)
}

func _Register_DoServiceRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Services)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).DoServiceRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Register/DoServiceRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).DoServiceRegister(ctx, req.(*Services))
	}
	return interceptor(ctx, in, info, handler)
}

func _Register_DoServiceInstanceRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInstances)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).DoServiceInstanceRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Register/DoServiceInstanceRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).DoServiceInstanceRegister(ctx, req.(*ServiceInstances))
	}
	return interceptor(ctx, in, info, handler)
}

func _Register_DoEndpointRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Enpoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).DoEndpointRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Register/DoEndpointRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).DoEndpointRegister(ctx, req.(*Enpoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Register_DoNetworkAddressRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetAddresses)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).DoNetworkAddressRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Register/DoNetworkAddressRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).DoNetworkAddressRegister(ctx, req.(*NetAddresses))
	}
	return interceptor(ctx, in, info, handler)
}

func _Register_DoServiceAndNetworkAddressMappingRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceAndNetworkAddressMappings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).DoServiceAndNetworkAddressMappingRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Register/DoServiceAndNetworkAddressMappingRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).DoServiceAndNetworkAddressMappingRegister(ctx, req.(*ServiceAndNetworkAddressMappings))
	}
	return interceptor(ctx, in, info, handler)
}

var _Register_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Register",
	HandlerType: (*RegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "doServiceRegister",
			Handler:    _Register_DoServiceRegister_Handler,
		},
		{
			MethodName: "doServiceInstanceRegister",
			Handler:    _Register_DoServiceInstanceRegister_Handler,
		},
		{
			MethodName: "doEndpointRegister",
			Handler:    _Register_DoEndpointRegister_Handler,
		},
		{
			MethodName: "doNetworkAddressRegister",
			Handler:    _Register_DoNetworkAddressRegister_Handler,
		},
		{
			MethodName: "doServiceAndNetworkAddressMappingRegister",
			Handler:    _Register_DoServiceAndNetworkAddressMappingRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register/Register.proto",
}
