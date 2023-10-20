package qianniu

import (
	"sync"
)

// QTaskMetadata 结构体
type QTaskMetadata struct {
	// 任务标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 任务摘要内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 任务开始时间，格式：当前时间毫秒数
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 任务结束时间，格式：当前时间毫秒数
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 发起者用户昵称
	SenderNick string `json:"sender_nick,omitempty" xml:"sender_nick,omitempty"`
	// 任务元备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 轻任务附件信息，userId_timestamp_随机字符串，例如：23434_1234458623_seresfto
	Attachments string `json:"attachments,omitempty" xml:"attachments,omitempty"`
	// 我的安排的任务上的提醒时间
	BizRemindTime string `json:"biz_remind_time,omitempty" xml:"biz_remind_time,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 任务更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 接收的任务类型，所有相关的任务的类型一致时有效。当任务类型不一致时为diff
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// biztype的中文名
	BizTypeStr string `json:"biz_type_str,omitempty" xml:"biz_type_str,omitempty"`
	// 点击动作的协议
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 语音备注的文件名
	VoiceFile string `json:"voice_file,omitempty" xml:"voice_file,omitempty"`
	// 任务接收人的列表
	Receiver string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// newYunpanAttachments
	NewYunpanAttachments string `json:"new_yunpan_attachments,omitempty" xml:"new_yunpan_attachments,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 来源系统ID
	BizSysId int64 `json:"biz_sys_id,omitempty" xml:"biz_sys_id,omitempty"`
	// 任务在来源系统的类型, 这个是业务系统的自定义类型
	BizSysTaskType int64 `json:"biz_sys_task_type,omitempty" xml:"biz_sys_task_type,omitempty"`
	// 发起者用户数字ID
	SenderUid int64 `json:"sender_uid,omitempty" xml:"sender_uid,omitempty"`
	// 提醒标志位： 0表示不需要提醒，1表示需要状态变更提醒
	ReminderFlag int64 `json:"reminder_flag,omitempty" xml:"reminder_flag,omitempty"`
	// 任务完成标识： 0表示只要有一个人完成任务即可，1表示所有人需要各自都完成任务
	FinishStrategy int64 `json:"finish_strategy,omitempty" xml:"finish_strategy,omitempty"`
	// 优先级，0低，1中，2高
	Priority int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 与此任务元相关的任务数
	TaskCount int64 `json:"task_count,omitempty" xml:"task_count,omitempty"`
	// 是biz_remind_time的数字格式
	BizRemindTimeLong int64 `json:"biz_remind_time_long,omitempty" xml:"biz_remind_time_long,omitempty"`
	// 完成的任务数。如果完成策略是只需要1个人完成的，则只要一个人完成就会返回总任务条数。如果是所有人都要完成的，则会返回实际完成条数。
	FinishCount int64 `json:"finish_count,omitempty" xml:"finish_count,omitempty"`
	// start_time的数字格式
	StartTimeLong int64 `json:"start_time_long,omitempty" xml:"start_time_long,omitempty"`
	// end_time的数字格式
	EndTimeLong int64 `json:"end_time_long,omitempty" xml:"end_time_long,omitempty"`
	// gmt_create的数字格式
	GmtCreateLong int64 `json:"gmt_create_long,omitempty" xml:"gmt_create_long,omitempty"`
	// gmt_modified的数字格式
	GmtModifiedLong int64 `json:"gmt_modified_long,omitempty" xml:"gmt_modified_long,omitempty"`
	// 同次操作的任务元数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 0为未完成，2为完成，4为取消
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 当前任务的评论数
	CommentCount int64 `json:"comment_count,omitempty" xml:"comment_count,omitempty"`
}

var poolQTaskMetadata = sync.Pool{
	New: func() any {
		return new(QTaskMetadata)
	},
}

// GetQTaskMetadata() 从对象池中获取QTaskMetadata
func GetQTaskMetadata() *QTaskMetadata {
	return poolQTaskMetadata.Get().(*QTaskMetadata)
}

// ReleaseQTaskMetadata 释放QTaskMetadata
func ReleaseQTaskMetadata(v *QTaskMetadata) {
	v.Title = ""
	v.Content = ""
	v.StartTime = ""
	v.EndTime = ""
	v.SenderNick = ""
	v.Memo = ""
	v.Attachments = ""
	v.BizRemindTime = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.BizType = ""
	v.BizTypeStr = ""
	v.Action = ""
	v.VoiceFile = ""
	v.Receiver = ""
	v.NewYunpanAttachments = ""
	v.Id = 0
	v.BizSysId = 0
	v.BizSysTaskType = 0
	v.SenderUid = 0
	v.ReminderFlag = 0
	v.FinishStrategy = 0
	v.Priority = 0
	v.TaskCount = 0
	v.BizRemindTimeLong = 0
	v.FinishCount = 0
	v.StartTimeLong = 0
	v.EndTimeLong = 0
	v.GmtCreateLong = 0
	v.GmtModifiedLong = 0
	v.TotalCount = 0
	v.Status = 0
	v.CommentCount = 0
	poolQTaskMetadata.Put(v)
}
