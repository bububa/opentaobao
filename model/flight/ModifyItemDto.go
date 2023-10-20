package flight

import (
	"sync"
)

// ModifyItemDto 结构体
type ModifyItemDto struct {
	// 票号
	Tickets []string `json:"tickets,omitempty" xml:"tickets>string,omitempty"`
	// 改签后航段
	AfterChangeSegments []ModifySegmentDto `json:"after_change_segments,omitempty" xml:"after_change_segments>modify_segment_dto,omitempty"`
	// 改前航段
	BeforeChangeSegments []ModifyBeforeSegmentDto `json:"before_change_segments,omitempty" xml:"before_change_segments>modify_before_segment_dto,omitempty"`
	// 票号
	TicketNos []string `json:"ticket_nos,omitempty" xml:"ticket_nos>string,omitempty"`
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// pnr(必填),无编码出票时可不填
	Pnr string `json:"pnr,omitempty" xml:"pnr,omitempty"`
	// 改签费用
	ModifyFee int64 `json:"modify_fee,omitempty" xml:"modify_fee,omitempty"`
	// 升舱费用
	UpgradeFee int64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
}

var poolModifyItemDto = sync.Pool{
	New: func() any {
		return new(ModifyItemDto)
	},
}

// GetModifyItemDto() 从对象池中获取ModifyItemDto
func GetModifyItemDto() *ModifyItemDto {
	return poolModifyItemDto.Get().(*ModifyItemDto)
}

// ReleaseModifyItemDto 释放ModifyItemDto
func ReleaseModifyItemDto(v *ModifyItemDto) {
	v.Tickets = v.Tickets[:0]
	v.AfterChangeSegments = v.AfterChangeSegments[:0]
	v.BeforeChangeSegments = v.BeforeChangeSegments[:0]
	v.TicketNos = v.TicketNos[:0]
	v.PassengerName = ""
	v.Pnr = ""
	v.ModifyFee = 0
	v.UpgradeFee = 0
	poolModifyItemDto.Put(v)
}
