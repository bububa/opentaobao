package wms

// Consignorderitemlist 
/* model for simplify = false
type Consignorderitemlist struct {

    // 仓库物流订单信息列表
    
    ConsignOrderItem  *struct {
        Consignorderitem  *Consignorderitem `json:"consignorderitem,omitempty"`
    } `json:"consign_order_item,omitempty"`
    

}
*/

// Consignorderitemlist 
type Consignorderitemlist struct {

    // 仓库物流订单信息列表
    ConsignOrderItem   *Consignorderitem `json:"consign_order_item,omitempty"`

}
