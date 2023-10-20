package product

import (
	"sync"
)

// BatchGoodsStatusResultDto 结构体
type BatchGoodsStatusResultDto struct {
	// 商品状态列表
	GoodsStatusList []GoodsStatusDto `json:"goods_status_list,omitempty" xml:"goods_status_list>goods_status_dto,omitempty"`
}

var poolBatchGoodsStatusResultDto = sync.Pool{
	New: func() any {
		return new(BatchGoodsStatusResultDto)
	},
}

// GetBatchGoodsStatusResultDto() 从对象池中获取BatchGoodsStatusResultDto
func GetBatchGoodsStatusResultDto() *BatchGoodsStatusResultDto {
	return poolBatchGoodsStatusResultDto.Get().(*BatchGoodsStatusResultDto)
}

// ReleaseBatchGoodsStatusResultDto 释放BatchGoodsStatusResultDto
func ReleaseBatchGoodsStatusResultDto(v *BatchGoodsStatusResultDto) {
	v.GoodsStatusList = v.GoodsStatusList[:0]
	poolBatchGoodsStatusResultDto.Put(v)
}
