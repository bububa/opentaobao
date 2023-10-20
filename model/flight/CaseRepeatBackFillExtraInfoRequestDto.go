package flight

import (
	"sync"
)

// CaseRepeatBackFillExtraInfoRequestDto 结构体
type CaseRepeatBackFillExtraInfoRequestDto struct {
	// 票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 人商品Id
	PassengerItemId int64 `json:"passenger_item_id,omitempty" xml:"passenger_item_id,omitempty"`
}

var poolCaseRepeatBackFillExtraInfoRequestDto = sync.Pool{
	New: func() any {
		return new(CaseRepeatBackFillExtraInfoRequestDto)
	},
}

// GetCaseRepeatBackFillExtraInfoRequestDto() 从对象池中获取CaseRepeatBackFillExtraInfoRequestDto
func GetCaseRepeatBackFillExtraInfoRequestDto() *CaseRepeatBackFillExtraInfoRequestDto {
	return poolCaseRepeatBackFillExtraInfoRequestDto.Get().(*CaseRepeatBackFillExtraInfoRequestDto)
}

// ReleaseCaseRepeatBackFillExtraInfoRequestDto 释放CaseRepeatBackFillExtraInfoRequestDto
func ReleaseCaseRepeatBackFillExtraInfoRequestDto(v *CaseRepeatBackFillExtraInfoRequestDto) {
	v.TicketNo = ""
	v.PassengerItemId = 0
	poolCaseRepeatBackFillExtraInfoRequestDto.Put(v)
}
