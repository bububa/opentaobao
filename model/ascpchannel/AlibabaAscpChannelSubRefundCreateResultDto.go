package ascpchannel

import (
	"sync"
)

// AlibabaAscpChannelSubRefundCreateResultDto 结构体
type AlibabaAscpChannelSubRefundCreateResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回主体
	Data *AlibabaAscpChannelSubRefundCreateData `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAscpChannelSubRefundCreateResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelSubRefundCreateResultDto)
	},
}

// GetAlibabaAscpChannelSubRefundCreateResultDto() 从对象池中获取AlibabaAscpChannelSubRefundCreateResultDto
func GetAlibabaAscpChannelSubRefundCreateResultDto() *AlibabaAscpChannelSubRefundCreateResultDto {
	return poolAlibabaAscpChannelSubRefundCreateResultDto.Get().(*AlibabaAscpChannelSubRefundCreateResultDto)
}

// ReleaseAlibabaAscpChannelSubRefundCreateResultDto 释放AlibabaAscpChannelSubRefundCreateResultDto
func ReleaseAlibabaAscpChannelSubRefundCreateResultDto(v *AlibabaAscpChannelSubRefundCreateResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = nil
	v.Success = false
	poolAlibabaAscpChannelSubRefundCreateResultDto.Put(v)
}
