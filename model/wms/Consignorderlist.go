package wms

// Consignorderlist 
/* model for simplify = false
type Consignorderlist struct {

    // 发货订单信息
    
    ConsignOrder  *struct {
        Consignorder  *Consignorder `json:"consignorder,omitempty"`
    } `json:"consign_order,omitempty"`
    

}
*/

// Consignorderlist 
type Consignorderlist struct {

    // 发货订单信息
    ConsignOrder   *Consignorder `json:"consign_order,omitempty"`

}
