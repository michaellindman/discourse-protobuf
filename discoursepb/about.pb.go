// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.3
// source: about.proto

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

type About struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	About *About_About `protobuf:"bytes,1,opt,name=about,proto3" json:"about,omitempty"`
}

func (x *About) Reset() {
	*x = About{}
	if protoimpl.UnsafeEnabled {
		mi := &file_about_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *About) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*About) ProtoMessage() {}

func (x *About) ProtoReflect() protoreflect.Message {
	mi := &file_about_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use About.ProtoReflect.Descriptor instead.
func (*About) Descriptor() ([]byte, []int) {
	return file_about_proto_rawDescGZIP(), []int{0}
}

func (x *About) GetAbout() *About_About {
	if x != nil {
		return x.About
	}
	return nil
}

type About_About struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats              *About_About_Stats                `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	Description        string                            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Title              string                            `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Locale             string                            `protobuf:"bytes,4,opt,name=locale,proto3" json:"locale,omitempty"`
	Version            string                            `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Https              bool                              `protobuf:"varint,6,opt,name=https,proto3" json:"https,omitempty"`
	CanSeeAboutStats   bool                              `protobuf:"varint,7,opt,name=can_see_about_stats,json=canSeeAboutStats,proto3" json:"can_see_about_stats,omitempty"`
	Moderators         []*About_About_Moderators         `protobuf:"bytes,8,rep,name=moderators,proto3" json:"moderators,omitempty"`
	Admins             []*About_About_Admins             `protobuf:"bytes,9,rep,name=admins,proto3" json:"admins,omitempty"`
	CategoryModerators []*About_About_CategoryModerators `protobuf:"bytes,10,rep,name=category_moderators,json=categoryModerators,proto3" json:"category_moderators,omitempty"`
}

func (x *About_About) Reset() {
	*x = About_About{}
	if protoimpl.UnsafeEnabled {
		mi := &file_about_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *About_About) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*About_About) ProtoMessage() {}

func (x *About_About) ProtoReflect() protoreflect.Message {
	mi := &file_about_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use About_About.ProtoReflect.Descriptor instead.
func (*About_About) Descriptor() ([]byte, []int) {
	return file_about_proto_rawDescGZIP(), []int{0, 0}
}

func (x *About_About) GetStats() *About_About_Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *About_About) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *About_About) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *About_About) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *About_About) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *About_About) GetHttps() bool {
	if x != nil {
		return x.Https
	}
	return false
}

func (x *About_About) GetCanSeeAboutStats() bool {
	if x != nil {
		return x.CanSeeAboutStats
	}
	return false
}

func (x *About_About) GetModerators() []*About_About_Moderators {
	if x != nil {
		return x.Moderators
	}
	return nil
}

func (x *About_About) GetAdmins() []*About_About_Admins {
	if x != nil {
		return x.Admins
	}
	return nil
}

func (x *About_About) GetCategoryModerators() []*About_About_CategoryModerators {
	if x != nil {
		return x.CategoryModerators
	}
	return nil
}

type About_About_Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicCount         int32 `protobuf:"varint,1,opt,name=topic_count,json=topicCount,proto3" json:"topic_count,omitempty"`
	PostCount          int32 `protobuf:"varint,2,opt,name=post_count,json=postCount,proto3" json:"post_count,omitempty"`
	UserCount          int32 `protobuf:"varint,3,opt,name=user_count,json=userCount,proto3" json:"user_count,omitempty"`
	Topics_7Days       int32 `protobuf:"varint,4,opt,name=topics_7_days,json=topics7Days,proto3" json:"topics_7_days,omitempty"`
	Topics_30Days      int32 `protobuf:"varint,5,opt,name=topics_30_days,json=topics30Days,proto3" json:"topics_30_days,omitempty"`
	Posts_7Days        int32 `protobuf:"varint,6,opt,name=posts_7_days,json=posts7Days,proto3" json:"posts_7_days,omitempty"`
	Posts_30Days       int32 `protobuf:"varint,7,opt,name=posts_30_days,json=posts30Days,proto3" json:"posts_30_days,omitempty"`
	Users_7Days        int32 `protobuf:"varint,8,opt,name=users_7_days,json=users7Days,proto3" json:"users_7_days,omitempty"`
	Users_30Days       int32 `protobuf:"varint,9,opt,name=users_30_days,json=users30Days,proto3" json:"users_30_days,omitempty"`
	ActiveUsers_7Days  int32 `protobuf:"varint,10,opt,name=active_users_7_days,json=activeUsers7Days,proto3" json:"active_users_7_days,omitempty"`
	ActiveUsers_30Days int32 `protobuf:"varint,11,opt,name=active_users_30_days,json=activeUsers30Days,proto3" json:"active_users_30_days,omitempty"`
	LikeCount          int32 `protobuf:"varint,12,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	Likes_7Days        int32 `protobuf:"varint,13,opt,name=likes_7_days,json=likes7Days,proto3" json:"likes_7_days,omitempty"`
	Likes_30Days       int32 `protobuf:"varint,14,opt,name=likes_30_days,json=likes30Days,proto3" json:"likes_30_days,omitempty"`
}

