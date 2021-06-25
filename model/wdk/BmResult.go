package wdk

// BmResult 
type BmResult struct {

    // 结果数据
    DataList   []AlibabaWdkBmCouponQueryData `json:"data_list,omitempty"`

    // 错误码，当失败返回错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 描述信息，当成功返回OK，当失败返回错误描述
    Message   string `json:"message,omitempty"`

    // 库存数量
    Data   int64 `json:"data,omitempty"`

    // 详细结果
    PublishResults   []SkuStockPublishResult `json:"publish_results,omitempty"`

    // 额外属性
    ExtData   string `json:"ext_data,omitempty"`

}
