package cainiaohandover

// Courierlist 
type Courierlist struct {
    // 承运商名字
    CourierName   string `json:"courier_name,omitempty" xml:"courier_name,omitempty"`
    // 承运商code
    CourierCode   string `json:"courier_code,omitempty" xml:"courier_code,omitempty"`
}
