package flight

import (
	"sync"
)

// CaseChangePassengerExtraInfoRequestDto 结构体
type CaseChangePassengerExtraInfoRequestDto struct {
	// 票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 手工单人ID
	ManualPassengerId int64 `json:"manual_passenger_id,omitempty" xml:"manual_passenger_id,omitempty"`
}

var poolCaseChangePassengerExtraInfoRequestDto = sync.Pool{
	New: func() any {
		return new(CaseChangePassengerExtraInfoRequestDto)
	},
}

// GetCaseChangePassengerExtraInfoRequestDto() 从对象池中获取CaseChangePassengerExtraInfoRequestDto
func GetCaseChangePassengerExtraInfoRequestDto() *CaseChangePassengerExtraInfoRequestDto {
	return poolCaseChangePassengerExtraInfoRequestDto.Get().(*CaseChangePassengerExtraInfoRequestDto)
}

// ReleaseCaseChangePassengerExtraInfoRequestDto 释放CaseChangePassengerExtraInfoRequestDto
func ReleaseCaseChangePassengerExtraInfoRequestDto(v *CaseChangePassengerExtraInfoRequestDto) {
	v.TicketNo = ""
	v.ManualPassengerId = 0
	poolCaseChangePassengerExtraInfoRequestDto.Put(v)
}
