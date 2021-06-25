package wdk

// AlibabaWdkSkuUpdateApiResult 
type AlibabaWdkSkuUpdateApiResult struct {

    // sku编码
    Model   string `json:"model,omitempty"`

    // sku商品操作成功标志
    Success   bool `json:"success,omitempty"`

    // sku商品操作错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // sku商品操作错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 聚合之后的产品id，商家需要保存该值
    ProductId   string `json:"product_id,omitempty"`

}
