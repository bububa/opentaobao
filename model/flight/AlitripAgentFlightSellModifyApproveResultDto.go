package flight

import (
	"sync"
)

// AlitripAgentFlightSellModifyApproveResultDto 结构体
type AlitripAgentFlightSellModifyApproveResultDto struct {
	// 错误码:000:系统异常, 001:请求参数不合法, 002:权限不足, 003:操作失败, 004:流量管控
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentFlightSellModifyApproveResultDto = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellModifyApproveResultDto)
	},
}

// GetAlitripAgentFlightSellModifyApproveResultDto() 从对象池中获取AlitripAgentFlightSellModifyApproveResultDto
func GetAlitripAgentFlightSellModifyApproveResultDto() *AlitripAgentFlightSellModifyApproveResultDto {
	return poolAlitripAgentFlightSellModifyApproveResultDto.Get().(*AlitripAgentFlightSellModifyApproveResultDto)
}

// ReleaseAlitripAgentFlightSellModifyApproveResultDto 释放AlitripAgentFlightSellModifyApproveResultDto
func ReleaseAlitripAgentFlightSellModifyApproveResultDto(v *AlitripAgentFlightSellModifyApproveResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripAgentFlightSellModifyApproveResultDto.Put(v)
}
