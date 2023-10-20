package product

import (
	"sync"
)

// GoodsOnSaleCommandDto 结构体
type GoodsOnSaleCommandDto struct {
	// 批量上架商品id集合
	ExternalGoodsIdList []ExternalGoodsIdDto `json:"external_goods_id_list,omitempty" xml:"external_goods_id_list>external_goods_id_dto,omitempty"`
	// 外部批次ID，用于幂等
	ExternalBatchId string `json:"external_batch_id,omitempty" xml:"external_batch_id,omitempty"`
}

var poolGoodsOnSaleCommandDto = sync.Pool{
	New: func() any {
		return new(GoodsOnSaleCommandDto)
	},
}

// GetGoodsOnSaleCommandDto() 从对象池中获取GoodsOnSaleCommandDto
func GetGoodsOnSaleCommandDto() *GoodsOnSaleCommandDto {
	return poolGoodsOnSaleCommandDto.Get().(*GoodsOnSaleCommandDto)
}

// ReleaseGoodsOnSaleCommandDto 释放GoodsOnSaleCommandDto
func ReleaseGoodsOnSaleCommandDto(v *GoodsOnSaleCommandDto) {
	v.ExternalGoodsIdList = v.ExternalGoodsIdList[:0]
	v.ExternalBatchId = ""
	poolGoodsOnSaleCommandDto.Put(v)
}
