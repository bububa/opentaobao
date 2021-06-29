package wdk

// AlibabaWdkSkuCombineskuUpdateApiResult 
type AlibabaWdkSkuCombineskuUpdateApiResult struct {
    // 单个商品是否更新成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 单个商品更新异常编码（异常才有值）
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 单个商品更新异常信息（异常才有值）
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 商品编码值
    Model   string `json:"model,omitempty" xml:"model,omitempty"`
}
