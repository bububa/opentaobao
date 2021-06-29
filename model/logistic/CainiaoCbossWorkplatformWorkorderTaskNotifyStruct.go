package logistic

// CainiaoCbossWorkplatformWorkorderTaskNotifyStruct 
type CainiaoCbossWorkplatformWorkorderTaskNotifyStruct struct {
    // 扩展字段
    Features   string `json:"features,omitempty" xml:"features,omitempty"`
    // 任务超时时间
    Timeout   string `json:"timeout,omitempty" xml:"timeout,omitempty"`
    // 任务下发次数
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    // 任务创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 工单任务状态
    TaskStatus   int64 `json:"task_status,omitempty" xml:"task_status,omitempty"`
    // 工单任务名称
    TaskName   string `json:"task_name,omitempty" xml:"task_name,omitempty"`
    // 工单任务id
    TaskId   string `json:"task_id,omitempty" xml:"task_id,omitempty"`
    // 工单id
    WorkOrderId   string `json:"work_order_id,omitempty" xml:"work_order_id,omitempty"`
    // 是否成功
    ResSuccess   bool `json:"res_success,omitempty" xml:"res_success,omitempty"`
    // 错误码
    ResErrorCode   string `json:"res_error_code,omitempty" xml:"res_error_code,omitempty"`
    // 错误信息
    ResErrorMsg   string `json:"res_error_msg,omitempty" xml:"res_error_msg,omitempty"`
}
