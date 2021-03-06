// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.3
// source: directory_items.proto

package discoursepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type DirectoryItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirectoryItems []*DirectoryItems_DirectoryItems `protobuf:"bytes,1,rep,name=directory_items,json=directoryItems,proto3" json:"directory_items,omitempty"`
	Meta           *DirectoryItems_Meta             `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *DirectoryItems) Reset() {
	*x = DirectoryItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directory_items_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryItems) ProtoMessage() {}

func (x *DirectoryItems) ProtoReflect() protoreflect.Message {
	mi := &file_directory_items_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryItems.ProtoReflect.Descriptor instead.
func (*DirectoryItems) Descriptor() ([]byte, []int) {
	return file_directory_items_proto_rawDescGZIP(), []int{0}
}

func (x *DirectoryItems) GetDirectoryItems() []*DirectoryItems_DirectoryItems {
	if x != nil {
		return x.DirectoryItems
	}
	return nil
}

func (x *DirectoryItems) GetMeta() *DirectoryItems_Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type DirectoryItems_DirectoryItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32                               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TimeRead      int32                               `protobuf:"varint,2,opt,name=time_read,json=timeRead,proto3" json:"time_read,omitempty"`
	LikesReceived int32                               `protobuf:"varint,3,opt,name=likes_received,json=likesReceived,proto3" json:"likes_received,omitempty"`
	LikesGiven    int32                               `protobuf:"varint,4,opt,name=likes_given,json=likesGiven,proto3" json:"likes_given,omitempty"`
	TopicsEntered int32                               `protobuf:"varint,5,opt,name=topics_entered,json=topicsEntered,proto3" json:"topics_entered,omitempty"`
	TopicCount    int32                               `protobuf:"varint,6,opt,name=topic_count,json=topicCount,proto3" json:"topic_count,omitempty"`
	PostCount     int32                               `protobuf:"varint,7,opt,name=post_count,json=postCount,proto3" json:"post_count,omitempty"`
	PostsRead     int32                               `protobuf:"varint,8,opt,name=posts_read,json=postsRead,proto3" json:"posts_read,omitempty"`
	DaysVisited   int32                               `protobuf:"varint,9,opt,name=days_visited,json=daysVisited,proto3" json:"days_visited,omitempty"`
	User          *DirectoryItems_DirectoryItems_User `protobuf:"bytes,10,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *DirectoryItems_DirectoryItems) Reset() {
	*x = DirectoryItems_DirectoryItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directory_items_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryItems_DirectoryItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryItems_DirectoryItems) ProtoMessage() {}

func (x *DirectoryItems_DirectoryItems) ProtoReflect() protoreflect.Message {
	mi := &file_directory_items_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryItems_DirectoryItems.ProtoReflect.Descriptor instead.
func (*DirectoryItems_DirectoryItems) Descriptor() ([]byte, []int) {
	return file_directory_items_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DirectoryItems_DirectoryItems) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DirectoryItems_DirectoryItems) GetTimeRead() int32 {
	if x != nil {
		return x.TimeRead
	}
	return 0
}

func (x *DirectoryItems_DirectoryItems) GetLikesReceived() int32 {
	if x != nil {
		return x.LikesReceived
	}
	return 0
}

func (x *DirectoryItems_DirectoryItems) GetLikesGiven() int32 {
	if x != nil {
		return x.LikesGiven
	}
	return 0
}

func (x *DirectoryItems_DirectoryItems) GetTopicsEntered() int32 {
	if x != nil {
		return x.TopicsEntered
	}
	return 0
}

func (x *DirectoryItems_DirectoryItems) GetTopicCount() int32 {
	if x != nil {
		return x.TopicCount
	}
	return 0
}

func (x *DirectoryItems_DirectoryItems) GetPostCount() int32 {
	if x != nil {
		return x.PostCount
	}
	return 0
}

func (x *DirectoryItems_DirectoryItems) GetPostsRead() int32 {
	if x != nil {
		return x.PostsRead
	}
	return 0
}

func (x *DirectoryItems_DirectoryItems) GetDaysVisited() int32 {
	if x != nil {
		return x.DaysVisited
	}
	return 0
}

func (x *DirectoryItems_DirectoryItems) GetUser() *DirectoryItems_DirectoryItems_User {
	if x != nil {
		return x.User
	}
	return nil
}

type DirectoryItems_Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastUpdatedAt           *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=last_updated_at,json=lastUpdatedAt,proto3" json:"last_updated_at,omitempty"`
	TotalRowsDirectoryItems int32                  `protobuf:"varint,2,opt,name=total_rows_directory_items,json=totalRowsDirectoryItems,proto3" json:"total_rows_directory_items,omitempty"`
	LoadMoreDirectoryItems  string                 `protobuf:"bytes,3,opt,name=load_more_directory_items,json=loadMoreDirectoryItems,proto3" json:"load_more_directory_items,omitempty"`
}

func (x *DirectoryItems_Meta) Reset() {
	*x = DirectoryItems_Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directory_items_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryItems_Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryItems_Meta) ProtoMessage() {}

