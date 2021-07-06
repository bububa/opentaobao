package xhotel

// TaobaoXhotelDataServiceSellerServiceindexResultSet 结构体
type TaobaoXhotelDataServiceSellerServiceindexResultSet struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// firstResult
	FirstResult *TopAdsSlrQueryResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
}
