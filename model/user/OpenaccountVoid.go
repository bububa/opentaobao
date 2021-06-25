package user

// OpenaccountVoid 
type OpenaccountVoid struct {

    // 返回信息
    Message   string `json:"message,omitempty"`

    // 是否成功
    Successful   bool `json:"successful,omitempty"`

    // 错误码
    Code   int64 `json:"code,omitempty"`

}
