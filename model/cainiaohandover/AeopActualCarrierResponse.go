package cainiaohandover

// AeopActualCarrierResponse 
type AeopActualCarrierResponse struct {
    // 实际承运商
    CourierList   []Courierlist `json:"courier_list,omitempty" xml:"courier_list>courierlist,omitempty"`
}
