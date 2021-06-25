package wdk

// AlibabaWdkItemCurrentpriceQueryResult 
type AlibabaWdkItemCurrentpriceQueryResult struct {

    // 返回码
    ReturnCode   string `json:"return_code,omitempty"`

    // 异常信息
    ReturnMsg   string `json:"return_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 返回的当前当前商品价格列表
    Models   []AlibabaWdkItemCurrentpriceQueryModel `json:"models,omitempty"`

}
