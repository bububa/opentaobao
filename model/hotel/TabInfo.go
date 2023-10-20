package hotel

import (
	"sync"
)

// TabInfo 结构体
type TabInfo struct {
	// tab编码
	TabCode string `json:"tab_code,omitempty" xml:"tab_code,omitempty"`
	// tab下的统计数据
	TabDetail string `json:"tab_detail,omitempty" xml:"tab_detail,omitempty"`
	// tab名称
	TabName string `json:"tab_name,omitempty" xml:"tab_name,omitempty"`
	// tab对应的埋点名
	TabTrack string `json:"tab_track,omitempty" xml:"tab_track,omitempty"`
	// 标签的态度 1好 0中性 -1 差
	Attitude int64 `json:"attitude,omitempty" xml:"attitude,omitempty"`
	// tabId
	TabId int64 `json:"tab_id,omitempty" xml:"tab_id,omitempty"`
	// 标签的类型 0正常 1热词 2房型 3度假
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// tab是否点击
	IsClick bool `json:"is_click,omitempty" xml:"is_click,omitempty"`
}

var poolTabInfo = sync.Pool{
	New: func() any {
		return new(TabInfo)
	},
}

// GetTabInfo() 从对象池中获取TabInfo
func GetTabInfo() *TabInfo {
	return poolTabInfo.Get().(*TabInfo)
}

// ReleaseTabInfo 释放TabInfo
func ReleaseTabInfo(v *TabInfo) {
	v.TabCode = ""
	v.TabDetail = ""
	v.TabName = ""
	v.TabTrack = ""
	v.Attitude = 0
	v.TabId = 0
	v.Type = 0
	v.IsClick = false
	poolTabInfo.Put(v)
}
