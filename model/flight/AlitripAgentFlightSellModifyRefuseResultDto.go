package flight

import (
	"sync"
)

// AlitripAgentFlightSellModifyRefuseResultDto 结构体
type AlitripAgentFlightSellModifyRefuseResultDto struct {
	// 错误码:000:系统异常, 001:请求参数不合法, 002:权限不足, 003:操作失败, 004:流量管控
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentFlightSellModifyRefuseResultDto = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellModifyRefuseResultDto)
	},
}

// GetAlitripAgentFlightSellModifyRefuseResultDto() 从对象池中获取AlitripAgentFlightSellModifyRefuseResultDto
func GetAlitripAgentFlightSellModifyRefuseResultDto() *AlitripAgentFlightSellModifyRefuseResultDto {
	return poolAlitripAgentFlightSellModifyRefuseResultDto.Get().(*AlitripAgentFlightSellModifyRefuseResultDto)
}

// ReleaseAlitripAgentFlightSellModifyRefuseResultDto 释放AlitripAgentFlightSellModifyRefuseResultDto
func ReleaseAlitripAgentFlightSellModifyRefuseResultDto(v *AlitripAgentFlightSellModifyRefuseResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlitripAgentFlightSellModifyRefuseResultDto.Put(v)
}
