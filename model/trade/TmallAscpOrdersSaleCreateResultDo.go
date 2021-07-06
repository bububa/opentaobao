package trade

// TmallAscpOrdersSaleCreateResultDo 结构体
type TmallAscpOrdersSaleCreateResultDo struct {
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// module
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
