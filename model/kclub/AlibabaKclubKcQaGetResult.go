package kclub

// AlibabaKclubKcQaGetResult 
type AlibabaKclubKcQaGetResult struct {
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 返回数据
    Data   *KcQaRead `json:"data,omitempty" xml:"data,omitempty"`
    // 错误编码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
