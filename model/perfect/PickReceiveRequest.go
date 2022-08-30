package perfect

// PickReceiveRequest 结构体
type PickReceiveRequest struct {
	// 拣货单参数
	PickOrders []PickOrderRequest `json:"pick_orders,omitempty" xml:"pick_orders>pick_order_request,omitempty"`
	// 1
	OutboundOrders []PickOutboundOrderRequest `json:"outbound_orders,omitempty" xml:"outbound_orders>pick_outbound_order_request,omitempty"`
	// 1
	DockBarcode string `json:"dock_barcode,omitempty" xml:"dock_barcode,omitempty"`
	// 1
	WorkMode string `json:"work_mode,omitempty" xml:"work_mode,omitempty"`
	// 1
	WaveCode string `json:"wave_code,omitempty" xml:"wave_code,omitempty"`
	// 1
	DockType string `json:"dock_type,omitempty" xml:"dock_type,omitempty"`
	// 1
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 1
	DockCode string `json:"dock_code,omitempty" xml:"dock_code,omitempty"`
	// 1
	CargoOwnerCode string `json:"cargo_owner_code,omitempty" xml:"cargo_owner_code,omitempty"`
	// 1
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}
