package wms

// Consignorderitem 
/* model for simplify = false
type Consignorderitem struct {

    // 商品数量
    
    ItemQuantity   int64 `json:"item_quantity,omitempty"`
    

    // ERP订单明细行号ID
    
    OrderItemId   string `json:"order_item_id,omitempty"`
    

    // 商品ID
    
    ItemId   string `json:"item_id,omitempty"`
    

    // 商品编码
    
    ItemCode   string `json:"item_code,omitempty"`
    

}
*/

// Consignorderitem 
type Consignorderitem struct {

    // 商品数量
    ItemQuantity   int64 `json:"item_quantity,omitempty"`

    // ERP订单明细行号ID
    OrderItemId   string `json:"order_item_id,omitempty"`

    // 商品ID
    ItemId   string `json:"item_id,omitempty"`

    // 商品编码
    ItemCode   string `json:"item_code,omitempty"`

}
