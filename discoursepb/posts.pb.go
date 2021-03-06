// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.3
// source: posts.proto

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

type PostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title            string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	TopicId          int32                  `protobuf:"varint,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	Raw              string                 `protobuf:"bytes,3,opt,name=raw,proto3" json:"raw,omitempty"`
	Category         int32                  `protobuf:"varint,4,opt,name=category,proto3" json:"category,omitempty"`
	TargetRecipients string                 `protobuf:"bytes,5,opt,name=target_recipients,json=targetRecipients,proto3" json:"target_recipients,omitempty"`
	Archetype        string                 `protobuf:"bytes,6,opt,name=archetype,proto3" json:"archetype,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PostsRequest) Reset() {
	*x = PostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostsRequest) ProtoMessage() {}

func (x *PostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostsRequest.ProtoReflect.Descriptor instead.
func (*PostsRequest) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{0}
}

func (x *PostsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostsRequest) GetTopicId() int32 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

func (x *PostsRequest) GetRaw() string {
	if x != nil {
		return x.Raw
	}
	return ""
}

func (x *PostsRequest) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *PostsRequest) GetTargetRecipients() string {
	if x != nil {
		return x.TargetRecipients
	}
	return ""
}

func (x *PostsRequest) GetArchetype() string {
	if x != nil {
		return x.Archetype
	}
	return ""
}

