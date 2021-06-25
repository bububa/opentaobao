package blackvip

// ResultSupport 
type ResultSupport struct {

    // 用户数据
    Models   *Models `json:"models,omitempty"`

    // 结果是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

}
