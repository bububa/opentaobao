package crm

import (
	"sync"
)

// ExchangeActivityCreateDto 结构体
type ExchangeActivityCreateDto struct {
	// 不包邮地区
	ExcludeArea string `json:"exclude_area,omitempty" xml:"exclude_area,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 优惠标签
	ActivityTag string `json:"activity_tag,omitempty" xml:"activity_tag,omitempty"`
	// 活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 商品ID
	ItmeId int64 `json:"itme_id,omitempty" xml:"itme_id,omitempty"`
	// 商品一口价
	FixPrice int64 `json:"fix_price,omitempty" xml:"fix_price,omitempty"`
	// 商品是否包邮
	FreePostage bool `json:"free_postage,omitempty" xml:"free_postage,omitempty"`
}

var poolExchangeActivityCreateDto = sync.Pool{
	New: func() any {
		return new(ExchangeActivityCreateDto)
	},
}

// GetExchangeActivityCreateDto() 从对象池中获取ExchangeActivityCreateDto
func GetExchangeActivityCreateDto() *ExchangeActivityCreateDto {
	return poolExchangeActivityCreateDto.Get().(*ExchangeActivityCreateDto)
}

// ReleaseExchangeActivityCreateDto 释放ExchangeActivityCreateDto
func ReleaseExchangeActivityCreateDto(v *ExchangeActivityCreateDto) {
	v.ExcludeArea = ""
	v.EndTime = ""
	v.ActivityTag = ""
	v.StartTime = ""
	v.ActivityName = ""
	v.ItmeId = 0
	v.FixPrice = 0
	v.FreePostage = false
	poolExchangeActivityCreateDto.Put(v)
}
