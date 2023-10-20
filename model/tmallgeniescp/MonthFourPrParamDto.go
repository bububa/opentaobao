package tmallgeniescp

import (
	"sync"
)

// MonthFourPrParamDto 结构体
type MonthFourPrParamDto struct {
	// 关键值日期
	KeyFigureDate string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
	// PR数量
	PrQty string `json:"pr_qty,omitempty" xml:"pr_qty,omitempty"`
	// 产品线
	ProductionLine string `json:"production_line,omitempty" xml:"production_line,omitempty"`
	// 扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

var poolMonthFourPrParamDto = sync.Pool{
	New: func() any {
		return new(MonthFourPrParamDto)
	},
}

// GetMonthFourPrParamDto() 从对象池中获取MonthFourPrParamDto
func GetMonthFourPrParamDto() *MonthFourPrParamDto {
	return poolMonthFourPrParamDto.Get().(*MonthFourPrParamDto)
}

// ReleaseMonthFourPrParamDto 释放MonthFourPrParamDto
func ReleaseMonthFourPrParamDto(v *MonthFourPrParamDto) {
	v.KeyFigureDate = ""
	v.PrQty = ""
	v.ProductionLine = ""
	v.ExtendJson = ""
	v.Tenant = ""
	poolMonthFourPrParamDto.Put(v)
}
