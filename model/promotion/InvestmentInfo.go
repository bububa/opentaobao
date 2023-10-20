package promotion

import (
	"sync"
)

// InvestmentInfo 结构体
type InvestmentInfo struct {
	// 出资人
	Investor string `json:"investor,omitempty" xml:"investor,omitempty"`
	// 出资比例 2000 = 20%
	InvestorRatio int64 `json:"investor_ratio,omitempty" xml:"investor_ratio,omitempty"`
}

var poolInvestmentInfo = sync.Pool{
	New: func() any {
		return new(InvestmentInfo)
	},
}

// GetInvestmentInfo() 从对象池中获取InvestmentInfo
func GetInvestmentInfo() *InvestmentInfo {
	return poolInvestmentInfo.Get().(*InvestmentInfo)
}

// ReleaseInvestmentInfo 释放InvestmentInfo
func ReleaseInvestmentInfo(v *InvestmentInfo) {
	v.Investor = ""
	v.InvestorRatio = 0
	poolInvestmentInfo.Put(v)
}
