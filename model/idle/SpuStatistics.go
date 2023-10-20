package idle

import (
	"sync"
)

// SpuStatistics 结构体
type SpuStatistics struct {
	// 开始时间（预留）
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 场景：3C（3C行业）、BULKS（大件）
	SceneType string `json:"scene_type,omitempty" xml:"scene_type,omitempty"`
	// 预留字段
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 该spu
	Descr string `json:"descr,omitempty" xml:"descr,omitempty"`
	// 结束时间（预留）
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 结构数据，涉及金额的数据单位为分，精确到元
	SpuData string `json:"spu_data,omitempty" xml:"spu_data,omitempty"`
	// 数据状态：1（可用）2（不可用）
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// spuId
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// 版本号（预留）
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolSpuStatistics = sync.Pool{
	New: func() any {
		return new(SpuStatistics)
	},
}

// GetSpuStatistics() 从对象池中获取SpuStatistics
func GetSpuStatistics() *SpuStatistics {
	return poolSpuStatistics.Get().(*SpuStatistics)
}

// ReleaseSpuStatistics 释放SpuStatistics
func ReleaseSpuStatistics(v *SpuStatistics) {
	v.StartTime = ""
	v.SceneType = ""
	v.SkuId = ""
	v.Descr = ""
	v.EndTime = ""
	v.SpuData = ""
	v.Status = 0
	v.SpuId = 0
	v.Version = 0
	poolSpuStatistics.Put(v)
}
