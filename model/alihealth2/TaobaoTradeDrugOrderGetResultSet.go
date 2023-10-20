package alihealth2

// TaobaoTradeDrugOrderGetResultSet 结构体
type TaobaoTradeDrugOrderGetResultSet struct {
	// 异常信息
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// TakeoutThirdOrder订单数据
	Result *TakeoutThirdOrder `json:"result,omitempty" xml:"result,omitempty"`
}
