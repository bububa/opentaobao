package wdk

// WorkResult 
type WorkResult struct {

    // 返回数据
    Data   *PackageResultDto `json:"data,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 错误码
    Code   string `json:"code,omitempty"`

    // 响应结果
    Success   bool `json:"success,omitempty"`

}
