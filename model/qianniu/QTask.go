package qianniu

// QTask 
/* model for simplify = false
type QTask struct {

    // 任务ID
    
    Id   int64 `json:"id,omitempty"`
    

    // 执行者用户数字ID
    
    ReceiverUid   int64 `json:"receiver_uid,omitempty"`
    

    // 执行者用户昵称
    
    ReceiverNick   string `json:"receiver_nick,omitempty"`
    

    // 任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
    
    Status   int64 `json:"status,omitempty"`
    

    // 子任务状态，由业务方自定义
    
    SubStatus   int64 `json:"sub_status,omitempty"`
    

    // 任务完成标识, 0-一个人完成整个任务, 1-所有人完成整个任务完成，冗余任务元数据字段
    
    FinishStrategy   int64 `json:"finish_strategy,omitempty"`
    

    // 任务完成时间，格式：当前时间毫秒数
    
    GmtFinished   string `json:"gmt_finished,omitempty"`
    

    // 业务类型
    
    BizType   string `json:"biz_type,omitempty"`
    

    // 子业务类型
    
    SubBizType   string `json:"sub_biz_type,omitempty"`
    

    // 业务ID
    
    BizId   string `json:"biz_id,omitempty"`
    

    // 业务参数
    
    BizParam   string `json:"biz_param,omitempty"`
    

    // 业务入口
    
    BizEntry   string `json:"biz_entry,omitempty"`
    

    // 任务标签
    
    Tag   string `json:"tag,omitempty"`
    

    // 任务备注
    
    Memo   string `json:"memo,omitempty"`
    

    // 关联的任务元数据
    
    Meta  *struct {
        QTaskMetadata  *QTaskMetadata `json:"q_task_metadata,omitempty"`
    } `json:"meta,omitempty"`
    

    // newYunpanAttachments
    
    NewYunpanAttachments   string `json:"new_yunpan_attachments,omitempty"`
    

    // 任务创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

}
*/

// QTask 
type QTask struct {

    // 任务ID
    Id   int64 `json:"id,omitempty"`

    // 执行者用户数字ID
    ReceiverUid   int64 `json:"receiver_uid,omitempty"`

    // 执行者用户昵称
    ReceiverNick   string `json:"receiver_nick,omitempty"`

    // 任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
    Status   int64 `json:"status,omitempty"`

    // 子任务状态，由业务方自定义
    SubStatus   int64 `json:"sub_status,omitempty"`

    // 任务完成标识, 0-一个人完成整个任务, 1-所有人完成整个任务完成，冗余任务元数据字段
    FinishStrategy   int64 `json:"finish_strategy,omitempty"`

    // 任务完成时间，格式：当前时间毫秒数
    GmtFinished   string `json:"gmt_finished,omitempty"`

    // 业务类型
    BizType   string `json:"biz_type,omitempty"`

    // 子业务类型
    SubBizType   string `json:"sub_biz_type,omitempty"`

    // 业务ID
    BizId   string `json:"biz_id,omitempty"`

    // 业务参数
    BizParam   string `json:"biz_param,omitempty"`

    // 业务入口
    BizEntry   string `json:"biz_entry,omitempty"`

    // 任务标签
    Tag   string `json:"tag,omitempty"`

    // 任务备注
    Memo   string `json:"memo,omitempty"`

    // 关联的任务元数据
    Meta   *QTaskMetadata `json:"meta,omitempty"`

    // newYunpanAttachments
    NewYunpanAttachments   string `json:"new_yunpan_attachments,omitempty"`

    // 任务创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

}
