// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: example/demo/v1/type/employee.proto

package types

import (
	date "google.golang.org/genproto/googleapis/type/date"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Role int32

const (
	// default role
	Role_USER    Role = 0
	Role_HR      Role = 1
	Role_MANAGER Role = 2
	Role_ADMIN   Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "USER",
		1: "HR",
		2: "MANAGER",
		3: "ADMIN",
	}
	Role_value = map[string]int32{
		"USER":    0,
		"HR":      1,
		"MANAGER": 2,
		"ADMIN":   3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_example_demo_v1_type_employee_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_example_demo_v1_type_employee_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_example_demo_v1_type_employee_proto_rawDescGZIP(), []int{0}
}

type EmployeeProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmpId    string     `protobuf:"bytes,1,opt,name=emp_id,json=empId,proto3" json:"emp_id,omitempty"`
	Name     string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	JoinDate *date.Date `protobuf:"bytes,3,opt,name=join_date,json=joinDate,proto3" json:"join_date,omitempty"`
	Role     Role       `protobuf:"varint,4,opt,name=role,proto3,enum=example.demo.v1.Role" json:"role,omitempty"`
}

func (x *EmployeeProto) Reset() {
	*x = EmployeeProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_demo_v1_type_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeProto) ProtoMessage() {}

func (x *EmployeeProto) ProtoReflect() protoreflect.Message {
	mi := &file_example_demo_v1_type_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeProto.ProtoReflect.Descriptor instead.
func (*EmployeeProto) Descriptor() ([]byte, []int) {
	return file_example_demo_v1_type_employee_proto_rawDescGZIP(), []int{0}
}

func (x *EmployeeProto) GetEmpId() string {
	if x != nil {
		return x.EmpId
	}
	return ""
}

func (x *EmployeeProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EmployeeProto) GetJoinDate() *date.Date {
	if x != nil {
		return x.JoinDate
	}
	return nil
}

func (x *EmployeeProto) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_USER
}

var File_example_demo_v1_type_employee_proto protoreflect.FileDescriptor

var file_example_demo_v1_type_employee_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95,
	0x01, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x0a, 0x06, 0x65, 0x6d, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x6a,
	0x6f, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x65, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x2a, 0x30, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x48, 0x52, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x03, 0x42, 0x36, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x50, 0x01, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_demo_v1_type_employee_proto_rawDescOnce sync.Once
	file_example_demo_v1_type_employee_proto_rawDescData = file_example_demo_v1_type_employee_proto_rawDesc
)

func file_example_demo_v1_type_employee_proto_rawDescGZIP() []byte {
	file_example_demo_v1_type_employee_proto_rawDescOnce.Do(func() {
		file_example_demo_v1_type_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_demo_v1_type_employee_proto_rawDescData)
	})
	return file_example_demo_v1_type_employee_proto_rawDescData
}

var file_example_demo_v1_type_employee_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_example_demo_v1_type_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_example_demo_v1_type_employee_proto_goTypes = []interface{}{
	(Role)(0),             // 0: example.demo.v1.Role
	(*EmployeeProto)(nil), // 1: example.demo.v1.EmployeeProto
	(*date.Date)(nil),     // 2: google.type.Date
}
var file_example_demo_v1_type_employee_proto_depIdxs = []int32{
	2, // 0: example.demo.v1.EmployeeProto.join_date:type_name -> google.type.Date
	0, // 1: example.demo.v1.EmployeeProto.role:type_name -> example.demo.v1.Role
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_example_demo_v1_type_employee_proto_init() }
func file_example_demo_v1_type_employee_proto_init() {
	if File_example_demo_v1_type_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_demo_v1_type_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_demo_v1_type_employee_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_demo_v1_type_employee_proto_goTypes,
		DependencyIndexes: file_example_demo_v1_type_employee_proto_depIdxs,
		EnumInfos:         file_example_demo_v1_type_employee_proto_enumTypes,
		MessageInfos:      file_example_demo_v1_type_employee_proto_msgTypes,
	}.Build()
	File_example_demo_v1_type_employee_proto = out.File
	file_example_demo_v1_type_employee_proto_rawDesc = nil
	file_example_demo_v1_type_employee_proto_goTypes = nil
	file_example_demo_v1_type_employee_proto_depIdxs = nil
}
