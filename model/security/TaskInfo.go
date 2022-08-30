package security

// TaskInfo 结构体
type TaskInfo struct {
	// 任务唯一标识
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 任务处理进度: 1-处理完成 (可立即通过对应的查询接口查询处理结果), 2-异步处理中(需等待app_info.callback_url接收到反向通知后再查询结果) ; 对于app_info.data_type为1目前都是异步处理，此字段返回2; 对于app_info.data_type为2目前都是同步处理，此字段返回1
	Progress int64 `json:"progress,omitempty" xml:"progress,omitempty"`
}