func (x *DirectoryItems_Meta) ProtoReflect() protoreflect.Message {
	mi := &file_directory_items_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryItems_Meta.ProtoReflect.Descriptor instead.
func (*DirectoryItems_Meta) Descriptor() ([]byte, []int) {
	return file_directory_items_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DirectoryItems_Meta) GetLastUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdatedAt
	}
	return nil
}

func (x *DirectoryItems_Meta) GetTotalRowsDirectoryItems() int32 {
	if x != nil {
		return x.TotalRowsDirectoryItems
	}
	return 0
}

func (x *DirectoryItems_Meta) GetLoadMoreDirectoryItems() string {
	if x != nil {
		return x.LoadMoreDirectoryItems
	}
	return ""
}

type DirectoryItems_DirectoryItems_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username       string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AvatarTemplate string `protobuf:"bytes,4,opt,name=avatar_template,json=avatarTemplate,proto3" json:"avatar_template,omitempty"`
	Title          string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *DirectoryItems_DirectoryItems_User) Reset() {
	*x = DirectoryItems_DirectoryItems_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_directory_items_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryItems_DirectoryItems_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryItems_DirectoryItems_User) ProtoMessage() {}

func (x *DirectoryItems_DirectoryItems_User) ProtoReflect() protoreflect.Message {
	mi := &file_directory_items_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryItems_DirectoryItems_User.ProtoReflect.Descriptor instead.
func (*DirectoryItems_DirectoryItems_User) Descriptor() ([]byte, []int) {
	return file_directory_items_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *DirectoryItems_DirectoryItems_User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DirectoryItems_DirectoryItems_User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DirectoryItems_DirectoryItems_User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DirectoryItems_DirectoryItems_User) GetAvatarTemplate() string {
	if x != nil {
		return x.AvatarTemplate
	}
	return ""
}

func (x *DirectoryItems_DirectoryItems_User) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_directory_items_proto protoreflect.FileDescriptor

var file_directory_items_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x06, 0x0a,
	0x0e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x5a, 0x0a, 0x0f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x0e, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3b, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x82, 0x04, 0x0a, 0x0e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x61, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x69, 0x6b, 0x65,
	0x73, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x47, 0x69, 0x76, 0x65, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x6f,
	0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x6f, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x76,
	0x69, 0x73, 0x69, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x61,
	0x79, 0x73, 0x56, 0x69, 0x73, 0x69, 0x74, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x85, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x1a, 0xc2, 0x01,
	0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x1a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x6f, 0x77, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x6d, 0x6f, 0x72, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6c, 0x6f, 0x61, 0x64,
	0x4d, 0x6f, 0x72, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x64, 0x6d, 0x61, 0x6e, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_directory_items_proto_rawDescOnce sync.Once
	file_directory_items_proto_rawDescData = file_directory_items_proto_rawDesc
)

func file_directory_items_proto_rawDescGZIP() []byte {
	file_directory_items_proto_rawDescOnce.Do(func() {
		file_directory_items_proto_rawDescData = protoimpl.X.CompressGZIP(file_directory_items_proto_rawDescData)
	})
	return file_directory_items_proto_rawDescData
}

var file_directory_items_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_directory_items_proto_goTypes = []interface{}{
	(*DirectoryItems)(nil),                     // 0: discourse.protobuf.DirectoryItems
	(*DirectoryItems_DirectoryItems)(nil),      // 1: discourse.protobuf.DirectoryItems.DirectoryItems
	(*DirectoryItems_Meta)(nil),                // 2: discourse.protobuf.DirectoryItems.Meta
	(*DirectoryItems_DirectoryItems_User)(nil), // 3: discourse.protobuf.DirectoryItems.DirectoryItems.User
	(*timestamppb.Timestamp)(nil),              // 4: google.protobuf.Timestamp
}
var file_directory_items_proto_depIdxs = []int32{
	1, // 0: discourse.protobuf.DirectoryItems.directory_items:type_name -> discourse.protobuf.DirectoryItems.DirectoryItems
	2, // 1: discourse.protobuf.DirectoryItems.meta:type_name -> discourse.protobuf.DirectoryItems.Meta
	3, // 2: discourse.protobuf.DirectoryItems.DirectoryItems.user:type_name -> discourse.protobuf.DirectoryItems.DirectoryItems.User
	4, // 3: discourse.protobuf.DirectoryItems.Meta.last_updated_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_directory_items_proto_init() }
func file_directory_items_proto_init() {
	if File_directory_items_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_directory_items_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryItems); i {
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
		file_directory_items_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryItems_DirectoryItems); i {
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
		file_directory_items_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryItems_Meta); i {
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
		file_directory_items_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryItems_DirectoryItems_User); i {
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
			RawDescriptor: file_directory_items_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_directory_items_proto_goTypes,
		DependencyIndexes: file_directory_items_proto_depIdxs,
		MessageInfos:      file_directory_items_proto_msgTypes,
	}.Build()
	File_directory_items_proto = out.File
	file_directory_items_proto_rawDesc = nil
	file_directory_items_proto_goTypes = nil
	file_directory_items_proto_depIdxs = nil
}
