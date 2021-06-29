package idleitem

// EasyResultDTO 
type EasyResultDTO struct {
    // 错误编码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 上传成功的文件id
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 成功与否
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
