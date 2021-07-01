package tmallnr

// AlibabaLsyCrmActivityPageUpdateResultDo 结构体
type AlibabaLsyCrmActivityPageUpdateResultDo struct {
	// 活动ID
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}