func (x *About_About_Stats) Reset() {
	*x = About_About_Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_about_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *About_About_Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*About_About_Stats) ProtoMessage() {}

func (x *About_About_Stats) ProtoReflect() protoreflect.Message {
	mi := &file_about_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use About_About_Stats.ProtoReflect.Descriptor instead.
func (*About_About_Stats) Descriptor() ([]byte, []int) {
	return file_about_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *About_About_Stats) GetTopicCount() int32 {
	if x != nil {
		return x.TopicCount
	}
	return 0
}

func (x *About_About_Stats) GetPostCount() int32 {
	if x != nil {
		return x.PostCount
	}
	return 0
}

func (x *About_About_Stats) GetUserCount() int32 {
	if x != nil {
		return x.UserCount
	}
	return 0
}

func (x *About_About_Stats) GetTopics_7Days() int32 {
	if x != nil {
		return x.Topics_7Days
	}
	return 0
}

func (x *About_About_Stats) GetTopics_30Days() int32 {
	if x != nil {
		return x.Topics_30Days
	}
	return 0
}

func (x *About_About_Stats) GetPosts_7Days() int32 {
	if x != nil {
		return x.Posts_7Days
	}
	return 0
}

func (x *About_About_Stats) GetPosts_30Days() int32 {
	if x != nil {
		return x.Posts_30Days
	}
	return 0
}

func (x *About_About_Stats) GetUsers_7Days() int32 {
	if x != nil {
		return x.Users_7Days
	}
	return 0
}

func (x *About_About_Stats) GetUsers_30Days() int32 {
	if x != nil {
		return x.Users_30Days
	}
	return 0
}

func (x *About_About_Stats) GetActiveUsers_7Days() int32 {
	if x != nil {
		return x.ActiveUsers_7Days
	}
	return 0
}

func (x *About_About_Stats) GetActiveUsers_30Days() int32 {
	if x != nil {
		return x.ActiveUsers_30Days
	}
	return 0
}

func (x *About_About_Stats) GetLikeCount() int32 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *About_About_Stats) GetLikes_7Days() int32 {
	if x != nil {
		return x.Likes_7Days
	}
	return 0
}

func (x *About_About_Stats) GetLikes_30Days() int32 {
	if x != nil {
		return x.Likes_30Days
	}
	return 0
}

type About_About_Moderators struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username       string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name           string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AvatarTemplate string                 `protobuf:"bytes,4,opt,name=avatar_template,json=avatarTemplate,proto3" json:"avatar_template,omitempty"`
	Title          string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	LastSeenAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_seen_at,json=lastSeenAt,proto3" json:"last_seen_at,omitempty"`
}

func (x *About_About_Moderators) Reset() {
	*x = About_About_Moderators{}
	if protoimpl.UnsafeEnabled {
		mi := &file_about_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *About_About_Moderators) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*About_About_Moderators) ProtoMessage() {}

func (x *About_About_Moderators) ProtoReflect() protoreflect.Message {
	mi := &file_about_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use About_About_Moderators.ProtoReflect.Descriptor instead.
func (*About_About_Moderators) Descriptor() ([]byte, []int) {
	return file_about_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *About_About_Moderators) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *About_About_Moderators) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *About_About_Moderators) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *About_About_Moderators) GetAvatarTemplate() string {
	if x != nil {
		return x.AvatarTemplate
	}
	return ""
}

