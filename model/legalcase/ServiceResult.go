package legalcase

// ServiceResult 
type ServiceResult struct {

    // success
    Success   bool `json:"success,omitempty"`

    // errorcode
    ErrorCode   string `json:"error_code,omitempty"`

    // 内容
    Contents   []Content `json:"contents,omitempty"`

    // errormasg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // content
    Content   string `json:"content,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误编码
    Errcode   string `json:"errcode,omitempty"`

    // 错误信息
    Errmsg   string `json:"errmsg,omitempty"`

}
