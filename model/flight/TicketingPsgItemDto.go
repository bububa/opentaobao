package flight

import (
	"sync"
)

// TicketingPsgItemDto 结构体
type TicketingPsgItemDto struct {
	// 票号
	Tickets []string `json:"tickets,omitempty" xml:"tickets>string,omitempty"`
	// 航段
	Segments []Segments `json:"segments,omitempty" xml:"segments>segments,omitempty"`
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// pnr
	Pnr string `json:"pnr,omitempty" xml:"pnr,omitempty"`
}

var poolTicketingPsgItemDto = sync.Pool{
	New: func() any {
		return new(TicketingPsgItemDto)
	},
}

// GetTicketingPsgItemDto() 从对象池中获取TicketingPsgItemDto
func GetTicketingPsgItemDto() *TicketingPsgItemDto {
	return poolTicketingPsgItemDto.Get().(*TicketingPsgItemDto)
}

// ReleaseTicketingPsgItemDto 释放TicketingPsgItemDto
func ReleaseTicketingPsgItemDto(v *TicketingPsgItemDto) {
	v.Tickets = v.Tickets[:0]
	v.Segments = v.Segments[:0]
	v.PassengerName = ""
	v.Pnr = ""
	poolTicketingPsgItemDto.Put(v)
}
