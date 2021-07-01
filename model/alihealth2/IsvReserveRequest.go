package alihealth2

// IsvReserveRequest 结构体
type IsvReserveRequest struct {
	// 预约单ID
	ReserveId int64 `json:"reserve_id,omitempty" xml:"reserve_id,omitempty"`
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// isv 预约单ID
	SpReserveId string `json:"sp_reserve_id,omitempty" xml:"sp_reserve_id,omitempty"`
	// 预约时间
	ReserveTime string `json:"reserve_time,omitempty" xml:"reserve_time,omitempty"`
}
