// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: vote/v1/service.proto

package pbuserv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VoterId   string                 `protobuf:"bytes,3,opt,name=voter_id,json=voterId,proto3" json:"voter_id,omitempty"`
	Vote      string                 `protobuf:"bytes,4,opt,name=vote,proto3" json:"vote,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_vote_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_vote_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Vote) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Vote) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Vote) GetVoterId() string {
	if x != nil {
		return x.VoterId
	}
	return ""
}

func (x *Vote) GetVote() string {
	if x != nil {
		return x.Vote
	}
	return ""
}

func (x *Vote) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type TotalVotes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TotalVotes int64  `protobuf:"varint,2,opt,name=total_votes,json=totalVotes,proto3" json:"total_votes,omitempty"`
}

func (x *TotalVotes) Reset() {
	*x = TotalVotes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalVotes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalVotes) ProtoMessage() {}

func (x *TotalVotes) ProtoReflect() protoreflect.Message {
	mi := &file_vote_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalVotes.ProtoReflect.Descriptor instead.
func (*TotalVotes) Descriptor() ([]byte, []int) {
	return file_vote_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *TotalVotes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TotalVotes) GetTotalVotes() int64 {
	if x != nil {
		return x.TotalVotes
	}
	return 0
}

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vote *Vote `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vote_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_vote_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *VoteRequest) GetVote() *Vote {
	if x != nil {
		return x.Vote
	}
	return nil
}

type TotalVotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *TotalVotesRequest) Reset() {
	*x = TotalVotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalVotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalVotesRequest) ProtoMessage() {}

func (x *TotalVotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vote_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalVotesRequest.ProtoReflect.Descriptor instead.
func (*TotalVotesRequest) Descriptor() ([]byte, []int) {
	return file_vote_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *TotalVotesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type TotalVotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalVotes *TotalVotes `protobuf:"bytes,1,opt,name=total_votes,json=totalVotes,proto3" json:"total_votes,omitempty"`
}

func (x *TotalVotesResponse) Reset() {
	*x = TotalVotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vote_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalVotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalVotesResponse) ProtoMessage() {}

func (x *TotalVotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vote_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalVotesResponse.ProtoReflect.Descriptor instead.
func (*TotalVotesResponse) Descriptor() ([]byte, []int) {
	return file_vote_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *TotalVotesResponse) GetTotalVotes() *TotalVotes {
	if x != nil {
		return x.TotalVotes
	}
	return nil
}

var File_vote_v1_service_proto protoreflect.FileDescriptor

var file_vote_v1_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99,
	0x01, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76,
	0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x46, 0x0a, 0x0a, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x6f, 0x74,
	0x65, 0x73, 0x22, 0x30, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x04,
	0x76, 0x6f, 0x74, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x6f, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x4a, 0x0a, 0x12, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x6f, 0x74,
	0x65, 0x73, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x32, 0x8e,
	0x01, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36,
	0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x56,
	0x6f, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x53,
	0x4f, 0x4e, 0x53, 0x74, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62,
	0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vote_v1_service_proto_rawDescOnce sync.Once
	file_vote_v1_service_proto_rawDescData = file_vote_v1_service_proto_rawDesc
)

func file_vote_v1_service_proto_rawDescGZIP() []byte {
	file_vote_v1_service_proto_rawDescOnce.Do(func() {
		file_vote_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_vote_v1_service_proto_rawDescData)
	})
	return file_vote_v1_service_proto_rawDescData
}

var file_vote_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vote_v1_service_proto_goTypes = []interface{}{
	(*Vote)(nil),                  // 0: vote.v1.Vote
	(*TotalVotes)(nil),            // 1: vote.v1.TotalVotes
	(*VoteRequest)(nil),           // 2: vote.v1.VoteRequest
	(*TotalVotesRequest)(nil),     // 3: vote.v1.TotalVotesRequest
	(*TotalVotesResponse)(nil),    // 4: vote.v1.TotalVotesResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_vote_v1_service_proto_depIdxs = []int32{
	5, // 0: vote.v1.Vote.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: vote.v1.VoteRequest.vote:type_name -> vote.v1.Vote
	1, // 2: vote.v1.TotalVotesResponse.total_votes:type_name -> vote.v1.TotalVotes
	2, // 3: vote.v1.VoteService.Vote:input_type -> vote.v1.VoteRequest
	3, // 4: vote.v1.VoteService.TotalVotes:input_type -> vote.v1.TotalVotesRequest
	6, // 5: vote.v1.VoteService.Vote:output_type -> google.protobuf.Empty
	4, // 6: vote.v1.VoteService.TotalVotes:output_type -> vote.v1.TotalVotesResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_vote_v1_service_proto_init() }
func file_vote_v1_service_proto_init() {
	if File_vote_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vote_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
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
		file_vote_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalVotes); i {
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
		file_vote_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequest); i {
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
		file_vote_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalVotesRequest); i {
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
		file_vote_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalVotesResponse); i {
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
			RawDescriptor: file_vote_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vote_v1_service_proto_goTypes,
		DependencyIndexes: file_vote_v1_service_proto_depIdxs,
		MessageInfos:      file_vote_v1_service_proto_msgTypes,
	}.Build()
	File_vote_v1_service_proto = out.File
	file_vote_v1_service_proto_rawDesc = nil
	file_vote_v1_service_proto_goTypes = nil
	file_vote_v1_service_proto_depIdxs = nil
}
