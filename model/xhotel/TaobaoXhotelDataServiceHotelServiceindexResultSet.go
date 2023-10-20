package xhotel

// TaobaoxhoteldataservicehotelserviceindexResultSet 结构体
type TaobaoxhoteldataservicehotelserviceindexResultSet struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// firstResult
	FirstResult *TopAdsHtlDataQueryResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
