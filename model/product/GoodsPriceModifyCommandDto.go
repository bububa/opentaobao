package product

import (
	"sync"
)

// GoodsPriceModifyCommandDto 结构体
type GoodsPriceModifyCommandDto struct {
	// 商品价格传输对象
	GoodsPriceList []GoodsPriceDto `json:"goods_price_list,omitempty" xml:"goods_price_list>goods_price_dto,omitempty"`
	// 外部批次ID，用于幂等
	ExternalBatchId string `json:"external_batch_id,omitempty" xml:"external_batch_id,omitempty"`
}

var poolGoodsPriceModifyCommandDto = sync.Pool{
	New: func() any {
		return new(GoodsPriceModifyCommandDto)
	},
}

// GetGoodsPriceModifyCommandDto() 从对象池中获取GoodsPriceModifyCommandDto
func GetGoodsPriceModifyCommandDto() *GoodsPriceModifyCommandDto {
	return poolGoodsPriceModifyCommandDto.Get().(*GoodsPriceModifyCommandDto)
}

// ReleaseGoodsPriceModifyCommandDto 释放GoodsPriceModifyCommandDto
func ReleaseGoodsPriceModifyCommandDto(v *GoodsPriceModifyCommandDto) {
	v.GoodsPriceList = v.GoodsPriceList[:0]
	v.ExternalBatchId = ""
	poolGoodsPriceModifyCommandDto.Put(v)
}
