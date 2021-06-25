package user

// OpenaccountLong 
type OpenaccountLong struct {

    // 返回信息
    Message   string `json:"message,omitempty"`

    // 返回数据
    Data   int64 `json:"data,omitempty"`

    // 返回码
    Code   int64 `json:"code,omitempty"`

    // 是否成功
    Successful   bool `json:"successful,omitempty"`

}
