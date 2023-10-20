package singletreasure

import (
	"sync"
)

// ItemProcessErrorResultDto 结构体
type ItemProcessErrorResultDto struct {
	// 处理失败的item和sku: map结构为:itemId: skuId-错误信息,所有返回的外层key是itemId，里面的key是skuId，商品级别的skuId为-1
	SkuIdMap string `json:"sku_id_map,omitempty" xml:"sku_id_map,omitempty"`
	// activityId
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

var poolItemProcessErrorResultDto = sync.Pool{
	New: func() any {
		return new(ItemProcessErrorResultDto)
	},
}

// GetItemProcessErrorResultDto() 从对象池中获取ItemProcessErrorResultDto
func GetItemProcessErrorResultDto() *ItemProcessErrorResultDto {
	return poolItemProcessErrorResultDto.Get().(*ItemProcessErrorResultDto)
}

// ReleaseItemProcessErrorResultDto 释放ItemProcessErrorResultDto
func ReleaseItemProcessErrorResultDto(v *ItemProcessErrorResultDto) {
	v.SkuIdMap = ""
	v.ActivityId = 0
	poolItemProcessErrorResultDto.Put(v)
}
