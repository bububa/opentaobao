package kbalgo

import (
	"sync"
)

// PoiToBaiduData 结构体
type PoiToBaiduData struct {
	// poiid
	PoiId string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	// 数据日期
	Dt string `json:"dt,omitempty" xml:"dt,omitempty"`
	// poi明细
	Content *Content `json:"content,omitempty" xml:"content,omitempty"`
}

var poolPoiToBaiduData = sync.Pool{
	New: func() any {
		return new(PoiToBaiduData)
	},
}

// GetPoiToBaiduData() 从对象池中获取PoiToBaiduData
func GetPoiToBaiduData() *PoiToBaiduData {
	return poolPoiToBaiduData.Get().(*PoiToBaiduData)
}

// ReleasePoiToBaiduData 释放PoiToBaiduData
func ReleasePoiToBaiduData(v *PoiToBaiduData) {
	v.PoiId = ""
	v.Dt = ""
	v.Content = nil
	poolPoiToBaiduData.Put(v)
}
