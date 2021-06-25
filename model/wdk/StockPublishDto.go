package wdk

// StockPublishDto 
type StockPublishDto struct {

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 更新数量，矢量
    Quantity   string `json:"quantity,omitempty"`

    // 订单号（商品粒度操作单）
    OrderNo   string `json:"order_no,omitempty"`

    // 订单类型，参见订单类型约束
    OrderType   string `json:"order_type,omitempty"`

    // 订单描述
    OrderDesc   string `json:"order_desc,omitempty"`

}
