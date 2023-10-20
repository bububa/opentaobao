package ascpchannel

import (
	"sync"
)

// AlibabaAscpChannelDistributorInventoryGetResultDto 结构体
type AlibabaAscpChannelDistributorInventoryGetResultDto struct {
	// 结果
	Datas []ChannelInventoryDto `json:"datas,omitempty" xml:"datas>channel_inventory_dto,omitempty"`
	// 调用链路ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAscpChannelDistributorInventoryGetResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelDistributorInventoryGetResultDto)
	},
}

// GetAlibabaAscpChannelDistributorInventoryGetResultDto() 从对象池中获取AlibabaAscpChannelDistributorInventoryGetResultDto
func GetAlibabaAscpChannelDistributorInventoryGetResultDto() *AlibabaAscpChannelDistributorInventoryGetResultDto {
	return poolAlibabaAscpChannelDistributorInventoryGetResultDto.Get().(*AlibabaAscpChannelDistributorInventoryGetResultDto)
}

// ReleaseAlibabaAscpChannelDistributorInventoryGetResultDto 释放AlibabaAscpChannelDistributorInventoryGetResultDto
func ReleaseAlibabaAscpChannelDistributorInventoryGetResultDto(v *AlibabaAscpChannelDistributorInventoryGetResultDto) {
	v.Datas = v.Datas[:0]
	v.TraceId = ""
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaAscpChannelDistributorInventoryGetResultDto.Put(v)
}
