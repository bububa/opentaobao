package wdk

// AlibabaWdkItemChangepriceQueryResult 
type AlibabaWdkItemChangepriceQueryResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 结果码
    ReturnCode   string `json:"return_code,omitempty" xml:"return_code,omitempty"`
    // 结果文案
    ReturnMsg   string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
    // 返回的商品改价单
    Models   []string `json:"models,omitempty" xml:"models>string,omitempty"`
}
