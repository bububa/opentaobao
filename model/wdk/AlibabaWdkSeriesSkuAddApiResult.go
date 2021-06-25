package wdk

// AlibabaWdkSeriesSkuAddApiResult 
type AlibabaWdkSeriesSkuAddApiResult struct {

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 系列品添加商品成功结果
    Model   *SkuSeriesEditResult `json:"model,omitempty"`

    // 错误详情
    ErrMsg   string `json:"err_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}