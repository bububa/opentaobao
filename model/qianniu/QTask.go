package qianniu

// QTask 结构体
type QTask struct {
	// 执行者用户昵称
	ReceiverNick string `json:"receiver_nick,omitempty" xml:"receiver_nick,omitempty"`
	// 发起人nick
	SenderNick string `json:"sender_nick,omitempty" xml:"sender_nick,omitempty"`
	// 任务创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 任务更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 任务完成时间，格式：当前时间毫秒数
	GmtFinished string `json:"gmt_finished,omitempty" xml:"gmt_finished,omitempty"`
	// 业务类型
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 子业务类型
	SubBizType string `json:"sub_biz_type,omitempty" xml:"sub_biz_type,omitempty"`
	// 业务ID
	BizId string `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
	// 业务参数
	BizParam string `json:"biz_param,omitempty" xml:"biz_param,omitempty"`
	// 业务入口
	BizEntry string `json:"biz_entry,omitempty" xml:"biz_entry,omitempty"`
	// 任务标签
	Tag string `json:"tag,omitempty" xml:"tag,omitempty"`
	// 任务备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 提醒时间
	RemindTime string `json:"remind_time,omitempty" xml:"remind_time,omitempty"`
	// 与业务相关的对象。当前主要用于保存买家nick
	BizNick string `json:"biz_nick,omitempty" xml:"biz_nick,omitempty"`
	// biz_type的类型中文名
	BizTypeStr string `json:"biz_type_str,omitempty" xml:"biz_type_str,omitempty"`
	// json格式的字符串，包含跳转协议
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 点击业务号时的动作
	BizIdAction string `json:"biz_id_action,omitempty" xml:"biz_id_action,omitempty"`
	// 业务号的展示名称
	BizIdName string `json:"biz_id_name,omitempty" xml:"biz_id_name,omitempty"`
	// 新任务的内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 附件的文件名
	Attachments string `json:"attachments,omitempty" xml:"attachments,omitempty"`
	// 语音附件的文件名
	VoiceFile string `json:"voice_file,omitempty" xml:"voice_file,omitempty"`
	// newYunpanAttachments
	NewYunpanAttachments string `json:"new_yunpan_attachments,omitempty" xml:"new_yunpan_attachments,omitempty"`
	// 买家uid密文
	OpenBuyerUid string `json:"open_buyer_uid,omitempty" xml:"open_buyer_uid,omitempty"`
	// 任务ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 执行者用户数字ID
	ReceiverUid int64 `json:"receiver_uid,omitempty" xml:"receiver_uid,omitempty"`
	// 发起人uid
	SenderUid int64 `json:"sender_uid,omitempty" xml:"sender_uid,omitempty"`
	// 表示是否为本条记录接收人实际完成
	FinishFlag int64 `json:"finish_flag,omitempty" xml:"finish_flag,omitempty"`
	// 任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，6-已转发
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 子任务状态，由业务方自定义
	SubStatus int64 `json:"sub_status,omitempty" xml:"sub_status,omitempty"`
	// 任务完成标识, 0-一个人完成整个任务, 1-所有人完成整个任务完成，冗余任务元数据字段
	FinishStrategy int64 `json:"finish_strategy,omitempty" xml:"finish_strategy,omitempty"`
	// gmt_create的数字格式
	GmtCreateLong int64 `json:"gmt_create_long,omitempty" xml:"gmt_create_long,omitempty"`
	// gmt_modified的数字格式
	GmtModifiedLong int64 `json:"gmt_modified_long,omitempty" xml:"gmt_modified_long,omitempty"`
	// gmt_finished的数字格式
	GmtFinishedLong int64 `json:"gmt_finished_long,omitempty" xml:"gmt_finished_long,omitempty"`
	// 到期提醒的方式。0-不提醒 1-pc和移动提醒 2-仅pc提醒 3-仅移动提醒。在查询类接口中，只需要传2或3即可，为1的数据都会包含在其中。
	RemindFlag int64 `json:"remind_flag,omitempty" xml:"remind_flag,omitempty"`
	// remind_time的数字格式
	RemindTimeLong int64 `json:"remind_time_long,omitempty" xml:"remind_time_long,omitempty"`
	// 同次操作相关的任务数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 优先级，同任务元中的priority值，方便查询使用。
	Priority int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 任务元id
	MetadataId int64 `json:"metadata_id,omitempty" xml:"metadata_id,omitempty"`
	// 当前任务的评论数
	CommentCount int64 `json:"comment_count,omitempty" xml:"comment_count,omitempty"`
	// 任务&ldquo;已读&rdquo;、&ldquo;未读&rdquo;状态，0：已读，1：未读
	ReadStatus int64 `json:"read_status,omitempty" xml:"read_status,omitempty"`
	// 关联的任务元数据
	Meta *QTaskMetadata `json:"meta,omitempty" xml:"meta,omitempty"`
	// 是否删除
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 父任务的id
	ParentTaskId int64 `json:"parent_task_id,omitempty" xml:"parent_task_id,omitempty"`
}
