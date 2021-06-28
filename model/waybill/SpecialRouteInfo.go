package waybill

// SpecialRouteInfo 
type SpecialRouteInfo struct {

    // 快递公司code
    
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    

    // 到货区域
    
    ReceiveArea   *AddressArea `json:"receive_area,omitempty" xml:"receive_area,omitempty"`
    

}