func (x *About_About_Moderators) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *About_About_Moderators) GetLastSeenAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSeenAt
	}
	return nil
}

type About_About_Admins struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username       string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name           string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AvatarTemplate string                 `protobuf:"bytes,4,opt,name=avatar_template,json=avatarTemplate,proto3" json:"avatar_template,omitempty"`
	Title          string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	LastSeenAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_seen_at,json=lastSeenAt,proto3" json:"last_seen_at,omitempty"`
}

func (x *About_About_Admins) Reset() {
	*x = About_About_Admins{}
	if protoimpl.UnsafeEnabled {
		mi := &file_about_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *About_About_Admins) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*About_About_Admins) ProtoMessage() {}

func (x *About_About_Admins) ProtoReflect() protoreflect.Message {
	mi := &file_about_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use About_About_Admins.ProtoReflect.Descriptor instead.
func (*About_About_Admins) Descriptor() ([]byte, []int) {
	return file_about_proto_rawDescGZIP(), []int{0, 0, 2}
}

func (x *About_About_Admins) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *About_About_Admins) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *About_About_Admins) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *About_About_Admins) GetAvatarTemplate() string {
	if x != nil {
		return x.AvatarTemplate
	}
	return ""
}

func (x *About_About_Admins) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *About_About_Admins) GetLastSeenAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSeenAt
	}
	return nil
}

type About_About_CategoryModerators struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username       string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name           string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AvatarTemplate string                 `protobuf:"bytes,4,opt,name=avatar_template,json=avatarTemplate,proto3" json:"avatar_template,omitempty"`
	Title          string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	LastSeenAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_seen_at,json=lastSeenAt,proto3" json:"last_seen_at,omitempty"`
}

func (x *About_About_CategoryModerators) Reset() {
	*x = About_About_CategoryModerators{}
	if protoimpl.UnsafeEnabled {
		mi := &file_about_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *About_About_CategoryModerators) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*About_About_CategoryModerators) ProtoMessage() {}

func (x *About_About_CategoryModerators) ProtoReflect() protoreflect.Message {
	mi := &file_about_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use About_About_CategoryModerators.ProtoReflect.Descriptor instead.
func (*About_About_CategoryModerators) Descriptor() ([]byte, []int) {
	return file_about_proto_rawDescGZIP(), []int{0, 0, 3}
}

func (x *About_About_CategoryModerators) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *About_About_CategoryModerators) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *About_About_CategoryModerators) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *About_About_CategoryModerators) GetAvatarTemplate() string {
	if x != nil {
		return x.AvatarTemplate
	}
	return ""
}

func (x *About_About_CategoryModerators) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *About_About_CategoryModerators) GetLastSeenAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSeenAt
	}
	return nil
}

var File_about_proto protoreflect.FileDescriptor

