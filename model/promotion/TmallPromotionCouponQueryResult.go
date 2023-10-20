package promotion

// TmallpromotioncouponqueryResult 结构体
type TmallpromotioncouponqueryResult struct {
	// data
	DataList []TmallpromotioncouponqueryData `json:"data_list,omitempty" xml:"data_list>tmallpromotioncouponquery_data,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// resultCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
