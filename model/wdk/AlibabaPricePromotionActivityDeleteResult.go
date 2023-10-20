package wdk

// AlibabapricepromotionactivitydeleteResult 结构体
type AlibabapricepromotionactivitydeleteResult struct {
	// data
	DataList []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
	// msg
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// totalRecord
	TotalRecord int64 `json:"total_record,omitempty" xml:"total_record,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
