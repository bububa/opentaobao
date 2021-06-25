package qianniu

// QTaskMetadata 
type QTaskMetadata struct {

    // 主键
    Id   int64 `json:"id,omitempty"`

    // 任务标题
    Title   string `json:"title,omitempty"`

    // 任务摘要内容
    Content   string `json:"content,omitempty"`

    // 来源系统ID
    BizSysId   int64 `json:"biz_sys_id,omitempty"`

    // 任务在来源系统的类型, 这个是业务系统的自定义类型
    BizSysTaskType   int64 `json:"biz_sys_task_type,omitempty"`

    // 任务开始时间，格式：当前时间毫秒数
    StartTime   string `json:"start_time,omitempty"`

    // 任务结束时间，格式：当前时间毫秒数
    EndTime   string `json:"end_time,omitempty"`

    // 发起者用户数字ID
    SenderUid   int64 `json:"sender_uid,omitempty"`

    // 发起者用户昵称
    SenderNick   string `json:"sender_nick,omitempty"`

    // 提醒标志位： 0表示不需要提醒，1表示需要状态变更提醒
    ReminderFlag   int64 `json:"reminder_flag,omitempty"`

    // 任务完成标识： 0表示只要有一个人完成任务即可，1表示所有人需要各自都完成任务
    FinishStrategy   int64 `json:"finish_strategy,omitempty"`

    // 与此任务元相关的任务数
    TaskCount   int64 `json:"task_count,omitempty"`

    // 优先级，0低，1中，2高
    Priority   int64 `json:"priority,omitempty"`

    // 任务元备注
    Memo   string `json:"memo,omitempty"`

    // 轻任务附件信息，userId_timestamp_随机字符串，例如：23434_1234458623_seresfto
    Attachments   string `json:"attachments,omitempty"`

    // 语音备注的文件名
    VoiceFile   string `json:"voice_file,omitempty"`

    // 任务的接收人
    Receiver   string `json:"receiver,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 完成的任务数。如果完成策略是只需要1个人完成的，则只要一个人完成就会返回总任务条数。如果是所有人都要完成的，则会返回实际完成条数。
    FinishCount   int64 `json:"finish_count,omitempty"`

    // 任务更新时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // start_time的数字格式
    StartTimeLong   int64 `json:"start_time_long,omitempty"`

    // end_time的数字格式
    EndTimeLong   int64 `json:"end_time_long,omitempty"`

    // gmt_create的数字格式
    GmtCreateLong   int64 `json:"gmt_create_long,omitempty"`

    // gmt_modified的数字格式
    GmtModifiedLong   int64 `json:"gmt_modified_long,omitempty"`

    // 同次操作的任务元数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 接收的任务类型，所有相关的任务的类型一致时有效。当任务类型不一致时为diff
    BizType   string `json:"biz_type,omitempty"`

    // 0为未完成，2为完成，4为取消
    Status   int64 `json:"status,omitempty"`

    // biztype的中文名
    BizTypeStr   string `json:"biz_type_str,omitempty"`

    // 当前任务的评论数
    CommentCount   int64 `json:"comment_count,omitempty"`

    // 点击动作的协议
    Action   string `json:"action,omitempty"`

    // 我的安排的任务上的提醒时间
    BizRemindTime   string `json:"biz_remind_time,omitempty"`

    // 是biz_remind_time的数字格式
    BizRemindTimeLong   int64 `json:"biz_remind_time_long,omitempty"`

    // newYunpanAttachments
    NewYunpanAttachments   string `json:"new_yunpan_attachments,omitempty"`

}
