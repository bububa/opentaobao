package wdk

// SkuStockPublishParamDo 
/* model for simplify = false
type SkuStockPublishParamDo struct {

    // 商家门店编码
    
    ShopCode   string `json:"shop_code,omitempty"`
    

    // 这笔单据发生的原因
    
    Reason   string `json:"reason,omitempty"`
    

    // 移动的数量，正数表式增加，负数表式减少
    
    Quantity   int64 `json:"quantity,omitempty"`
    

    // 唯一单据号，用于幂等操作
    
    BillNo   string `json:"bill_no,omitempty"`
    

    // 商品对应的69码
    
    Barcode   string `json:"barcode,omitempty"`
    

    // 更新类型，1表式覆盖，0表式增量
    
    UpdateType   int64 `json:"update_type,omitempty"`
    

    // 当时业务发生的时间戳，单位：ms
    
    OperationTs   int64 `json:"operation_ts,omitempty"`
    

}
*/

// SkuStockPublishParamDo 
type SkuStockPublishParamDo struct {

    // 商家门店编码
    ShopCode   string `json:"shop_code,omitempty"`

    // 这笔单据发生的原因
    Reason   string `json:"reason,omitempty"`

    // 移动的数量，正数表式增加，负数表式减少
    Quantity   int64 `json:"quantity,omitempty"`

    // 唯一单据号，用于幂等操作
    BillNo   string `json:"bill_no,omitempty"`

    // 商品对应的69码
    Barcode   string `json:"barcode,omitempty"`

    // 更新类型，1表式覆盖，0表式增量
    UpdateType   int64 `json:"update_type,omitempty"`

    // 当时业务发生的时间戳，单位：ms
    OperationTs   int64 `json:"operation_ts,omitempty"`

}
