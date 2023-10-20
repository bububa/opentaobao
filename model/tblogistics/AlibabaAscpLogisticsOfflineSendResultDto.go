package tblogistics

import (
	"sync"
)

// AlibabaAscpLogisticsOfflineSendResultDto 结构体
type AlibabaAscpLogisticsOfflineSendResultDto struct {
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAscpLogisticsOfflineSendResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsOfflineSendResultDto)
	},
}

// GetAlibabaAscpLogisticsOfflineSendResultDto() 从对象池中获取AlibabaAscpLogisticsOfflineSendResultDto
func GetAlibabaAscpLogisticsOfflineSendResultDto() *AlibabaAscpLogisticsOfflineSendResultDto {
	return poolAlibabaAscpLogisticsOfflineSendResultDto.Get().(*AlibabaAscpLogisticsOfflineSendResultDto)
}

// ReleaseAlibabaAscpLogisticsOfflineSendResultDto 释放AlibabaAscpLogisticsOfflineSendResultDto
func ReleaseAlibabaAscpLogisticsOfflineSendResultDto(v *AlibabaAscpLogisticsOfflineSendResultDto) {
	v.Success = false
	poolAlibabaAscpLogisticsOfflineSendResultDto.Put(v)
}
