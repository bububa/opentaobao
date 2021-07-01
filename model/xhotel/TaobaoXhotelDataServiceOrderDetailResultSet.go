package xhotel

// TaobaoXhotelDataServiceOrderDetailResultSet 结构体
type TaobaoXhotelDataServiceOrderDetailResultSet struct {
	// firstResult
	FirstResult *TopAdsTripSvcQueryResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
