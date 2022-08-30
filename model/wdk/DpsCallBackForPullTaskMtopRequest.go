package wdk

// DpsCallBackForPullTaskMtopRequest 结构体
type DpsCallBackForPullTaskMtopRequest struct {
	// 任务列表
	TaskCodeList []string `json:"task_code_list,omitempty" xml:"task_code_list>string,omitempty"`
	// 仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}
