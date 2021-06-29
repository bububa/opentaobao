package wdk

// BmResult 
type BmResult struct {
    // 结果数据
    DataList   []AlibabaWdkBmCouponQueryData `json:"data_list,omitempty" xml:"data_list>alibaba_wdk_bm_coupon_query_data,omitempty"`
    // 错误码，当失败返回错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 描述信息，当成功返回OK，当失败返回错误描述
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 库存数量
    Data   int64 `json:"data,omitempty" xml:"data,omitempty"`
    // 详细结果
    PublishResults   []SkuStockPublishResult `json:"publish_results,omitempty" xml:"publish_results>sku_stock_publish_result,omitempty"`
    // 额外属性
    ExtData   string `json:"ext_data,omitempty" xml:"ext_data,omitempty"`
}
