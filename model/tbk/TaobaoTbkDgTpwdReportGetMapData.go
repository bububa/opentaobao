package tbk

// TaobaotbkdgtpwdreportgetMapData 结构体
type TaobaotbkdgtpwdreportgetMapData struct {
	// 截止查询时刻近1小时回流pv
	Hourpv int64 `json:"hour_pv,omitempty" xml:"hour_pv,omitempty"`
	// 截止查询时刻近1小时回流uv
	Houruv int64 `json:"hour_uv,omitempty" xml:"hour_uv,omitempty"`
	// 今日截止查询时刻累计uv
	Uv int64 `json:"uv,omitempty" xml:"uv,omitempty"`
	// 今日截止查询时刻累计pv
	Pv int64 `json:"pv,omitempty" xml:"pv,omitempty"`
}
