package perfect

// OutboundOrderCancelDto 结构体
type OutboundOrderCancelDto struct {
	// 1
	WorkMode string `json:"work_mode,omitempty" xml:"work_mode,omitempty"`
	// 1
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 11
	SubOrderCode string `json:"sub_order_code,omitempty" xml:"sub_order_code,omitempty"`
	// 1
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	InterceptStatus int64 `json:"intercept_status,omitempty" xml:"intercept_status,omitempty"`
	// 1
	WarehouseId int64 `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
	// 1
	CancelSuccess bool `json:"cancel_success,omitempty" xml:"cancel_success,omitempty"`
}
