package product

import (
	"sync"
)

// GoodsOffSaleCommandDto 结构体
type GoodsOffSaleCommandDto struct {
	// 批量下架商品id集合
	ExternalGoodsIdList []ExternalGoodsIdDto `json:"external_goods_id_list,omitempty" xml:"external_goods_id_list>external_goods_id_dto,omitempty"`
	// 外部批次ID，用于幂等
	ExternalBatchId string `json:"external_batch_id,omitempty" xml:"external_batch_id,omitempty"`
}

var poolGoodsOffSaleCommandDto = sync.Pool{
	New: func() any {
		return new(GoodsOffSaleCommandDto)
	},
}

// GetGoodsOffSaleCommandDto() 从对象池中获取GoodsOffSaleCommandDto
func GetGoodsOffSaleCommandDto() *GoodsOffSaleCommandDto {
	return poolGoodsOffSaleCommandDto.Get().(*GoodsOffSaleCommandDto)
}

// ReleaseGoodsOffSaleCommandDto 释放GoodsOffSaleCommandDto
func ReleaseGoodsOffSaleCommandDto(v *GoodsOffSaleCommandDto) {
	v.ExternalGoodsIdList = v.ExternalGoodsIdList[:0]
	v.ExternalBatchId = ""
	poolGoodsOffSaleCommandDto.Put(v)
}
