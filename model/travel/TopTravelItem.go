package travel

// TopTravelItem 
type TopTravelItem struct {

    // 商品id
    ItemId   int64 `json:"item_id,omitempty"`

    // 商家自定义商品编码
    OutProductId   string `json:"out_product_id,omitempty"`

    // 商品修改时间
    Modified   string `json:"modified,omitempty"`

    // 扩展信息
    Extend   string `json:"extend,omitempty"`

    // 商品创建时间
    Created   string `json:"created,omitempty"`

    // skuid
    SkuId   int64 `json:"sku_id,omitempty"`

    // 商家编码
    OuterId   string `json:"outer_id,omitempty"`

}
