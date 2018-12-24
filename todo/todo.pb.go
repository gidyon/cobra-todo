// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Task struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Done                 bool     `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
	TaskId               int32    `protobuf:"varint,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Task) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Task) GetTaskId() int32 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

type Todo struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type TaskId struct {
	TaskId               int32    `protobuf:"varint,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskId) Reset()         { *m = TaskId{} }
func (m *TaskId) String() string { return proto.CompactTextString(m) }
func (*TaskId) ProtoMessage()    {}
func (*TaskId) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *TaskId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskId.Unmarshal(m, b)
}
func (m *TaskId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskId.Marshal(b, m, deterministic)
}
func (m *TaskId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskId.Merge(m, src)
}
func (m *TaskId) XXX_Size() int {
	return xxx_messageInfo_TaskId.Size(m)
}
func (m *TaskId) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskId.DiscardUnknown(m)
}

var xxx_messageInfo_TaskId proto.InternalMessageInfo

func (m *TaskId) GetTaskId() int32 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func init() {
	proto.RegisterType((*Task)(nil), "todo.Task")
	proto.RegisterType((*Todo)(nil), "todo.Todo")
	proto.RegisterType((*Void)(nil), "todo.Void")
	proto.RegisterType((*TaskId)(nil), "todo.TaskId")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0x2f, 0x6e, 0x2e, 0xea, 0x9c, 0xd5, 0x34, 0x86, 0x6b, 0x8c, 0xc1, 0x22, 0xd5, 0x15,
	0xe7, 0x13, 0x08, 0x07, 0x72, 0xa0, 0x4d, 0x58, 0x6c, 0x25, 0x32, 0x41, 0x96, 0x93, 0x1b, 0xd9,
	0xa4, 0xf0, 0x51, 0x7c, 0x5c, 0x99, 0xa8, 0xb8, 0x08, 0xdb, 0xfd, 0x99, 0xfc, 0xf9, 0x3e, 0x26,
	0x00, 0x95, 0x89, 0x37, 0xef, 0x23, 0x57, 0x46, 0x2d, 0xd9, 0xdf, 0x83, 0xee, 0x53, 0x39, 0x20,
	0x82, 0xae, 0xf9, 0xa3, 0x5a, 0xe5, 0x54, 0x38, 0x8f, 0x2d, 0xcb, 0x8c, 0xf8, 0x98, 0xed, 0x89,
	0x53, 0xe1, 0x2c, 0xb6, 0x8c, 0x97, 0x70, 0x5a, 0x53, 0x39, 0x3c, 0x0f, 0x64, 0x3b, 0xa7, 0xc2,
	0x32, 0x1a, 0x39, 0xee, 0xc9, 0x07, 0xd0, 0x3d, 0x13, 0xa3, 0x83, 0xa5, 0x4c, 0x8a, 0x55, 0xae,
	0x0b, 0xab, 0x2d, 0x6c, 0x9a, 0x52, 0x1c, 0xf1, 0xfb, 0xc2, 0x1b, 0xd0, 0x4f, 0x3c, 0x90, 0xbf,
	0x06, 0xd3, 0xb7, 0xb7, 0xb3, 0xd0, 0xed, 0xa7, 0x82, 0x95, 0x50, 0x1f, 0xd3, 0x31, 0xbd, 0xe6,
	0x11, 0x1d, 0xe8, 0x87, 0xa1, 0x54, 0xfc, 0xa1, 0x0a, 0x66, 0xfd, 0x6b, 0x90, 0x6d, 0x16, 0x78,
	0x05, 0xdd, 0x1d, 0x11, 0x4e, 0xb4, 0xeb, 0x49, 0xd9, 0x2f, 0xf0, 0x06, 0xcc, 0x2e, 0xbf, 0xe5,
	0x9a, 0xf1, 0xe2, 0xaf, 0xb3, 0xa7, 0x7f, 0x2d, 0x07, 0x7a, 0x27, 0xeb, 0xce, 0x72, 0x5e, 0x4c,
	0xfb, 0xc5, 0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0x47, 0xa3, 0x81, 0x53, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoManagerClient is the client API for TodoManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoManagerClient interface {
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Todo, error)
	Add(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *TaskId, opts ...grpc.CallOption) (*Void, error)
	Done(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Void, error)
}

type todoManagerClient struct {
	cc *grpc.ClientConn
}

func NewTodoManagerClient(cc *grpc.ClientConn) TodoManagerClient {
	return &todoManagerClient{cc}
}

func (c *todoManagerClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todo.TodoManager/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) Add(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/todo.TodoManager/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) Delete(ctx context.Context, in *TaskId, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/todo.TodoManager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) Done(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/todo.TodoManager/Done", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoManagerServer is the server API for TodoManager service.
type TodoManagerServer interface {
	List(context.Context, *Void) (*Todo, error)
	Add(context.Context, *Task) (*Void, error)
	Delete(context.Context, *TaskId) (*Void, error)
	Done(context.Context, *Task) (*Void, error)
}

func RegisterTodoManagerServer(s *grpc.Server, srv TodoManagerServer) {
	s.RegisterService(&_TodoManager_serviceDesc, srv)
}

func _TodoManager_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoManager/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoManager/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).Add(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).Delete(ctx, req.(*TaskId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_Done_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).Done(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoManager/Done",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).Done(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoManager",
	HandlerType: (*TodoManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TodoManager_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _TodoManager_Add_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TodoManager_Delete_Handler,
		},
		{
			MethodName: "Done",
			Handler:    _TodoManager_Done_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
