package media

// ResultCode 
type ResultCode struct {

    // 错误代码
    ErrorCode   int64 `json:"error_code,omitempty"`

    // 错误描述
    Info   string `json:"info,omitempty"`

}
