package wdk

// AlibabaPricePromotionActivityDeleteResult 结构体
type AlibabaPricePromotionActivityDeleteResult struct {
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// data
	DataList []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
	// code
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// totalRecord
	TotalRecord int64 `json:"total_record,omitempty" xml:"total_record,omitempty"`
	// msg
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
