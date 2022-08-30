package waybill

// WaybillOrderConfirmRequest 结构体
type WaybillOrderConfirmRequest struct {
	// 物流单号信息
	WaybillInfo []WaybillOrderConfirmWaybillInfo `json:"waybill_info,omitempty" xml:"waybill_info>waybill_order_confirm_waybill_info,omitempty"`
	// cpCode
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 预约上门截止时间
	DoorPickUpEndTime string `json:"door_pick_up_end_time,omitempty" xml:"door_pick_up_end_time,omitempty"`
	// 预约上门时间
	DoorPickUpTime string `json:"door_pick_up_time,omitempty" xml:"door_pick_up_time,omitempty"`
	// 扩展信息，json String
	ExtraInfo string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	// 物流服务， json String
	LogisticsServices string `json:"logistics_services,omitempty" xml:"logistics_services,omitempty"`
	// 快递产品编码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 客户订单货物总高，单位厘米
	TotalHeight int64 `json:"total_height,omitempty" xml:"total_height,omitempty"`
	// 订单货物总长,单位厘米
	TotalLength int64 `json:"total_length,omitempty" xml:"total_length,omitempty"`
	// 订单货物总宽，单位厘米
	TotalWidth int64 `json:"total_width,omitempty" xml:"total_width,omitempty"`
	// 货物总体积，单位立方厘米
	TotalVolume int64 `json:"total_volume,omitempty" xml:"total_volume,omitempty"`
	// 货物总重量，单位g
	TotalWeight int64 `json:"total_weight,omitempty" xml:"total_weight,omitempty"`
	// 预约上门收件
	CallDoorPickUp bool `json:"call_door_pick_up,omitempty" xml:"call_door_pick_up,omitempty"`
}
