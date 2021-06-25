package aliqin

// AlibabaAliqinFcIotCardofferResult 
type AlibabaAliqinFcIotCardofferResult struct {

    // 结果对象
    Models   []AlibabaAliqinFcIotCardofferModel `json:"models,omitempty"`

    // 1.成功；2.失败
    Code   string `json:"code,omitempty"`

    // 状态
    Success   bool `json:"success,omitempty"`

    // 错误信息
    Msg   string `json:"msg,omitempty"`

}
