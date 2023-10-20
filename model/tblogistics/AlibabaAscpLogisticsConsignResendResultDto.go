package tblogistics

import (
	"sync"
)

// AlibabaAscpLogisticsConsignResendResultDto 结构体
type AlibabaAscpLogisticsConsignResendResultDto struct {
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAscpLogisticsConsignResendResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsConsignResendResultDto)
	},
}

// GetAlibabaAscpLogisticsConsignResendResultDto() 从对象池中获取AlibabaAscpLogisticsConsignResendResultDto
func GetAlibabaAscpLogisticsConsignResendResultDto() *AlibabaAscpLogisticsConsignResendResultDto {
	return poolAlibabaAscpLogisticsConsignResendResultDto.Get().(*AlibabaAscpLogisticsConsignResendResultDto)
}

// ReleaseAlibabaAscpLogisticsConsignResendResultDto 释放AlibabaAscpLogisticsConsignResendResultDto
func ReleaseAlibabaAscpLogisticsConsignResendResultDto(v *AlibabaAscpLogisticsConsignResendResultDto) {
	v.Success = false
	poolAlibabaAscpLogisticsConsignResendResultDto.Put(v)
}
