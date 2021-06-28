package wms

// Orderitem 
/* model for simplify = false
type Orderitem struct {

    // 订单明细行编码
    
    OrderItemId   string `json:"order_item_id,omitempty"`
    

    // 交易单号
    
    OrderSourceCode   string `json:"order_source_code,omitempty"`
    

    // 商品ID
    
    ItemId   string `json:"item_id,omitempty"`
    

    // 商家编码
    
    ItemCode   string `json:"item_code,omitempty"`
    

    // 商品属性列表
    
    InventoryItemList  struct {
        Inventoryitemlist  []Inventoryitemlist `json:"inventoryitemlist,omitempty"`
    } `json:"inventory_item_list,omitempty"`
    

}
*/

// Orderitem 
type Orderitem struct {

    // 订单明细行编码
    OrderItemId   string `json:"order_item_id,omitempty"`

    // 交易单号
    OrderSourceCode   string `json:"order_source_code,omitempty"`

    // 商品ID
    ItemId   string `json:"item_id,omitempty"`

    // 商家编码
    ItemCode   string `json:"item_code,omitempty"`

    // 商品属性列表
    InventoryItemList   []Inventoryitemlist `json:"inventory_item_list,omitempty"`

}
