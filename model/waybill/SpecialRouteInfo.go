package waybill

// SpecialRouteInfo 
/* model for simplify = false
type SpecialRouteInfo struct {

    // 快递公司code
    
    CpCode   string `json:"cp_code,omitempty"`
    

    // 到货区域
    
    ReceiveArea  *struct {
        AddressArea  *AddressArea `json:"address_area,omitempty"`
    } `json:"receive_area,omitempty"`
    

}
*/

// SpecialRouteInfo 
type SpecialRouteInfo struct {

    // 快递公司code
    CpCode   string `json:"cp_code,omitempty"`

    // 到货区域
    ReceiveArea   *AddressArea `json:"receive_area,omitempty"`

}
