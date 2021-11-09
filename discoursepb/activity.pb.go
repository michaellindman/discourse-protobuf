// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.3
// source: activity.proto

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

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post []*Activity_Post `protobuf:"bytes,1,rep,name=post,proto3" json:"post,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{0}
}

func (x *Activity) GetPost() []*Activity_Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type Activity_Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                          int32                           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                        string                          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Username                    string                          `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	AvatarTemplate              string                          `protobuf:"bytes,4,opt,name=avatar_template,json=avatarTemplate,proto3" json:"avatar_template,omitempty"`
	CreatedAt                   *timestamppb.Timestamp          `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Cooked                      string                          `protobuf:"bytes,6,opt,name=cooked,proto3" json:"cooked,omitempty"`
	PostNumber                  int32                           `protobuf:"varint,7,opt,name=post_number,json=postNumber,proto3" json:"post_number,omitempty"`
	PostType                    int32                           `protobuf:"varint,8,opt,name=post_type,json=postType,proto3" json:"post_type,omitempty"`
	UpdatedAt                   *timestamppb.Timestamp          `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ReplyCount                  int32                           `protobuf:"varint,10,opt,name=reply_count,json=replyCount,proto3" json:"reply_count,omitempty"`
	ReplyToPostNumber           int32                           `protobuf:"varint,11,opt,name=reply_to_post_number,json=replyToPostNumber,proto3" json:"reply_to_post_number,omitempty"`
	QouteCount                  int32                           `protobuf:"varint,12,opt,name=qoute_count,json=qouteCount,proto3" json:"qoute_count,omitempty"`
	IncomingLinkCount           int32                           `protobuf:"varint,13,opt,name=incoming_link_count,json=incomingLinkCount,proto3" json:"incoming_link_count,omitempty"`
	Reads                       int32                           `protobuf:"varint,14,opt,name=reads,proto3" json:"reads,omitempty"`
	ReadersCount                int32                           `protobuf:"varint,15,opt,name=readers_count,json=readersCount,proto3" json:"readers_count,omitempty"`
	Score                       int32                           `protobuf:"varint,16,opt,name=score,proto3" json:"score,omitempty"`
	Yours                       bool                            `protobuf:"varint,17,opt,name=yours,proto3" json:"yours,omitempty"`
	TopicId                     int32                           `protobuf:"varint,18,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	TopicSlug                   string                          `protobuf:"bytes,19,opt,name=topic_slug,json=topicSlug,proto3" json:"topic_slug,omitempty"`
	DisplayUsername             string                          `protobuf:"bytes,20,opt,name=display_username,json=displayUsername,proto3" json:"display_username,omitempty"`
	PrimaryGroupName            string                          `protobuf:"bytes,21,opt,name=primary_group_name,json=primaryGroupName,proto3" json:"primary_group_name,omitempty"`
	PrimaryGroupFlairUrl        string                          `protobuf:"bytes,22,opt,name=primary_group_flair_url,json=primaryGroupFlairUrl,proto3" json:"primary_group_flair_url,omitempty"`
	PrimaryGroupFlairBgColor    string                          `protobuf:"bytes,23,opt,name=primary_group_flair_bg_color,json=primaryGroupFlairBgColor,proto3" json:"primary_group_flair_bg_color,omitempty"`
	PrimaryGroupFlairColor      string                          `protobuf:"bytes,24,opt,name=primary_group_flair_color,json=primaryGroupFlairColor,proto3" json:"primary_group_flair_color,omitempty"`
	Version                     int32                           `protobuf:"varint,25,opt,name=version,proto3" json:"version,omitempty"`
	CanEdit                     bool                            `protobuf:"varint,26,opt,name=can_edit,json=canEdit,proto3" json:"can_edit,omitempty"`
	CanDelete                   bool                            `protobuf:"varint,27,opt,name=can_delete,json=canDelete,proto3" json:"can_delete,omitempty"`
	CanRecover                  bool                            `protobuf:"varint,28,opt,name=can_recover,json=canRecover,proto3" json:"can_recover,omitempty"`
	CanWiki                     bool                            `protobuf:"varint,29,opt,name=can_wiki,json=canWiki,proto3" json:"can_wiki,omitempty"`
	UserTitle                   string                          `protobuf:"bytes,30,opt,name=user_title,json=userTitle,proto3" json:"user_title,omitempty"`
	TitleIsGroup                bool                            `protobuf:"varint,31,opt,name=title_is_group,json=titleIsGroup,proto3" json:"title_is_group,omitempty"`
	Bookmarked                  bool                            `protobuf:"varint,32,opt,name=bookmarked,proto3" json:"bookmarked,omitempty"`
	ActionsSummary              []*Activity_Post_ActionsSummary `protobuf:"bytes,33,rep,name=actions_summary,json=actionsSummary,proto3" json:"actions_summary,omitempty"`
	Moderator                   bool                            `protobuf:"varint,34,opt,name=moderator,proto3" json:"moderator,omitempty"`
	Admin                       bool                            `protobuf:"varint,35,opt,name=admin,proto3" json:"admin,omitempty"`
	Staff                       bool                            `protobuf:"varint,36,opt,name=staff,proto3" json:"staff,omitempty"`
	UserId                      int32                           `protobuf:"varint,37,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Hidden                      bool                            `protobuf:"varint,38,opt,name=hidden,proto3" json:"hidden,omitempty"`
	TrustLevel                  int32                           `protobuf:"varint,39,opt,name=trust_level,json=trustLevel,proto3" json:"trust_level,omitempty"`
	DeletedAt                   *timestamppb.Timestamp          `protobuf:"bytes,40,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	UserDeleted                 bool                            `protobuf:"varint,41,opt,name=user_deleted,json=userDeleted,proto3" json:"user_deleted,omitempty"`
	EditReason                  string                          `protobuf:"bytes,42,opt,name=edit_reason,json=editReason,proto3" json:"edit_reason,omitempty"`
	CanViewEditHistory          bool                            `protobuf:"varint,43,opt,name=can_view_edit_history,json=canViewEditHistory,proto3" json:"can_view_edit_history,omitempty"`
	Wiki                        bool                            `protobuf:"varint,44,opt,name=wiki,proto3" json:"wiki,omitempty"`
	Excerpt                     string                          `protobuf:"bytes,45,opt,name=excerpt,proto3" json:"excerpt,omitempty"`
	ReviewableId                int32                           `protobuf:"varint,46,opt,name=reviewable_id,json=reviewableId,proto3" json:"reviewable_id,omitempty"`
	ReviewableScoreCount        int32                           `protobuf:"varint,47,opt,name=reviewable_score_count,json=reviewableScoreCount,proto3" json:"reviewable_score_count,omitempty"`
	ReviewableScorePendingCount int32                           `protobuf:"varint,48,opt,name=reviewable_score_pending_count,json=reviewableScorePendingCount,proto3" json:"reviewable_score_pending_count,omitempty"`
	CanAcceptAnswer             bool                            `protobuf:"varint,49,opt,name=can_accept_answer,json=canAcceptAnswer,proto3" json:"can_accept_answer,omitempty"`
	CanUnacceptAnswer           bool                            `protobuf:"varint,50,opt,name=can_unaccept_answer,json=canUnacceptAnswer,proto3" json:"can_unaccept_answer,omitempty"`
	AcceptedAnswer              bool                            `protobuf:"varint,51,opt,name=accepted_answer,json=acceptedAnswer,proto3" json:"accepted_answer,omitempty"`
}

func (x *Activity_Post) Reset() {
	*x = Activity_Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity_Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity_Post) ProtoMessage() {}

func (x *Activity_Post) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity_Post.ProtoReflect.Descriptor instead.
func (*Activity_Post) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Activity_Post) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Activity_Post) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Activity_Post) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Activity_Post) GetAvatarTemplate() string {
	if x != nil {
		return x.AvatarTemplate
	}
	return ""
}

func (x *Activity_Post) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Activity_Post) GetCooked() string {
	if x != nil {
		return x.Cooked
	}
	return ""
}

func (x *Activity_Post) GetPostNumber() int32 {
	if x != nil {
		return x.PostNumber
	}
	return 0
}

func (x *Activity_Post) GetPostType() int32 {
	if x != nil {
		return x.PostType
	}
	return 0
}

func (x *Activity_Post) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Activity_Post) GetReplyCount() int32 {
	if x != nil {
		return x.ReplyCount
	}
	return 0
}

func (x *Activity_Post) GetReplyToPostNumber() int32 {
	if x != nil {
		return x.ReplyToPostNumber
	}
	return 0
}

func (x *Activity_Post) GetQouteCount() int32 {
	if x != nil {
		return x.QouteCount
	}
	return 0
}

func (x *Activity_Post) GetIncomingLinkCount() int32 {
	if x != nil {
		return x.IncomingLinkCount
	}
	return 0
}

func (x *Activity_Post) GetReads() int32 {
	if x != nil {
		return x.Reads
	}
	return 0
}

func (x *Activity_Post) GetReadersCount() int32 {
	if x != nil {
		return x.ReadersCount
	}
	return 0
}

func (x *Activity_Post) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Activity_Post) GetYours() bool {
	if x != nil {
		return x.Yours
	}
	return false
}

func (x *Activity_Post) GetTopicId() int32 {
	if x != nil {
		return x.TopicId
	}
	return 0
}

func (x *Activity_Post) GetTopicSlug() string {
	if x != nil {
		return x.TopicSlug
	}
	return ""
}

func (x *Activity_Post) GetDisplayUsername() string {
	if x != nil {
		return x.DisplayUsername
	}
	return ""
}

func (x *Activity_Post) GetPrimaryGroupName() string {
	if x != nil {
		return x.PrimaryGroupName
	}
	return ""
}

func (x *Activity_Post) GetPrimaryGroupFlairUrl() string {
	if x != nil {
		return x.PrimaryGroupFlairUrl
	}
	return ""
}

func (x *Activity_Post) GetPrimaryGroupFlairBgColor() string {
	if x != nil {
		return x.PrimaryGroupFlairBgColor
	}
	return ""
}

func (x *Activity_Post) GetPrimaryGroupFlairColor() string {
	if x != nil {
		return x.PrimaryGroupFlairColor
	}
	return ""
}

func (x *Activity_Post) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Activity_Post) GetCanEdit() bool {
	if x != nil {
		return x.CanEdit
	}
	return false
}

func (x *Activity_Post) GetCanDelete() bool {
	if x != nil {
		return x.CanDelete
	}
	return false
}

func (x *Activity_Post) GetCanRecover() bool {
	if x != nil {
		return x.CanRecover
	}
	return false
}

func (x *Activity_Post) GetCanWiki() bool {
	if x != nil {
		return x.CanWiki
	}
	return false
}

func (x *Activity_Post) GetUserTitle() string {
	if x != nil {
		return x.UserTitle
	}
	return ""
}

func (x *Activity_Post) GetTitleIsGroup() bool {
	if x != nil {
		return x.TitleIsGroup
	}
	return false
}

func (x *Activity_Post) GetBookmarked() bool {
	if x != nil {
		return x.Bookmarked
	}
	return false
}

func (x *Activity_Post) GetActionsSummary() []*Activity_Post_ActionsSummary {
	if x != nil {
		return x.ActionsSummary
	}
	return nil
}

func (x *Activity_Post) GetModerator() bool {
	if x != nil {
		return x.Moderator
	}
	return false
}

func (x *Activity_Post) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *Activity_Post) GetStaff() bool {
	if x != nil {
		return x.Staff
	}
	return false
}

func (x *Activity_Post) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Activity_Post) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *Activity_Post) GetTrustLevel() int32 {
	if x != nil {
		return x.TrustLevel
	}
	return 0
}

func (x *Activity_Post) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Activity_Post) GetUserDeleted() bool {
	if x != nil {
		return x.UserDeleted
	}
	return false
}

func (x *Activity_Post) GetEditReason() string {
	if x != nil {
		return x.EditReason
	}
	return ""
}

func (x *Activity_Post) GetCanViewEditHistory() bool {
	if x != nil {
		return x.CanViewEditHistory
	}
	return false
}

func (x *Activity_Post) GetWiki() bool {
	if x != nil {
		return x.Wiki
	}
	return false
}

func (x *Activity_Post) GetExcerpt() string {
	if x != nil {
		return x.Excerpt
	}
	return ""
}

func (x *Activity_Post) GetReviewableId() int32 {
	if x != nil {
		return x.ReviewableId
	}
	return 0
}

func (x *Activity_Post) GetReviewableScoreCount() int32 {
	if x != nil {
		return x.ReviewableScoreCount
	}
	return 0
}

func (x *Activity_Post) GetReviewableScorePendingCount() int32 {
	if x != nil {
		return x.ReviewableScorePendingCount
	}
	return 0
}

func (x *Activity_Post) GetCanAcceptAnswer() bool {
	if x != nil {
		return x.CanAcceptAnswer
	}
	return false
}

func (x *Activity_Post) GetCanUnacceptAnswer() bool {
	if x != nil {
		return x.CanUnacceptAnswer
	}
	return false
}

func (x *Activity_Post) GetAcceptedAnswer() bool {
	if x != nil {
		return x.AcceptedAnswer
	}
	return false
}

type Activity_Post_ActionsSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Count  int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	CanAct bool  `protobuf:"varint,3,opt,name=can_act,json=canAct,proto3" json:"can_act,omitempty"`
}

func (x *Activity_Post_ActionsSummary) Reset() {
	*x = Activity_Post_ActionsSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity_Post_ActionsSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity_Post_ActionsSummary) ProtoMessage() {}

func (x *Activity_Post_ActionsSummary) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity_Post_ActionsSummary.ProtoReflect.Descriptor instead.
func (*Activity_Post_ActionsSummary) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Activity_Post_ActionsSummary) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Activity_Post_ActionsSummary) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Activity_Post_ActionsSummary) GetCanAct() bool {
	if x != nil {
		return x.CanAct
	}
	return false
}

var File_activity_proto protoreflect.FileDescriptor

var file_activity_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x10, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x12, 0x35, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x1a, 0xcc, 0x0f, 0x0a, 0x04, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x50,
	0x6f, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x71, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e,
	0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x61, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x79,
	0x6f, 0x75, 0x72, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x79, 0x6f, 0x75, 0x72,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x66, 0x6c, 0x61, 0x69, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x46, 0x6c, 0x61, 0x69, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x3e, 0x0a, 0x1c, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x66, 0x6c, 0x61,
	0x69, 0x72, 0x5f, 0x62, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x18, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46,
	0x6c, 0x61, 0x69, 0x72, 0x42, 0x67, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x19, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x66, 0x6c, 0x61,
	0x69, 0x72, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x6c, 0x61, 0x69,
	0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x19, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x6e, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x18, 0x1a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x63, 0x61, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x61, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x63, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x61, 0x6e, 0x5f, 0x77, 0x69, 0x6b, 0x69, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63,
	0x61, 0x6e, 0x57, 0x69, 0x6b, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x69,
	0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x49, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x62,
	0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x18, 0x20, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x12, 0x59, 0x0a, 0x0f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x21,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x22, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x23, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x18, 0x24, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x25, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64,
	0x64, 0x65, 0x6e, 0x18, 0x26, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x27, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x29, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x31, 0x0a, 0x15, 0x63, 0x61, 0x6e, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x65, 0x64,
	0x69, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x63, 0x61, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x45, 0x64, 0x69, 0x74, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x6b, 0x69, 0x18, 0x2c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x77, 0x69, 0x6b, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x63, 0x65,
	0x72, 0x70, 0x74, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x63, 0x65, 0x72,
	0x70, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x2f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x43, 0x0a,
	0x1e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x30, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1b, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x61, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x31, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63,
	0x61, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x13, 0x63, 0x61, 0x6e, 0x5f, 0x75, 0x6e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x32, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x61, 0x6e,
	0x55, 0x6e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x18, 0x33, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65,
	0x64, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x1a, 0x4f, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x61, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x63, 0x61, 0x6e, 0x41, 0x63, 0x74, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x6c, 0x69,
	0x6e, 0x64, 0x6d, 0x61, 0x6e, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activity_proto_rawDescOnce sync.Once
	file_activity_proto_rawDescData = file_activity_proto_rawDesc
)

func file_activity_proto_rawDescGZIP() []byte {
	file_activity_proto_rawDescOnce.Do(func() {
		file_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_activity_proto_rawDescData)
	})
	return file_activity_proto_rawDescData
}

var file_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_activity_proto_goTypes = []interface{}{
	(*Activity)(nil),                     // 0: discourse.protobuf.Activity
	(*Activity_Post)(nil),                // 1: discourse.protobuf.Activity.Post
	(*Activity_Post_ActionsSummary)(nil), // 2: discourse.protobuf.Activity.Post.ActionsSummary
	(*timestamppb.Timestamp)(nil),        // 3: google.protobuf.Timestamp
}
var file_activity_proto_depIdxs = []int32{
	1, // 0: discourse.protobuf.Activity.post:type_name -> discourse.protobuf.Activity.Post
	3, // 1: discourse.protobuf.Activity.Post.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: discourse.protobuf.Activity.Post.updated_at:type_name -> google.protobuf.Timestamp
	2, // 3: discourse.protobuf.Activity.Post.actions_summary:type_name -> discourse.protobuf.Activity.Post.ActionsSummary
	3, // 4: discourse.protobuf.Activity.Post.deleted_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_activity_proto_init() }
func file_activity_proto_init() {
	if File_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_activity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_activity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity_Post); i {
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
		file_activity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity_Post_ActionsSummary); i {
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
			RawDescriptor: file_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_activity_proto_goTypes,
		DependencyIndexes: file_activity_proto_depIdxs,
		MessageInfos:      file_activity_proto_msgTypes,
	}.Build()
	File_activity_proto = out.File
	file_activity_proto_rawDesc = nil
	file_activity_proto_goTypes = nil
	file_activity_proto_depIdxs = nil
}
