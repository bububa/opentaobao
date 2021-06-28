package wdk

// Skudetails 
/* model for simplify = false
type Skudetails struct {

    // 履约子单id
    
    FulfillSubOrderId   string `json:"fulfill_sub_order_id,omitempty"`
    

    // 货品编码
    
    SkuCode   string `json:"sku_code,omitempty"`
    

}
*/

// Skudetails 
type Skudetails struct {

    // 履约子单id
    FulfillSubOrderId   string `json:"fulfill_sub_order_id,omitempty"`

    // 货品编码
    SkuCode   string `json:"sku_code,omitempty"`

}
