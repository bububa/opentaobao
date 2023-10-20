package product

import (
	"sync"
)

// GoodsBatchResultDto 结构体
type GoodsBatchResultDto struct {
	// 商品下架批次ID
	BatchId int64 `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
}

var poolGoodsBatchResultDto = sync.Pool{
	New: func() any {
		return new(GoodsBatchResultDto)
	},
}

// GetGoodsBatchResultDto() 从对象池中获取GoodsBatchResultDto
func GetGoodsBatchResultDto() *GoodsBatchResultDto {
	return poolGoodsBatchResultDto.Get().(*GoodsBatchResultDto)
}

// ReleaseGoodsBatchResultDto 释放GoodsBatchResultDto
func ReleaseGoodsBatchResultDto(v *GoodsBatchResultDto) {
	v.BatchId = 0
	poolGoodsBatchResultDto.Put(v)
}
