package icbu

// InventoryDetail 
/* model for simplify = false
type InventoryDetail struct {

    // 仓库code，默认不填
    
    StoreCode   string `json:"store_code,omitempty"`
    

    // 想设置的库存
    
    CurrentInventory   int64 `json:"current_inventory,omitempty"`
    

    // 原始库存
    
    SrcInventory   int64 `json:"src_inventory,omitempty"`
    

}
*/

// InventoryDetail 
type InventoryDetail struct {

    // 仓库code，默认不填
    StoreCode   string `json:"store_code,omitempty"`

    // 想设置的库存
    CurrentInventory   int64 `json:"current_inventory,omitempty"`

    // 原始库存
    SrcInventory   int64 `json:"src_inventory,omitempty"`

}
