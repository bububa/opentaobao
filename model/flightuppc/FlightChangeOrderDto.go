package flightuppc

import (
	"sync"
)

// FlightChangeOrderDto 结构体
type FlightChangeOrderDto struct {
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 航变信息
	FlightChange *FlightChangeDto `json:"flight_change,omitempty" xml:"flight_change,omitempty"`
}

var poolFlightChangeOrderDto = sync.Pool{
	New: func() any {
		return new(FlightChangeOrderDto)
	},
}

// GetFlightChangeOrderDto() 从对象池中获取FlightChangeOrderDto
func GetFlightChangeOrderDto() *FlightChangeOrderDto {
	return poolFlightChangeOrderDto.Get().(*FlightChangeOrderDto)
}

// ReleaseFlightChangeOrderDto 释放FlightChangeOrderDto
func ReleaseFlightChangeOrderDto(v *FlightChangeOrderDto) {
	v.OrderId = 0
	v.FlightChange = nil
	poolFlightChangeOrderDto.Put(v)
}
