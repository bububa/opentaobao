package perfect

import (
	"sync"
)

// PerfectPerformanceItemPublishResp 结构体
type PerfectPerformanceItemPublishResp struct {
	// 商品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
}

var poolPerfectPerformanceItemPublishResp = sync.Pool{
	New: func() any {
		return new(PerfectPerformanceItemPublishResp)
	},
}

// GetPerfectPerformanceItemPublishResp() 从对象池中获取PerfectPerformanceItemPublishResp
func GetPerfectPerformanceItemPublishResp() *PerfectPerformanceItemPublishResp {
	return poolPerfectPerformanceItemPublishResp.Get().(*PerfectPerformanceItemPublishResp)
}

// ReleasePerfectPerformanceItemPublishResp 释放PerfectPerformanceItemPublishResp
func ReleasePerfectPerformanceItemPublishResp(v *PerfectPerformanceItemPublishResp) {
	v.ItemCode = ""
	poolPerfectPerformanceItemPublishResp.Put(v)
}
