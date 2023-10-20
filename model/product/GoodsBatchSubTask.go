package product

import (
	"sync"
)

// GoodsBatchSubTask 结构体
type GoodsBatchSubTask struct {
	// 子任务状态产生原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 外部商品ID，用于标识外部系统每次提交过来的商品
	ExternalGoodsId string `json:"external_goods_id,omitempty" xml:"external_goods_id,omitempty"`
	// 子任务状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 子任务id
	SubBatchId int64 `json:"sub_batch_id,omitempty" xml:"sub_batch_id,omitempty"`
	// 交易猫商品ID，如果商品发布失败则为空
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}

var poolGoodsBatchSubTask = sync.Pool{
	New: func() any {
		return new(GoodsBatchSubTask)
	},
}

// GetGoodsBatchSubTask() 从对象池中获取GoodsBatchSubTask
func GetGoodsBatchSubTask() *GoodsBatchSubTask {
	return poolGoodsBatchSubTask.Get().(*GoodsBatchSubTask)
}

// ReleaseGoodsBatchSubTask 释放GoodsBatchSubTask
func ReleaseGoodsBatchSubTask(v *GoodsBatchSubTask) {
	v.Reason = ""
	v.ExternalGoodsId = ""
	v.Status = 0
	v.SubBatchId = 0
	v.GoodsId = 0
	poolGoodsBatchSubTask.Put(v)
}
