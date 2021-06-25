package wdk

// AlibabaWdkItemChangepriceQueryResult 
type AlibabaWdkItemChangepriceQueryResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 结果码
    ReturnCode   string `json:"return_code,omitempty"`

    // 结果文案
    ReturnMsg   string `json:"return_msg,omitempty"`

    // 返回的商品改价单
    Models   []String `json:"models,omitempty"`

}
