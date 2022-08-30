package wdk

// DpsCallBackSortMtopRequest 结构体
type DpsCallBackSortMtopRequest struct {
	// 明细列表
	DetailRequestList []DpsCallBackSortDetailMtopRequest `json:"detail_request_list,omitempty" xml:"detail_request_list>dps_call_back_sort_detail_mtop_request,omitempty"`
	// 任务号
	TaskCode string `json:"task_code,omitempty" xml:"task_code,omitempty"`
	// 账号
	UserAccount string `json:"user_account,omitempty" xml:"user_account,omitempty"`
	// 仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}
