package flight

import (
	"sync"
)

// CaseExtraAddBabyRequestDto 结构体
type CaseExtraAddBabyRequestDto struct {
	// pnr号
	Pnr string `json:"pnr,omitempty" xml:"pnr,omitempty"`
	// 票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 人商品Id
	PassengerItemId int64 `json:"passenger_item_id,omitempty" xml:"passenger_item_id,omitempty"`
}

var poolCaseExtraAddBabyRequestDto = sync.Pool{
	New: func() any {
		return new(CaseExtraAddBabyRequestDto)
	},
}

// GetCaseExtraAddBabyRequestDto() 从对象池中获取CaseExtraAddBabyRequestDto
func GetCaseExtraAddBabyRequestDto() *CaseExtraAddBabyRequestDto {
	return poolCaseExtraAddBabyRequestDto.Get().(*CaseExtraAddBabyRequestDto)
}

// ReleaseCaseExtraAddBabyRequestDto 释放CaseExtraAddBabyRequestDto
func ReleaseCaseExtraAddBabyRequestDto(v *CaseExtraAddBabyRequestDto) {
	v.Pnr = ""
	v.TicketNo = ""
	v.PassengerItemId = 0
	poolCaseExtraAddBabyRequestDto.Put(v)
}
