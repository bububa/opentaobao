package product

import (
	"sync"
)

// GoodsDeleteCommandDto 结构体
type GoodsDeleteCommandDto struct {
	// 批量删除商品id集合
	ExternalGoodsIdList []ExternalGoodsIdDto `json:"external_goods_id_list,omitempty" xml:"external_goods_id_list>external_goods_id_dto,omitempty"`
	// 外部批次ID，用于幂等
	ExternalBatchId string `json:"external_batch_id,omitempty" xml:"external_batch_id,omitempty"`
}

var poolGoodsDeleteCommandDto = sync.Pool{
	New: func() any {
		return new(GoodsDeleteCommandDto)
	},
}

// GetGoodsDeleteCommandDto() 从对象池中获取GoodsDeleteCommandDto
func GetGoodsDeleteCommandDto() *GoodsDeleteCommandDto {
	return poolGoodsDeleteCommandDto.Get().(*GoodsDeleteCommandDto)
}

// ReleaseGoodsDeleteCommandDto 释放GoodsDeleteCommandDto
func ReleaseGoodsDeleteCommandDto(v *GoodsDeleteCommandDto) {
	v.ExternalGoodsIdList = v.ExternalGoodsIdList[:0]
	v.ExternalBatchId = ""
	poolGoodsDeleteCommandDto.Put(v)
}
