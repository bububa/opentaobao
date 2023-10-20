package singletreasure

import (
	"sync"
)

// ItemDetailInfoBatchCreateDto 结构体
type ItemDetailInfoBatchCreateDto struct {
	// 商品列表
	ItemDetailInfo []ItemDetailInfoCreateDto `json:"item_detail_info,omitempty" xml:"item_detail_info>item_detail_info_create_dto,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

var poolItemDetailInfoBatchCreateDto = sync.Pool{
	New: func() any {
		return new(ItemDetailInfoBatchCreateDto)
	},
}

// GetItemDetailInfoBatchCreateDto() 从对象池中获取ItemDetailInfoBatchCreateDto
func GetItemDetailInfoBatchCreateDto() *ItemDetailInfoBatchCreateDto {
	return poolItemDetailInfoBatchCreateDto.Get().(*ItemDetailInfoBatchCreateDto)
}

// ReleaseItemDetailInfoBatchCreateDto 释放ItemDetailInfoBatchCreateDto
func ReleaseItemDetailInfoBatchCreateDto(v *ItemDetailInfoBatchCreateDto) {
	v.ItemDetailInfo = v.ItemDetailInfo[:0]
	v.ActivityId = 0
	poolItemDetailInfoBatchCreateDto.Put(v)
}
