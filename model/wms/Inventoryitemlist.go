package wms

// Inventoryitemlist 
/* model for simplify = false
type Inventoryitemlist struct {

    // 商品属性列表
    
    InventoryItem  *struct {
        Inventoryitem  *Inventoryitem `json:"inventoryitem,omitempty"`
    } `json:"inventory_item,omitempty"`
    

}
*/

// Inventoryitemlist 
type Inventoryitemlist struct {

    // 商品属性列表
    InventoryItem   *Inventoryitem `json:"inventory_item,omitempty"`

}
