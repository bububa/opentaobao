package security

import (
	"sync"
)

// RiskDetailBatch 结构体
type RiskDetailBatch struct {
	// 各个应用的风险列表
	ResultInfos []RiskDetailBatchItem `json:"result_infos,omitempty" xml:"result_infos>risk_detail_batch_item,omitempty"`
	// 任务状态: 1-已完成,2-处理中,3-处理失败,4-处理超时
	TotalStatus int64 `json:"total_status,omitempty" xml:"total_status,omitempty"`
}

var poolRiskDetailBatch = sync.Pool{
	New: func() any {
		return new(RiskDetailBatch)
	},
}

// GetRiskDetailBatch() 从对象池中获取RiskDetailBatch
func GetRiskDetailBatch() *RiskDetailBatch {
	return poolRiskDetailBatch.Get().(*RiskDetailBatch)
}

// ReleaseRiskDetailBatch 释放RiskDetailBatch
func ReleaseRiskDetailBatch(v *RiskDetailBatch) {
	v.ResultInfos = v.ResultInfos[:0]
	v.TotalStatus = 0
	poolRiskDetailBatch.Put(v)
}
