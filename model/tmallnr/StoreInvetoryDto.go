package tmallnr

// StoreInvetoryDto 
type StoreInvetoryDto struct {
    // 商家的外部商品编码，写入值。
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // sku号
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // 天猫商品id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 天猫后端商品id
    ScItemId   int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    // 单据流水号，用于幂等操作
    BillNum   string `json:"bill_num,omitempty" xml:"bill_num,omitempty"`
    // CERTAINTY 表示确定性库存
    InventoryType   string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
    // 库存数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
