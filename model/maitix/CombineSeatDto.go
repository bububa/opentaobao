package maitix

import (
	"sync"
)

// CombineSeatDto 结构体
type CombineSeatDto struct {
	// 套票下的单票信息
	OrdinaryTickets []OrdinarySeatDto `json:"ordinary_tickets,omitempty" xml:"ordinary_tickets>ordinary_seat_dto,omitempty"`
	// 套票价格，单位为分
	CombineTicketPrice string `json:"combine_ticket_price,omitempty" xml:"combine_ticket_price,omitempty"`
	// 套票名称
	CombineTicketName string `json:"combine_ticket_name,omitempty" xml:"combine_ticket_name,omitempty"`
	// 套票组合ID,和下单参数的combine_id一个意思
	ThirdCombineJointTicketId int64 `json:"third_combine_joint_ticket_id,omitempty" xml:"third_combine_joint_ticket_id,omitempty"`
	// 套票ID,如果是套票。下单的时候传这个给ticket_item_id,和下面的combine_ticket_id一样
	ThirdCombineTicketId int64 `json:"third_combine_ticket_id,omitempty" xml:"third_combine_ticket_id,omitempty"`
	// 套票ID,如果是套票。下单的时候传这个给ticket_item_id,和上面的third_combine_ticket_id一样
	CombineTicketId int64 `json:"combine_ticket_id,omitempty" xml:"combine_ticket_id,omitempty"`
}

var poolCombineSeatDto = sync.Pool{
	New: func() any {
		return new(CombineSeatDto)
	},
}

// GetCombineSeatDto() 从对象池中获取CombineSeatDto
func GetCombineSeatDto() *CombineSeatDto {
	return poolCombineSeatDto.Get().(*CombineSeatDto)
}

// ReleaseCombineSeatDto 释放CombineSeatDto
func ReleaseCombineSeatDto(v *CombineSeatDto) {
	v.OrdinaryTickets = v.OrdinaryTickets[:0]
	v.CombineTicketPrice = ""
	v.CombineTicketName = ""
	v.ThirdCombineJointTicketId = 0
	v.ThirdCombineTicketId = 0
	v.CombineTicketId = 0
	poolCombineSeatDto.Put(v)
}
