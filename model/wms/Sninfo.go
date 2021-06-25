package wms

// Sninfo 
type Sninfo struct {

    // 商品ID
    ItemId   string `json:"item_id,omitempty"`

    // 商家对商品的编码
    ItemCode   string `json:"item_code,omitempty"`

    // 库存类型（1 可销售库存(正品) 101 残次 102 机损 103 箱损201 冻结库存）
    InventoryType   int64 `json:"inventory_type,omitempty"`

    // sn编码
    SnCode   string `json:"sn_code,omitempty"`

}
