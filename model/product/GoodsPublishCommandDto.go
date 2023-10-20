package product

import (
	"sync"
)

// GoodsPublishCommandDto 结构体
type GoodsPublishCommandDto struct {
	// 商品发布数据体
	GoodsList []GoodsPublishDto `json:"goods_list,omitempty" xml:"goods_list>goods_publish_dto,omitempty"`
	// 外部批次ID，用于幂等
	ExternalBatchId string `json:"external_batch_id,omitempty" xml:"external_batch_id,omitempty"`
}

var poolGoodsPublishCommandDto = sync.Pool{
	New: func() any {
		return new(GoodsPublishCommandDto)
	},
}

// GetGoodsPublishCommandDto() 从对象池中获取GoodsPublishCommandDto
func GetGoodsPublishCommandDto() *GoodsPublishCommandDto {
	return poolGoodsPublishCommandDto.Get().(*GoodsPublishCommandDto)
}

// ReleaseGoodsPublishCommandDto 释放GoodsPublishCommandDto
func ReleaseGoodsPublishCommandDto(v *GoodsPublishCommandDto) {
	v.GoodsList = v.GoodsList[:0]
	v.ExternalBatchId = ""
	poolGoodsPublishCommandDto.Put(v)
}
