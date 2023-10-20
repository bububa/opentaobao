package ascpchannel

import (
	"sync"
)

// AlibabaAscpChannelDistributorPriceGetResultDto 结构体
type AlibabaAscpChannelDistributorPriceGetResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 价格数据
	Data *TopDistributorPriceResult `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAscpChannelDistributorPriceGetResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelDistributorPriceGetResultDto)
	},
}

// GetAlibabaAscpChannelDistributorPriceGetResultDto() 从对象池中获取AlibabaAscpChannelDistributorPriceGetResultDto
func GetAlibabaAscpChannelDistributorPriceGetResultDto() *AlibabaAscpChannelDistributorPriceGetResultDto {
	return poolAlibabaAscpChannelDistributorPriceGetResultDto.Get().(*AlibabaAscpChannelDistributorPriceGetResultDto)
}

// ReleaseAlibabaAscpChannelDistributorPriceGetResultDto 释放AlibabaAscpChannelDistributorPriceGetResultDto
func ReleaseAlibabaAscpChannelDistributorPriceGetResultDto(v *AlibabaAscpChannelDistributorPriceGetResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = nil
	v.Success = false
	poolAlibabaAscpChannelDistributorPriceGetResultDto.Put(v)
}
