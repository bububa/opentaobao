package normalvisa

import (
	"sync"
)

// LogisticsCompanyInfo 结构体
type LogisticsCompanyInfo struct {
	// 物流公司列表
	LogisticsCompanyUnitTopVOList []LogisticsCompanyUnitTopVo `json:"logistics_company_unit_top_v_o_list,omitempty" xml:"logistics_company_unit_top_v_o_list>logistics_company_unit_top_vo,omitempty"`
}

var poolLogisticsCompanyInfo = sync.Pool{
	New: func() any {
		return new(LogisticsCompanyInfo)
	},
}

// GetLogisticsCompanyInfo() 从对象池中获取LogisticsCompanyInfo
func GetLogisticsCompanyInfo() *LogisticsCompanyInfo {
	return poolLogisticsCompanyInfo.Get().(*LogisticsCompanyInfo)
}

// ReleaseLogisticsCompanyInfo 释放LogisticsCompanyInfo
func ReleaseLogisticsCompanyInfo(v *LogisticsCompanyInfo) {
	v.LogisticsCompanyUnitTopVOList = v.LogisticsCompanyUnitTopVOList[:0]
	poolLogisticsCompanyInfo.Put(v)
}
