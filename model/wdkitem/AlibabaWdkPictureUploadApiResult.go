package wdkitem

// AlibabaWdkPictureUploadApiResult 
type AlibabaWdkPictureUploadApiResult struct {
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // model
    Model   *PictureDO `json:"model,omitempty" xml:"model,omitempty"`
    // 错误code
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 错误原因
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}
