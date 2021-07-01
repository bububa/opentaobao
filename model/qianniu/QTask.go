package qianniu

// QTask 结构体
type QTask struct {
	// 任务ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 执行者用户数字ID
	ReceiverUid int64 `json:"receiver_uid,omitempty" xml:"receiver_uid,omitempty"`
	// 执行者用户昵称
	ReceiverNick string `json:"receiver_nick,omitempty" xml:"receiver_nick,omitempty"`
	// 任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 子任务状态，由业务方自定义
	SubStatus int64 `json:"sub_status,omitempty" xml:"sub_status,omitempty"`
	// 任务完成标识, 0-一个人完成整个任务, 1-所有人完成整个任务完成，冗余任务元数据字段
	FinishStrategy int64 `json:"finish_strategy,omitempty" xml:"finish_strategy,omitempty"`
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
	// 关联的任务元数据
	Meta *QTaskMetadata `json:"meta,omitempty" xml:"meta,omitempty"`
	// newYunpanAttachments
	NewYunpanAttachments string `json:"new_yunpan_attachments,omitempty" xml:"new_yunpan_attachments,omitempty"`
	// 任务创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
}
