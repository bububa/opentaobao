package wdk

import (
	"sync"
)

// SkuChannelConfigDto 结构体
type SkuChannelConfigDto struct {
	// 渠道对应的门店id列表[&#34;store001&#34;,&#34;store002&#34;]
	StoreIds []string `json:"store_ids,omitempty" xml:"store_ids>string,omitempty"`
	// 渠道编码 枚举： 2 美团 3 饿了么 26 京东到家 31 翱象淘鲜达
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
}

var poolSkuChannelConfigDto = sync.Pool{
	New: func() any {
		return new(SkuChannelConfigDto)
	},
}

// GetSkuChannelConfigDto() 从对象池中获取SkuChannelConfigDto
func GetSkuChannelConfigDto() *SkuChannelConfigDto {
	return poolSkuChannelConfigDto.Get().(*SkuChannelConfigDto)
}

// ReleaseSkuChannelConfigDto 释放SkuChannelConfigDto
func ReleaseSkuChannelConfigDto(v *SkuChannelConfigDto) {
	v.StoreIds = v.StoreIds[:0]
	v.Channel = ""
	poolSkuChannelConfigDto.Put(v)
}
