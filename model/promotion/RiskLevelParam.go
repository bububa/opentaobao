package promotion

import (
	"sync"
)

// RiskLevelParam 结构体
type RiskLevelParam struct {
	// 风险等级(可选值:higher-risk,middle-risk,low-risk)
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 风险等级比例左边值
	LeftRange int64 `json:"left_range,omitempty" xml:"left_range,omitempty"`
	// 风险等级比例右边值
	RightRange int64 `json:"right_range,omitempty" xml:"right_range,omitempty"`
}

var poolRiskLevelParam = sync.Pool{
	New: func() any {
		return new(RiskLevelParam)
	},
}

// GetRiskLevelParam() 从对象池中获取RiskLevelParam
func GetRiskLevelParam() *RiskLevelParam {
	return poolRiskLevelParam.Get().(*RiskLevelParam)
}

// ReleaseRiskLevelParam 释放RiskLevelParam
func ReleaseRiskLevelParam(v *RiskLevelParam) {
	v.Key = ""
	v.LeftRange = 0
	v.RightRange = 0
	poolRiskLevelParam.Put(v)
}
