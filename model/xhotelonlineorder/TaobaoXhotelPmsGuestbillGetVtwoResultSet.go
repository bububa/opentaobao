package xhotelonlineorder

// TaobaoXhotelPmsGuestbillGetVtwoResultSet 结构体
type TaobaoXhotelPmsGuestbillGetVtwoResultSet struct {
	// 账单列表中涉及到的金额费用单位均为分
	Results []OrderBillInfo `json:"results,omitempty" xml:"results>order_bill_info,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
