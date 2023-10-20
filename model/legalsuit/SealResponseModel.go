package legalsuit

import (
	"sync"
)

// SealResponseModel 结构体
type SealResponseModel struct {
	// 批次id
	BatchId string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 案件id
	SuitId int64 `json:"suit_id,omitempty" xml:"suit_id,omitempty"`
	// 委托id
	EntrustId int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
	// 任务id
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

var poolSealResponseModel = sync.Pool{
	New: func() any {
		return new(SealResponseModel)
	},
}

// GetSealResponseModel() 从对象池中获取SealResponseModel
func GetSealResponseModel() *SealResponseModel {
	return poolSealResponseModel.Get().(*SealResponseModel)
}

// ReleaseSealResponseModel 释放SealResponseModel
func ReleaseSealResponseModel(v *SealResponseModel) {
	v.BatchId = ""
	v.SuitId = 0
	v.EntrustId = 0
	v.TaskId = 0
	poolSealResponseModel.Put(v)
}
