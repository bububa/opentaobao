package ascpchannel

import (
	"sync"
)

// AlibabaAscpChannelDistributorInventoryListGetResultDto 结构体
type AlibabaAscpChannelDistributorInventoryListGetResultDto struct {
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

var poolAlibabaAscpChannelDistributorInventoryListGetResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelDistributorInventoryListGetResultDto)
	},
}

// GetAlibabaAscpChannelDistributorInventoryListGetResultDto() 从对象池中获取AlibabaAscpChannelDistributorInventoryListGetResultDto
func GetAlibabaAscpChannelDistributorInventoryListGetResultDto() *AlibabaAscpChannelDistributorInventoryListGetResultDto {
	return poolAlibabaAscpChannelDistributorInventoryListGetResultDto.Get().(*AlibabaAscpChannelDistributorInventoryListGetResultDto)
}

// ReleaseAlibabaAscpChannelDistributorInventoryListGetResultDto 释放AlibabaAscpChannelDistributorInventoryListGetResultDto
func ReleaseAlibabaAscpChannelDistributorInventoryListGetResultDto(v *AlibabaAscpChannelDistributorInventoryListGetResultDto) {
	v.Datas = v.Datas[:0]
	v.TraceId = ""
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaAscpChannelDistributorInventoryListGetResultDto.Put(v)
}
