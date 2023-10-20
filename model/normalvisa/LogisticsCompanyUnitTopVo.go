package normalvisa

import (
	"sync"
)

// LogisticsCompanyUnitTopVo 结构体
type LogisticsCompanyUnitTopVo struct {
	// 物流公司名
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 物流公司代码
	CompanyCode string `json:"company_code,omitempty" xml:"company_code,omitempty"`
}

var poolLogisticsCompanyUnitTopVo = sync.Pool{
	New: func() any {
		return new(LogisticsCompanyUnitTopVo)
	},
}

// GetLogisticsCompanyUnitTopVo() 从对象池中获取LogisticsCompanyUnitTopVo
func GetLogisticsCompanyUnitTopVo() *LogisticsCompanyUnitTopVo {
	return poolLogisticsCompanyUnitTopVo.Get().(*LogisticsCompanyUnitTopVo)
}

// ReleaseLogisticsCompanyUnitTopVo 释放LogisticsCompanyUnitTopVo
func ReleaseLogisticsCompanyUnitTopVo(v *LogisticsCompanyUnitTopVo) {
	v.CompanyName = ""
	v.CompanyCode = ""
	poolLogisticsCompanyUnitTopVo.Put(v)
}