func (x *PostsRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type PostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int32                           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               string                          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AvatarTemplate     string                          `protobuf:"bytes,3,opt,name=avatar_template,json=avatarTemplate,proto3" json:"avatar_template,omitempty"`
	CreatedAt          *timestamppb.Timestamp          `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Cooked             string                          `protobuf:"bytes,5,opt,name=cooked,proto3" json:"cooked,omitempty"`
	PostNumber         int32                           `protobuf:"varint,6,opt,name=post_number,json=postNumber,proto3" json:"post_number,omitempty"`
	PostType           int32                           `protobuf:"varint,7,opt,name=post_type,json=postType,proto3" json:"post_type,omitempty"`
	UpdatedAt          string                          `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ReplyCount         int32                           `protobuf:"varint,9,opt,name=reply_count,json=replyCount,proto3" json:"reply_count,omitempty"`
	ReplyToPostNumber  int32                           `protobuf:"varint,10,opt,name=reply_to_post_number,json=replyToPostNumber,proto3" json:"reply_to_post_number,omitempty"`
	QuoteCount         int32                           `protobuf:"varint,11,opt,name=quote_count,json=quoteCount,proto3" json:"quote_count,omitempty"`
	AvgTime            *timestamppb.Timestamp          `protobuf:"bytes,12,opt,name=avg_time,json=avgTime,proto3" json:"avg_time,omitempty"`
	IncomingLinkCount  int32                           `protobuf:"varint,13,opt,name=incoming_link_count,json=incomingLinkCount,proto3" json:"incoming_link_count,omitempty"`
	Reads              int32                           `protobuf:"varint,14,opt,name=reads,proto3" json:"reads,omitempty"`
	Score              float64                         `protobuf:"fixed64,15,opt,name=score,proto3" json:"score,omitempty"`
	Yours              bool                            `protobuf:"varint,16,opt,name=yours,proto3" json:"yours,omitempty"`
	TopicId            int32                           `protobuf:"varint,17,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	TopicSlug          string                          `protobuf:"bytes,18,opt,name=topic_slug,json=topicSlug,proto3" json:"topic_slug,omitempty"`
	DisplayUsername    string                          `protobuf:"bytes,19,opt,name=display_username,json=displayUsername,proto3" json:"display_username,omitempty"`
	Version            int32                           `protobuf:"varint,20,opt,name=version,proto3" json:"version,omitempty"`
	CanEdit            bool                            `protobuf:"varint,21,opt,name=can_edit,json=canEdit,proto3" json:"can_edit,omitempty"`
	CanDelete          bool                            `protobuf:"varint,22,opt,name=can_delete,json=canDelete,proto3" json:"can_delete,omitempty"`
	CanRecover         bool                            `protobuf:"varint,23,opt,name=can_recover,json=canRecover,proto3" json:"can_recover,omitempty"`
	CanWiki            bool                            `protobuf:"varint,24,opt,name=can_wiki,json=canWiki,proto3" json:"can_wiki,omitempty"`
	UserTitle          string                          `protobuf:"bytes,25,opt,name=user_title,json=userTitle,proto3" json:"user_title,omitempty"`
	ActionsSummary     []*PostsResponse_ActionsSummary `protobuf:"bytes,26,rep,name=actions_summary,json=actionsSummary,proto3" json:"actions_summary,omitempty"`
	Moderator          bool                            `protobuf:"varint,27,opt,name=moderator,proto3" json:"moderator,omitempty"`
	Admin              bool                            `protobuf:"varint,28,opt,name=admin,proto3" json:"admin,omitempty"`
	Staff              bool                            `protobuf:"varint,29,opt,name=staff,proto3" json:"staff,omitempty"`
	UserId             int32                           `protobuf:"varint,30,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DraftSequence      int32                           `protobuf:"varint,31,opt,name=draft_sequence,json=draftSequence,proto3" json:"draft_sequence,omitempty"`
	Hidden             bool                            `protobuf:"varint,32,opt,name=hidden,proto3" json:"hidden,omitempty"`
	HiddenReasonId     int32                           `protobuf:"varint,33,opt,name=hidden_reason_id,json=hiddenReasonId,proto3" json:"hidden_reason_id,omitempty"`
	TrustLevel         int32                           `protobuf:"varint,34,opt,name=trust_level,json=trustLevel,proto3" json:"trust_level,omitempty"`
	DeletedAt          *timestamppb.Timestamp          `protobuf:"bytes,35,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	UserDeleted        bool                            `protobuf:"varint,36,opt,name=user_deleted,json=userDeleted,proto3" json:"user_deleted,omitempty"`
	EditReason         string                          `protobuf:"bytes,37,opt,name=edit_reason,json=editReason,proto3" json:"edit_reason,omitempty"`
	CanViewEditHistory bool                            `protobuf:"varint,38,opt,name=can_view_edit_history,json=canViewEditHistory,proto3" json:"can_view_edit_history,omitempty"`
	Wiki               bool                            `protobuf:"varint,39,opt,name=wiki,proto3" json:"wiki,omitempty"`
}

func (x *PostsResponse) Reset() {
	*x = PostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostsResponse) ProtoMessage() {}

func (x *PostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostsResponse.ProtoReflect.Descriptor instead.
func (*PostsResponse) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{1}
}

func (x *PostsResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PostsResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PostsResponse) GetAvatarTemplate() string {
	if x != nil {
		return x.AvatarTemplate
	}
	return ""
}

func (x *PostsResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PostsResponse) GetCooked() string {
	if x != nil {
		return x.Cooked
	}
	return ""
}

func (x *PostsResponse) GetPostNumber() int32 {
	if x != nil {
		return x.PostNumber
	}
	return 0
}

func (x *PostsResponse) GetPostType() int32 {
	if x != nil {
		return x.PostType
	}
	return 0
}

func (x *PostsResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PostsResponse) GetReplyCount() int32 {
	if x != nil {
		return x.ReplyCount
	}
	return 0
}

func (x *PostsResponse) GetReplyToPostNumber() int32 {
	if x != nil {
		return x.ReplyToPostNumber
	}
	return 0
}

func (x *PostsResponse) GetQuoteCount() int32 {
	if x != nil {
		return x.QuoteCount
	}
	return 0
}

func (x *PostsResponse) GetAvgTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AvgTime
	}
	return nil
}

func (x *PostsResponse) GetIncomingLinkCount() int32 {
	if x != nil {
		return x.IncomingLinkCount
	}
	return 0
}

func (x *PostsResponse) GetReads() int32 {
	if x != nil {
		return x.Reads
	}
	return 0
}

func (x *PostsResponse) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *PostsResponse) GetYours() bool {
	if x != nil {
		return x.Yours
	}
	return false
}

func (x *PostsResponse) GetTopicId() int32 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

func (x *PostsResponse) GetTopicSlug() string {
	if x != nil {
		return x.TopicSlug
	}
	return ""
}

func (x *PostsResponse) GetDisplayUsername() string {
	if x != nil {
		return x.DisplayUsername
	}
	return ""
}

func (x *PostsResponse) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *PostsResponse) GetCanEdit() bool {
	if x != nil {
		return x.CanEdit
	}
	return false
}

func (x *PostsResponse) GetCanDelete() bool {
	if x != nil {
		return x.CanDelete
	}
	return false
}

func (x *PostsResponse) GetCanRecover() bool {
	if x != nil {
		return x.CanRecover
	}
	return false
}

func (x *PostsResponse) GetCanWiki() bool {
	if x != nil {
		return x.CanWiki
	}
	return false
}

func (x *PostsResponse) GetUserTitle() string {
	if x != nil {
		return x.UserTitle
	}
	return ""
}

func (x *PostsResponse) GetActionsSummary() []*PostsResponse_ActionsSummary {
	if x != nil {
		return x.ActionsSummary
	}
	return nil
}

func (x *PostsResponse) GetModerator() bool {
	if x != nil {
		return x.Moderator
	}
	return false
}

func (x *PostsResponse) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *PostsResponse) GetStaff() bool {
	if x != nil {
		return x.Staff
	}
	return false
}

func (x *PostsResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PostsResponse) GetDraftSequence() int32 {
	if x != nil {
		return x.DraftSequence
	}
	return 0
}

func (x *PostsResponse) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *PostsResponse) GetHiddenReasonId() int32 {
	if x != nil {
		return x.HiddenReasonId
	}
	return 0
}

func (x *PostsResponse) GetTrustLevel() int32 {
	if x != nil {
		return x.TrustLevel
	}
	return 0
}

func (x *PostsResponse) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *PostsResponse) GetUserDeleted() bool {
	if x != nil {
		return x.UserDeleted
	}
	return false
}

func (x *PostsResponse) GetEditReason() string {
	if x != nil {
		return x.EditReason
	}
	return ""
}

func (x *PostsResponse) GetCanViewEditHistory() bool {
	if x != nil {
		return x.CanViewEditHistory
	}
	return false
}

func (x *PostsResponse) GetWiki() bool {
	if x != nil {
		return x.Wiki
	}
	return false
}

type PostsResponse_ActionsSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Hidden bool  `protobuf:"varint,2,opt,name=hidden,proto3" json:"hidden,omitempty"`
	CanAct bool  `protobuf:"varint,3,opt,name=can_act,json=canAct,proto3" json:"can_act,omitempty"`
}

func (x *PostsResponse_ActionsSummary) Reset() {
	*x = PostsResponse_ActionsSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostsResponse_ActionsSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostsResponse_ActionsSummary) ProtoMessage() {}

func (x *PostsResponse_ActionsSummary) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostsResponse_ActionsSummary.ProtoReflect.Descriptor instead.
func (*PostsResponse_ActionsSummary) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PostsResponse_ActionsSummary) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PostsResponse_ActionsSummary) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *PostsResponse_ActionsSummary) GetCanAct() bool {
	if x != nil {
		return x.CanAct
	}
	return false
}

var File_posts_proto protoreflect.FileDescriptor

var file_posts_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x72, 0x63, 0x68, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x63, 0x68, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9d, 0x0b, 0x0a, 0x0d, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x50,
	0x6f, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x61, 0x76,
	0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x61, 0x76, 0x67, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11,
	0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x79, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x79, 0x6f,
	0x75, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x29, 0x0a,
	0x10, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x6e, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x61, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x61, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x61, 0x6e, 0x5f, 0x77, 0x69, 0x6b, 0x69, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x63, 0x61, 0x6e, 0x57, 0x69, 0x6b, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x59, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x64,
	0x72, 0x61, 0x66, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x20, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69,
	0x64, 0x64, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x21, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x22, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x23, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x24, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x25, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x31,
	0x0a, 0x15, 0x63, 0x61, 0x6e, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x5f,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x26, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x63,
	0x61, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x45, 0x64, 0x69, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x6b, 0x69, 0x18, 0x27, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x77, 0x69, 0x6b, 0x69, 0x1a, 0x51, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x61, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x63, 0x61, 0x6e, 0x41, 0x63, 0x74, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x6c, 0x69,
	0x6e, 0x64, 0x6d, 0x61, 0x6e, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_posts_proto_rawDescOnce sync.Once
	file_posts_proto_rawDescData = file_posts_proto_rawDesc
)

func file_posts_proto_rawDescGZIP() []byte {
	file_posts_proto_rawDescOnce.Do(func() {
		file_posts_proto_rawDescData = protoimpl.X.CompressGZIP(file_posts_proto_rawDescData)
	})
	return file_posts_proto_rawDescData
}

var file_posts_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_posts_proto_goTypes = []interface{}{
	(*PostsRequest)(nil),                 // 0: discourse.protobuf.PostsRequest
	(*PostsResponse)(nil),                // 1: discourse.protobuf.PostsResponse
	(*PostsResponse_ActionsSummary)(nil), // 2: discourse.protobuf.PostsResponse.ActionsSummary
	(*timestamppb.Timestamp)(nil),        // 3: google.protobuf.Timestamp
}
var file_posts_proto_depIdxs = []int32{
	3, // 0: discourse.protobuf.PostsRequest.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: discourse.protobuf.PostsResponse.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: discourse.protobuf.PostsResponse.avg_time:type_name -> google.protobuf.Timestamp
	2, // 3: discourse.protobuf.PostsResponse.actions_summary:type_name -> discourse.protobuf.PostsResponse.ActionsSummary
	3, // 4: discourse.protobuf.PostsResponse.deleted_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_posts_proto_init() }
func file_posts_proto_init() {
	if File_posts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_posts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostsRequest); i {
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
		file_posts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostsResponse); i {
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
		file_posts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostsResponse_ActionsSummary); i {
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
			RawDescriptor: file_posts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_posts_proto_goTypes,
		DependencyIndexes: file_posts_proto_depIdxs,
		MessageInfos:      file_posts_proto_msgTypes,
	}.Build()
	File_posts_proto = out.File
	file_posts_proto_rawDesc = nil
	file_posts_proto_goTypes = nil
	file_posts_proto_depIdxs = nil
}
