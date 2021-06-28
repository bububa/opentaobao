package wms

// Consignsendinfolist 
/* model for simplify = false
type Consignsendinfolist struct {

    // 物流订单发货信息
    
    ConsignSendInfo  *struct {
        Consignsendinfo  *Consignsendinfo `json:"consignsendinfo,omitempty"`
    } `json:"consign_send_info,omitempty"`
    

}
*/

// Consignsendinfolist 
type Consignsendinfolist struct {

    // 物流订单发货信息
    ConsignSendInfo   *Consignsendinfo `json:"consign_send_info,omitempty"`

}
