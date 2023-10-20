package train

import (
	"sync"
)

// ChangeTicketDto 结构体
type ChangeTicketDto struct {
	// 改签票坐席类型名
	ChangeSeatTypeName string `json:"change_seat_type_name,omitempty" xml:"change_seat_type_name,omitempty"`
	// 改签票坐席类型code
	ChangeSeatTypeCode string `json:"change_seat_type_code,omitempty" xml:"change_seat_type_code,omitempty"`
	// 改签票坐席号
	ChangeSeatNo string `json:"change_seat_no,omitempty" xml:"change_seat_no,omitempty"`
	// 改签票车厢号
	ChangeCoachNo string `json:"change_coach_no,omitempty" xml:"change_coach_no,omitempty"`
	// 票类型code
	TicketTypeCode string `json:"ticket_type_code,omitempty" xml:"ticket_type_code,omitempty"`
	// 改签票车厢名
	ChangeCoachName string `json:"change_coach_name,omitempty" xml:"change_coach_name,omitempty"`
	// 票类型名称
	TicketTypeName string `json:"ticket_type_name,omitempty" xml:"ticket_type_name,omitempty"`
	// 改签票票价，单位分
	ChangeTicketPrice string `json:"change_ticket_price,omitempty" xml:"change_ticket_price,omitempty"`
	// 改签申请单id
	ChangeApplyId int64 `json:"change_apply_id,omitempty" xml:"change_apply_id,omitempty"`
	// 改签手续费，单位分
	HandingFee int64 `json:"handing_fee,omitempty" xml:"handing_fee,omitempty"`
	// 子单单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

var poolChangeTicketDto = sync.Pool{
	New: func() any {
		return new(ChangeTicketDto)
	},
}

// GetChangeTicketDto() 从对象池中获取ChangeTicketDto
func GetChangeTicketDto() *ChangeTicketDto {
	return poolChangeTicketDto.Get().(*ChangeTicketDto)
}

// ReleaseChangeTicketDto 释放ChangeTicketDto
func ReleaseChangeTicketDto(v *ChangeTicketDto) {
	v.ChangeSeatTypeName = ""
	v.ChangeSeatTypeCode = ""
	v.ChangeSeatNo = ""
	v.ChangeCoachNo = ""
	v.TicketTypeCode = ""
	v.ChangeCoachName = ""
	v.TicketTypeName = ""
	v.ChangeTicketPrice = ""
	v.ChangeApplyId = 0
	v.HandingFee = 0
	v.SubOrderId = 0
	poolChangeTicketDto.Put(v)
}
