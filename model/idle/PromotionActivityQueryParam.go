package idle

import (
	"sync"
)

// PromotionActivityQueryParam 结构体
type PromotionActivityQueryParam struct {
	// 业务唯一标识
	UniqueKey string `json:"unique_key,omitempty" xml:"unique_key,omitempty"`
	// 预约时间戳，单位秒
	BookTime string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// 素材ID
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
	// 任务ID
	TaskId int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

var poolPromotionActivityQueryParam = sync.Pool{
	New: func() any {
		return new(PromotionActivityQueryParam)
	},
}

// GetPromotionActivityQueryParam() 从对象池中获取PromotionActivityQueryParam
func GetPromotionActivityQueryParam() *PromotionActivityQueryParam {
	return poolPromotionActivityQueryParam.Get().(*PromotionActivityQueryParam)
}

// ReleasePromotionActivityQueryParam 释放PromotionActivityQueryParam
func ReleasePromotionActivityQueryParam(v *PromotionActivityQueryParam) {
	v.UniqueKey = ""
	v.BookTime = ""
	v.MaterialId = 0
	v.TaskId = 0
	poolPromotionActivityQueryParam.Put(v)
}
