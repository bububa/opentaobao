package logistic

// LogisticsResult 
type LogisticsResult struct {

    // 错误编码
    Code   string `json:"code,omitempty"`

    // 数据
    Data   *Pagination `json:"data,omitempty"`

    // 失败消息
    Message   string `json:"message,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
