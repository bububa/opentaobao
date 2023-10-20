package flight

import (
	"sync"
)

// AlitripAgentFlightSellTicketingDetailResultDto 结构体
type AlitripAgentFlightSellTicketingDetailResultDto struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 出参对象
	Data *TicketingDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentFlightSellTicketingDetailResultDto = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellTicketingDetailResultDto)
	},
}

// GetAlitripAgentFlightSellTicketingDetailResultDto() 从对象池中获取AlitripAgentFlightSellTicketingDetailResultDto
func GetAlitripAgentFlightSellTicketingDetailResultDto() *AlitripAgentFlightSellTicketingDetailResultDto {
	return poolAlitripAgentFlightSellTicketingDetailResultDto.Get().(*AlitripAgentFlightSellTicketingDetailResultDto)
}

// ReleaseAlitripAgentFlightSellTicketingDetailResultDto 释放AlitripAgentFlightSellTicketingDetailResultDto
func ReleaseAlitripAgentFlightSellTicketingDetailResultDto(v *AlitripAgentFlightSellTicketingDetailResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolAlitripAgentFlightSellTicketingDetailResultDto.Put(v)
}
