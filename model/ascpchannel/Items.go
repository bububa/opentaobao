package ascpchannel

// Items 
type Items struct {
    // 商品编码
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    // 商品仓储系统编码
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 包裹内该商品的数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
