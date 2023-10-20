package product

import (
	"sync"
)

// GoodsBatchTaskResultDto 结构体
type GoodsBatchTaskResultDto struct {
	// 商品批次子任务对象集合
	GoodsBatchSubTaskList []GoodsBatchSubTask `json:"goods_batch_sub_task_list,omitempty" xml:"goods_batch_sub_task_list>goods_batch_sub_task,omitempty"`
	// 任务状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolGoodsBatchTaskResultDto = sync.Pool{
	New: func() any {
		return new(GoodsBatchTaskResultDto)
	},
}

// GetGoodsBatchTaskResultDto() 从对象池中获取GoodsBatchTaskResultDto
func GetGoodsBatchTaskResultDto() *GoodsBatchTaskResultDto {
	return poolGoodsBatchTaskResultDto.Get().(*GoodsBatchTaskResultDto)
}

// ReleaseGoodsBatchTaskResultDto 释放GoodsBatchTaskResultDto
func ReleaseGoodsBatchTaskResultDto(v *GoodsBatchTaskResultDto) {
	v.GoodsBatchSubTaskList = v.GoodsBatchSubTaskList[:0]
	v.Status = 0
	poolGoodsBatchTaskResultDto.Put(v)
}
