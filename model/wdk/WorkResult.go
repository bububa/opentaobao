package wdk

// WorkResult 
type WorkResult struct {
    // 返回数据
    Data   *PackageResultDTO `json:"data,omitempty" xml:"data,omitempty"`
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 错误码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 响应结果
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