var file_about_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x91, 0x0d, 0x0a, 0x05, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x12, 0x35, 0x0a, 0x05,
	0x61, 0x62, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x52, 0x05, 0x61, 0x62,
	0x6f, 0x75, 0x74, 0x1a, 0xd0, 0x0c, 0x0a, 0x05, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x12, 0x3b, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x74, 0x74, 0x70, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x74, 0x74, 0x70, 0x73, 0x12, 0x2d, 0x0a, 0x13, 0x63, 0x61,
	0x6e, 0x5f, 0x73, 0x65, 0x65, 0x5f, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x63, 0x61, 0x6e, 0x53, 0x65, 0x65, 0x41,
	0x62, 0x6f, 0x75, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x0a, 0x6d, 0x6f, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74,
	0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x52, 0x06, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x73, 0x12, 0x63, 0x0a, 0x13, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x2e, 0x41, 0x62,
	0x6f, 0x75, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x12, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x81, 0x04, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x5f, 0x37, 0x5f,
	0x64, 0x61, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x37, 0x44, 0x61, 0x79, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x5f, 0x33, 0x30, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x33, 0x30, 0x44, 0x61, 0x79, 0x73, 0x12, 0x20, 0x0a,
	0x0c, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x37, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x37, 0x44, 0x61, 0x79, 0x73, 0x12,
	0x22, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x33, 0x30, 0x5f, 0x64, 0x61, 0x79, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x33, 0x30, 0x44,
	0x61, 0x79, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x37, 0x5f, 0x64,
	0x61, 0x79, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x37, 0x44, 0x61, 0x79, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x33,
	0x30, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x33, 0x30, 0x44, 0x61, 0x79, 0x73, 0x12, 0x2d, 0x0a, 0x13, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x37, 0x5f, 0x64, 0x61, 0x79, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x37, 0x44, 0x61, 0x79, 0x73, 0x12, 0x2f, 0x0a, 0x14, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x33, 0x30, 0x5f, 0x64, 0x61, 0x79, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x33, 0x30, 0x44, 0x61, 0x79, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6b,
	0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c,
	0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x69, 0x6b, 0x65,
	0x73, 0x5f, 0x37, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x6c, 0x69, 0x6b, 0x65, 0x73, 0x37, 0x44, 0x61, 0x79, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x5f, 0x33, 0x30, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x33, 0x30, 0x44, 0x61, 0x79, 0x73, 0x1a, 0xc9,
	0x01, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x41, 0x74, 0x1a, 0xc5, 0x01, 0x0a, 0x06, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65,
	0x6e, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e,
	0x41, 0x74, 0x1a, 0xd1, 0x01, 0x0a, 0x12, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d,
	0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x53, 0x65, 0x65, 0x6e, 0x41, 0x74, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x64,
	0x6d, 0x61, 0x6e, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_about_proto_rawDescOnce sync.Once
	file_about_proto_rawDescData = file_about_proto_rawDesc
)

func file_about_proto_rawDescGZIP() []byte {
	file_about_proto_rawDescOnce.Do(func() {
		file_about_proto_rawDescData = protoimpl.X.CompressGZIP(file_about_proto_rawDescData)
	})
	return file_about_proto_rawDescData
}

var file_about_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_about_proto_goTypes = []interface{}{
	(*About)(nil),                          // 0: discourse.protobuf.About
	(*About_About)(nil),                    // 1: discourse.protobuf.About.About
	(*About_About_Stats)(nil),              // 2: discourse.protobuf.About.About.Stats
	(*About_About_Moderators)(nil),         // 3: discourse.protobuf.About.About.Moderators
	(*About_About_Admins)(nil),             // 4: discourse.protobuf.About.About.Admins
	(*About_About_CategoryModerators)(nil), // 5: discourse.protobuf.About.About.CategoryModerators
	(*timestamppb.Timestamp)(nil),          // 6: google.protobuf.Timestamp
}
var file_about_proto_depIdxs = []int32{
	1, // 0: discourse.protobuf.About.about:type_name -> discourse.protobuf.About.About
	2, // 1: discourse.protobuf.About.About.stats:type_name -> discourse.protobuf.About.About.Stats
	3, // 2: discourse.protobuf.About.About.moderators:type_name -> discourse.protobuf.About.About.Moderators
	4, // 3: discourse.protobuf.About.About.admins:type_name -> discourse.protobuf.About.About.Admins
	5, // 4: discourse.protobuf.About.About.category_moderators:type_name -> discourse.protobuf.About.About.CategoryModerators
	6, // 5: discourse.protobuf.About.About.Moderators.last_seen_at:type_name -> google.protobuf.Timestamp
	6, // 6: discourse.protobuf.About.About.Admins.last_seen_at:type_name -> google.protobuf.Timestamp
	6, // 7: discourse.protobuf.About.About.CategoryModerators.last_seen_at:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_about_proto_init() }
func file_about_proto_init() {
	if File_about_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_about_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*About); i {
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
		file_about_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*About_About); i {
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
		file_about_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*About_About_Stats); i {
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
		file_about_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*About_About_Moderators); i {
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
		file_about_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*About_About_Admins); i {
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
		file_about_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*About_About_CategoryModerators); i {
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
			RawDescriptor: file_about_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_about_proto_goTypes,
		DependencyIndexes: file_about_proto_depIdxs,
		MessageInfos:      file_about_proto_msgTypes,
	}.Build()
	File_about_proto = out.File
	file_about_proto_rawDesc = nil
	file_about_proto_goTypes = nil
	file_about_proto_depIdxs = nil
}
