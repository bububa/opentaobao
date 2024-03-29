package mydata

import (
	"sync"
)

// CompanyIndicators 结构体
type CompanyIndicators struct {
	// 点击率
	ClkRate string `json:"clk_rate,omitempty" xml:"clk_rate,omitempty"`
	// 最近30天询盘一次回复率
	Reply string `json:"reply,omitempty" xml:"reply,omitempty"`
	// 点击
	Clk int64 `json:"clk,omitempty" xml:"clk,omitempty"`
	// 反馈（询盘）
	Fb int64 `json:"fb,omitempty" xml:"fb,omitempty"`
	// 曝光
	Imps int64 `json:"imps,omitempty" xml:"imps,omitempty"`
	// 访客
	Visitor int64 `json:"visitor,omitempty" xml:"visitor,omitempty"`
}

var poolCompanyIndicators = sync.Pool{
	New: func() any {
		return new(CompanyIndicators)
	},
}

// GetCompanyIndicators() 从对象池中获取CompanyIndicators
func GetCompanyIndicators() *CompanyIndicators {
	return poolCompanyIndicators.Get().(*CompanyIndicators)
}

// ReleaseCompanyIndicators 释放CompanyIndicators
func ReleaseCompanyIndicators(v *CompanyIndicators) {
	v.ClkRate = ""
	v.Reply = ""
	v.Clk = 0
	v.Fb = 0
	v.Imps = 0
	v.Visitor = 0
	poolCompanyIndicators.Put(v)
}
