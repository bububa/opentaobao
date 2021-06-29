package simba

// ExtraAttributes 
type ExtraAttributes struct {
    // 商品在淘宝的发布时间
    PublishTime   string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
    // 库存数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 商品的累积销量
    SalesCount   int64 `json:"sales_count,omitempty" xml:"sales_count,omitempty"`
}
