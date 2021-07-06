package aesolution

// BaseResult 结构体
type BaseResult struct {
	// error message
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// time stamp
	TimeStamp string `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// error code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data *GlobalAeopTpOrderDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// success
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
