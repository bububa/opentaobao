package usergrowth

import (
	"sync"
)

// RegionData 结构体
type RegionData struct {
	// 坐标字符串
	CoordinateStr string `json:"coordinate_str,omitempty" xml:"coordinate_str,omitempty"`
	// 文本
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 宽
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 高
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// labelId
	LabelId string `json:"label_id,omitempty" xml:"label_id,omitempty"`
	// 序号
	Ranking int64 `json:"ranking,omitempty" xml:"ranking,omitempty"`
}

var poolRegionData = sync.Pool{
	New: func() any {
		return new(RegionData)
	},
}

// GetRegionData() 从对象池中获取RegionData
func GetRegionData() *RegionData {
	return poolRegionData.Get().(*RegionData)
}

// ReleaseRegionData 释放RegionData
func ReleaseRegionData(v *RegionData) {
	v.CoordinateStr = ""
	v.Text = ""
	v.Width = ""
	v.Height = ""
	v.LabelId = ""
	v.Ranking = 0
	poolRegionData.Put(v)
}
