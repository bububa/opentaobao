package category

// ItemPickPagingResult 
type ItemPickPagingResult struct {

    // result
    Results   []CategoryDto `json:"results,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
