package flight

import (
	"sync"
)

// AlitripAgentFlightSellTicketingIssueResultDto 结构体
type AlitripAgentFlightSellTicketingIssueResultDto struct {
	// 错误码:000:系统异常, 001:请求参数不合法, 002:权限不足, 003:操作失败, 004:流量管控
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentFlightSellTicketingIssueResultDto = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellTicketingIssueResultDto)
	},
}

// GetAlitripAgentFlightSellTicketingIssueResultDto() 从对象池中获取AlitripAgentFlightSellTicketingIssueResultDto
func GetAlitripAgentFlightSellTicketingIssueResultDto() *AlitripAgentFlightSellTicketingIssueResultDto {
	return poolAlitripAgentFlightSellTicketingIssueResultDto.Get().(*AlitripAgentFlightSellTicketingIssueResultDto)
}

// ReleaseAlitripAgentFlightSellTicketingIssueResultDto 释放AlitripAgentFlightSellTicketingIssueResultDto
func ReleaseAlitripAgentFlightSellTicketingIssueResultDto(v *AlitripAgentFlightSellTicketingIssueResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripAgentFlightSellTicketingIssueResultDto.Put(v)
}
