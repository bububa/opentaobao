package wdkitem

// AlibabawdkpictureuploadApiResult 结构体
type AlibabawdkpictureuploadApiResult struct {
	// 错误code
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误原因
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// model
	Model *PictureDo `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
