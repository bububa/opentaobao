package flight

import (
	"sync"
)

// AlitripAgentFlightSellRefundDetailResultDto 结构体
type AlitripAgentFlightSellRefundDetailResultDto struct {
	// 错误码:000:系统异常, 001:请求参数不合法, 002:权限不足, 003:操作失败, 004:流量管控
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 结果数据
	Data *RefundDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripAgentFlightSellRefundDetailResultDto = sync.Pool{
	New: func() any {
		return new(AlitripAgentFlightSellRefundDetailResultDto)
	},
}

// GetAlitripAgentFlightSellRefundDetailResultDto() 从对象池中获取AlitripAgentFlightSellRefundDetailResultDto
func GetAlitripAgentFlightSellRefundDetailResultDto() *AlitripAgentFlightSellRefundDetailResultDto {
	return poolAlitripAgentFlightSellRefundDetailResultDto.Get().(*AlitripAgentFlightSellRefundDetailResultDto)
}

// ReleaseAlitripAgentFlightSellRefundDetailResultDto 释放AlitripAgentFlightSellRefundDetailResultDto
func ReleaseAlitripAgentFlightSellRefundDetailResultDto(v *AlitripAgentFlightSellRefundDetailResultDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolAlitripAgentFlightSellRefundDetailResultDto.Put(v)
}
