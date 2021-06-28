package wlb

// ItemListWlbWmsInventoryLackUpload 
/* model for simplify = false
type ItemListWlbWmsInventoryLackUpload struct {

    // 报缺原因（系统报缺，实物报缺）
    
    Reason   string `json:"reason,omitempty"`
    

    // 应发商品数量
    
    ItemQty   int64 `json:"item_qty,omitempty"`
    

    // 后端商品ID
    
    ItemId   string `json:"item_id,omitempty"`
    

    // 商品缺货数量
    
    LackQty   int64 `json:"lack_qty,omitempty"`
    

}
*/

// ItemListWlbWmsInventoryLackUpload 
type ItemListWlbWmsInventoryLackUpload struct {

    // 报缺原因（系统报缺，实物报缺）
    Reason   string `json:"reason,omitempty"`

    // 应发商品数量
    ItemQty   int64 `json:"item_qty,omitempty"`

    // 后端商品ID
    ItemId   string `json:"item_id,omitempty"`

    // 商品缺货数量
    LackQty   int64 `json:"lack_qty,omitempty"`

}
