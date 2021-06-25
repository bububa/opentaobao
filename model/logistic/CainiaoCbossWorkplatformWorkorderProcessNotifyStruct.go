package logistic

// CainiaoCbossWorkplatformWorkorderProcessNotifyStruct 
type CainiaoCbossWorkplatformWorkorderProcessNotifyStruct struct {

    // 工单id
    WorkOrderId   string `json:"work_order_id,omitempty"`

    // 工单业务状态
    BizStatus   int64 `json:"biz_status,omitempty"`

    // 进度的操作时间
    ActionTime   string `json:"action_time,omitempty"`

    // 下发的进度类型
    OpType   int64 `json:"op_type,omitempty"`

    // 任务id
    TaskId   string `json:"task_id,omitempty"`

    // 任务名称
    TaskName   string `json:"task_name,omitempty"`

    // 操作者淘宝/旺旺 姓名/昵称
    DealerName   string `json:"dealer_name,omitempty"`

    // 操作者淘宝/旺旺角色
    DealerRole   int64 `json:"dealer_role,omitempty"`

    // 任务的操作备注
    Memo   string `json:"memo,omitempty"`

    // 工单进度的凭证图片地址
    AttachPath   string `json:"attach_path,omitempty"`

    // 扩展字段,半角分号分隔
    Features   string `json:"features,omitempty"`

    // 调用结果
    ResSuccess   bool `json:"res_success,omitempty"`

    // 错误码
    ResErrorCode   string `json:"res_error_code,omitempty"`

    // 错误信息
    ResErrorMsg   string `json:"res_error_msg,omitempty"`

}
