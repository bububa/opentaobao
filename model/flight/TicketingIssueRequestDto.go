package flight

import (
	"sync"
)

// TicketingIssueRequestDto 结构体
type TicketingIssueRequestDto struct {
	// 出票信息
	IssueList []TicketingPsgItemDto `json:"issue_list,omitempty" xml:"issue_list>ticketing_psg_item_dto,omitempty"`
	// 飞猪订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 国内国际标识:1:国内,2:国际
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
}

var poolTicketingIssueRequestDto = sync.Pool{
	New: func() any {
		return new(TicketingIssueRequestDto)
	},
}

// GetTicketingIssueRequestDto() 从对象池中获取TicketingIssueRequestDto
func GetTicketingIssueRequestDto() *TicketingIssueRequestDto {
	return poolTicketingIssueRequestDto.Get().(*TicketingIssueRequestDto)
}

// ReleaseTicketingIssueRequestDto 释放TicketingIssueRequestDto
func ReleaseTicketingIssueRequestDto(v *TicketingIssueRequestDto) {
	v.IssueList = v.IssueList[:0]
	v.OrderId = ""
	v.DomesticIntl = 0
	poolTicketingIssueRequestDto.Put(v)
}
