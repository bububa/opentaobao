package idle

import (
	"sync"
)

// LogisticsDto 结构体
type LogisticsDto struct {
	// 快递/物流名称
	LogisticsName string `json:"logistics_name,omitempty" xml:"logistics_name,omitempty"`
	// 物流id
	LogisticsId string `json:"logistics_id,omitempty" xml:"logistics_id,omitempty"`
	// 发货人地址信息
	SenderAddress *UserAddressDto `json:"sender_address,omitempty" xml:"sender_address,omitempty"`
}

var poolLogisticsDto = sync.Pool{
	New: func() any {
		return new(LogisticsDto)
	},
}

// GetLogisticsDto() 从对象池中获取LogisticsDto
func GetLogisticsDto() *LogisticsDto {
	return poolLogisticsDto.Get().(*LogisticsDto)
}

// ReleaseLogisticsDto 释放LogisticsDto
func ReleaseLogisticsDto(v *LogisticsDto) {
	v.LogisticsName = ""
	v.LogisticsId = ""
	v.SenderAddress = nil
	poolLogisticsDto.Put(v)
}
