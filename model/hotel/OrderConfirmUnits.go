package hotel

// OrderConfirmUnits 结构体
type OrderConfirmUnits struct {
	// h5OrderConfirmUrl
	H5OrderConfirmUrl string `json:"h5_order_confirm_url,omitempty" xml:"h5_order_confirm_url,omitempty"`
	// pcOrderConfirmUrl
	PcOrderConfirmUrl string `json:"pc_order_confirm_url,omitempty" xml:"pc_order_confirm_url,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
