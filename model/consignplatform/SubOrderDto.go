package consignplatform

// SubOrderDto 
type SubOrderDto struct {
    // 外部子订单id
    OuterOrderId   string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
    // 商品数量
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`
    // 商品价格（单位分）
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    // 商品图片链接
    PictureUrl   string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
    // 商品类别。1 日用品; 2 食品; 3 文件; 4 衣物; 5 数码产品; 6 其他
    Category   int64 `json:"category,omitempty" xml:"category,omitempty"`
    // 商品名称
    GoodsName   string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
    // 商品id
    GoodsId   string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}
