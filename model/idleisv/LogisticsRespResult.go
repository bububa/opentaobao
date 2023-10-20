package idleisv

import (
	"sync"
)

// LogisticsRespResult 结构体
type LogisticsRespResult struct {
	// 全部公司列表
	CompanyList []CompanyList `json:"company_list,omitempty" xml:"company_list>company_list,omitempty"`
	// 热门公司列表
	HotCompanyList []HotCompanyList `json:"hot_company_list,omitempty" xml:"hot_company_list>hot_company_list,omitempty"`
	// 快递公司总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolLogisticsRespResult = sync.Pool{
	New: func() any {
		return new(LogisticsRespResult)
	},
}

// GetLogisticsRespResult() 从对象池中获取LogisticsRespResult
func GetLogisticsRespResult() *LogisticsRespResult {
	return poolLogisticsRespResult.Get().(*LogisticsRespResult)
}

// ReleaseLogisticsRespResult 释放LogisticsRespResult
func ReleaseLogisticsRespResult(v *LogisticsRespResult) {
	v.CompanyList = v.CompanyList[:0]
	v.HotCompanyList = v.HotCompanyList[:0]
	v.Total = 0
	poolLogisticsRespResult.Put(v)
}
