package product

// DisplayQualifications 
type DisplayQualifications struct {

    // 返回消息
    Message   string `json:"message,omitempty"`

    // 是否成功
    Result   bool `json:"result,omitempty"`

    // 填充列表
    Model   string `json:"model,omitempty"`

}
