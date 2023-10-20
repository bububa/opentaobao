package flight

import (
	"sync"
)

// AlitripAgentFlightSellRefundRefuseResultDto 结构体
type AlitripAgentFlightSellRefundRefuseResultDto struct {
	// 错误码:000:系统异常, 001:请求参数不合法, 002:权限不足, 003:操作失败, 004:流量管控
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentFlightSellRefundRefuseResultDto = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellRefundRefuseResultDto)
	},
}

// GetAlitripAgentFlightSellRefundRefuseResultDto() 从对象池中获取AlitripAgentFlightSellRefundRefuseResultDto
func GetAlitripAgentFlightSellRefundRefuseResultDto() *AlitripAgentFlightSellRefundRefuseResultDto {
	return poolAlitripAgentFlightSellRefundRefuseResultDto.Get().(*AlitripAgentFlightSellRefundRefuseResultDto)
}

// ReleaseAlitripAgentFlightSellRefundRefuseResultDto 释放AlitripAgentFlightSellRefundRefuseResultDto
func ReleaseAlitripAgentFlightSellRefundRefuseResultDto(v *AlitripAgentFlightSellRefundRefuseResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripAgentFlightSellRefundRefuseResultDto.Put(v)
}
