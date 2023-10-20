package product

import (
	"sync"
)

// GoodsBatchTaskQueryDto 结构体
type GoodsBatchTaskQueryDto struct {
	// 任务批次ID
	BatchId int64 `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
}

var poolGoodsBatchTaskQueryDto = sync.Pool{
	New: func() any {
		return new(GoodsBatchTaskQueryDto)
	},
}

// GetGoodsBatchTaskQueryDto() 从对象池中获取GoodsBatchTaskQueryDto
func GetGoodsBatchTaskQueryDto() *GoodsBatchTaskQueryDto {
	return poolGoodsBatchTaskQueryDto.Get().(*GoodsBatchTaskQueryDto)
}

// ReleaseGoodsBatchTaskQueryDto 释放GoodsBatchTaskQueryDto
func ReleaseGoodsBatchTaskQueryDto(v *GoodsBatchTaskQueryDto) {
	v.BatchId = 0
	poolGoodsBatchTaskQueryDto.Put(v)
}
