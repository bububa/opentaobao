package alink

// TopServiceResult 
type TopServiceResult struct {
    // 接口返回数据
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // message
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    // 详细描述
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
}
