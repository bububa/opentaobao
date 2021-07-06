package wdk

// FulfillLogisticListResult 结构体
type FulfillLogisticListResult struct {
	// 小票批次信息
	Results []ReceiptBatchInfo `json:"results,omitempty" xml:"results>receipt_batch_info,omitempty"`
	// SYSTEM_ERROR("SYSTEM_ERROR", "系统异常"),PARAM_ERROR("PARAM_ERROR", "参数错误"),BUSINESS_ERROR("BUSINESS_ERROR", "业务异常");
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回码含义描述
	ErrDesc string `json:"err_desc,omitempty" xml:"err_desc,omitempty"`
	// true 调用成功 false 调用异常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
