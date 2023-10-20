package ascpchannel

import (
	"sync"
)

// AlibabaAscpChannelRefundGoodsWaybillResultDto 结构体
type AlibabaAscpChannelRefundGoodsWaybillResultDto struct {
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAscpChannelRefundGoodsWaybillResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelRefundGoodsWaybillResultDto)
	},
}

// GetAlibabaAscpChannelRefundGoodsWaybillResultDto() 从对象池中获取AlibabaAscpChannelRefundGoodsWaybillResultDto
func GetAlibabaAscpChannelRefundGoodsWaybillResultDto() *AlibabaAscpChannelRefundGoodsWaybillResultDto {
	return poolAlibabaAscpChannelRefundGoodsWaybillResultDto.Get().(*AlibabaAscpChannelRefundGoodsWaybillResultDto)
}

// ReleaseAlibabaAscpChannelRefundGoodsWaybillResultDto 释放AlibabaAscpChannelRefundGoodsWaybillResultDto
func ReleaseAlibabaAscpChannelRefundGoodsWaybillResultDto(v *AlibabaAscpChannelRefundGoodsWaybillResultDto) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaAscpChannelRefundGoodsWaybillResultDto.Put(v)
}
