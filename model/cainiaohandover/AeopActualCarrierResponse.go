package cainiaohandover

// AeopActualCarrierResponse 结构体
type AeopActualCarrierResponse struct {
	// 实际承运商
	CourierList []Courierlist `json:"courier_list,omitempty" xml:"courier_list>courierlist,omitempty"`
}
