package tbk

import (
	"sync"
)

// TaobaoTbkDgTpwdReportGetMapData 结构体
type TaobaoTbkDgTpwdReportGetMapData struct {
	// 截止查询时刻近1小时回流pv
	HourPv int64 `json:"hour_pv,omitempty" xml:"hour_pv,omitempty"`
	// 截止查询时刻近1小时回流uv
	HourUv int64 `json:"hour_uv,omitempty" xml:"hour_uv,omitempty"`
	// 今日截止查询时刻累计uv
	Uv int64 `json:"uv,omitempty" xml:"uv,omitempty"`
	// 今日截止查询时刻累计pv
	Pv int64 `json:"pv,omitempty" xml:"pv,omitempty"`
}

var poolTaobaoTbkDgTpwdReportGetMapData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgTpwdReportGetMapData)
	},
}

// GetTaobaoTbkDgTpwdReportGetMapData() 从对象池中获取TaobaoTbkDgTpwdReportGetMapData
func GetTaobaoTbkDgTpwdReportGetMapData() *TaobaoTbkDgTpwdReportGetMapData {
	return poolTaobaoTbkDgTpwdReportGetMapData.Get().(*TaobaoTbkDgTpwdReportGetMapData)
}

// ReleaseTaobaoTbkDgTpwdReportGetMapData 释放TaobaoTbkDgTpwdReportGetMapData
func ReleaseTaobaoTbkDgTpwdReportGetMapData(v *TaobaoTbkDgTpwdReportGetMapData) {
	v.HourPv = 0
	v.HourUv = 0
	v.Uv = 0
	v.Pv = 0
	poolTaobaoTbkDgTpwdReportGetMapData.Put(v)
}
