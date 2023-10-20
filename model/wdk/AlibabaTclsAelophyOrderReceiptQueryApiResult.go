package wdk

// AlibabatclsaelophyorderreceiptqueryApiResult 结构体
type AlibabatclsaelophyorderreceiptqueryApiResult struct {
	// 打印商家/顾客联小票数据列表
	OrderList []ReceiptDto `json:"order_list,omitempty" xml:"order_list>receipt_dto,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口状态
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}
