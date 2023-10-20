package tmallcar

import (
	"sync"
)

// XCarSysLineDto 结构体
type XCarSysLineDto struct {
	// 最高降幅
	CarLineMaxDecline string `json:"car_line_max_decline,omitempty" xml:"car_line_max_decline,omitempty"`
	// 车系保养排名
	LineRank string `json:"line_rank,omitempty" xml:"line_rank,omitempty"`
	// 本地参考价范围
	LocalRefPriceRange string `json:"local_ref_price_range,omitempty" xml:"local_ref_price_range,omitempty"`
	// 厂商指导价范围
	ManuGuiPriceRange string `json:"manu_gui_price_range,omitempty" xml:"manu_gui_price_range,omitempty"`
	// 车系封面图
	Pic string `json:"pic,omitempty" xml:"pic,omitempty"`
	// 年均养护
	YearCuring string `json:"year_curing,omitempty" xml:"year_curing,omitempty"`
	// 品牌属性id
	BrandPid int64 `json:"brand_pid,omitempty" xml:"brand_pid,omitempty"`
	// 品牌值id
	BrandVid int64 `json:"brand_vid,omitempty" xml:"brand_vid,omitempty"`
	// 车系属性id
	LinePid int64 `json:"line_pid,omitempty" xml:"line_pid,omitempty"`
	// 车系值id
	LineVid int64 `json:"line_vid,omitempty" xml:"line_vid,omitempty"`
	// 状态0.无效 1有效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolXCarSysLineDto = sync.Pool{
	New: func() any {
		return new(XCarSysLineDto)
	},
}

// GetXCarSysLineDto() 从对象池中获取XCarSysLineDto
func GetXCarSysLineDto() *XCarSysLineDto {
	return poolXCarSysLineDto.Get().(*XCarSysLineDto)
}

// ReleaseXCarSysLineDto 释放XCarSysLineDto
func ReleaseXCarSysLineDto(v *XCarSysLineDto) {
	v.CarLineMaxDecline = ""
	v.LineRank = ""
	v.LocalRefPriceRange = ""
	v.ManuGuiPriceRange = ""
	v.Pic = ""
	v.YearCuring = ""
	v.BrandPid = 0
	v.BrandVid = 0
	v.LinePid = 0
	v.LineVid = 0
	v.Status = 0
	poolXCarSysLineDto.Put(v)
}
