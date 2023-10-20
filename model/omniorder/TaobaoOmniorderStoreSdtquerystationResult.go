package omniorder

// TaobaoomniorderstoresdtquerystationResult 结构体
type TaobaoomniorderstoresdtquerystationResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *SdtQueryPackageResponse `json:"data,omitempty" xml:"data,omitempty"`
}
