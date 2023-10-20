package antifraud

import (
	"sync"
)

// CollinadataQueryResult 结构体
type CollinadataQueryResult struct {
	// 积分信息.千分制
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// pass,review,reject
	RiskLevel string `json:"risk_level,omitempty" xml:"risk_level,omitempty"`
	// 字符串格式, 关于score生成的描述信息. 本字段可能为空.
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
}

var poolCollinadataQueryResult = sync.Pool{
	New: func() any {
		return new(CollinadataQueryResult)
	},
}

// GetCollinadataQueryResult() 从对象池中获取CollinadataQueryResult
func GetCollinadataQueryResult() *CollinadataQueryResult {
	return poolCollinadataQueryResult.Get().(*CollinadataQueryResult)
}

// ReleaseCollinadataQueryResult 释放CollinadataQueryResult
func ReleaseCollinadataQueryResult(v *CollinadataQueryResult) {
	v.Score = ""
	v.RiskLevel = ""
	v.Detail = ""
	poolCollinadataQueryResult.Put(v)
}
