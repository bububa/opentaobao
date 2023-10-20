package moscm

import (
	"sync"
)

// AlibabaMosOrderRefundListGetResultDto 结构体
type AlibabaMosOrderRefundListGetResultDto struct {
	// 状态码
	SubCode string `json:"sub_code,omitempty" xml:"sub_code,omitempty"`
	// 信息
	SubMsg string `json:"sub_msg,omitempty" xml:"sub_msg,omitempty"`
	// 成功标志
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 结果集
	Data *PagedList `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaMosOrderRefundListGetResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaMosOrderRefundListGetResultDto)
	},
}

// GetAlibabaMosOrderRefundListGetResultDto() 从对象池中获取AlibabaMosOrderRefundListGetResultDto
func GetAlibabaMosOrderRefundListGetResultDto() *AlibabaMosOrderRefundListGetResultDto {
	return poolAlibabaMosOrderRefundListGetResultDto.Get().(*AlibabaMosOrderRefundListGetResultDto)
}

// ReleaseAlibabaMosOrderRefundListGetResultDto 释放AlibabaMosOrderRefundListGetResultDto
func ReleaseAlibabaMosOrderRefundListGetResultDto(v *AlibabaMosOrderRefundListGetResultDto) {
	v.SubCode = ""
	v.SubMsg = ""
	v.Success = ""
	v.Data = nil
	poolAlibabaMosOrderRefundListGetResultDto.Put(v)
}
