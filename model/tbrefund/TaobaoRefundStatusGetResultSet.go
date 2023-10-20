package tbrefund

// TaobaorefundstatusgetResultSet 结构体
type TaobaorefundstatusgetResultSet struct {
	// 数组对象
	ResultList []QueryRefundStatusResponse `json:"result_list,omitempty" xml:"result_list>query_refund_status_response,omitempty"`
	// 错误码，没有表示无异常
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
