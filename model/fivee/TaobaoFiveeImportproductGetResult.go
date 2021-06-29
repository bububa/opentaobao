package fivee

// TaobaoFiveeImportproductGetResult 
type TaobaoFiveeImportproductGetResult struct {
    // code
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // 返回素材id
    Data   *ImportProduct `json:"data,omitempty" xml:"data,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
