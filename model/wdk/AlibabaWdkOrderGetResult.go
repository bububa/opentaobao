package wdk

// AlibabaWdkOrderGetResult 结构体
type AlibabaWdkOrderGetResult struct {
	// 订单列表
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// returnCode
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// returnMsg
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 返回本查询条件下的数据总数
	TotalNumber int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}
