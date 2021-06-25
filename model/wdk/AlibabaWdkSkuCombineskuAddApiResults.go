package wdk

// AlibabaWdkSkuCombineskuAddApiResults 
type AlibabaWdkSkuCombineskuAddApiResults struct {

    // 接口调用是否成功
    Success   bool `json:"success,omitempty"`

    // 接口调用异常编码
    ErrCode   string `json:"err_code,omitempty"`

    // 接口调用异常信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 商品列表
    Models   []AlibabaWdkSkuCombineskuAddApiResult `json:"models,omitempty"`

}
