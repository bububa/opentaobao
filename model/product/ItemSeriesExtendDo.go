package product

import (
	"sync"
)

// ItemSeriesExtendDo 结构体
type ItemSeriesExtendDo struct {
	// 分组信息
	GroupItemList []GroupItem `json:"group_item_list,omitempty" xml:"group_item_list>group_item,omitempty"`
	// 系列玩法
	Mode string `json:"mode,omitempty" xml:"mode,omitempty"`
	// 系列名称
	SeriesName string `json:"series_name,omitempty" xml:"series_name,omitempty"`
	// 系列描述
	SeriesDesc string `json:"series_desc,omitempty" xml:"series_desc,omitempty"`
	// 类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
}

var poolItemSeriesExtendDo = sync.Pool{
	New: func() any {
		return new(ItemSeriesExtendDo)
	},
}

// GetItemSeriesExtendDo() 从对象池中获取ItemSeriesExtendDo
func GetItemSeriesExtendDo() *ItemSeriesExtendDo {
	return poolItemSeriesExtendDo.Get().(*ItemSeriesExtendDo)
}

// ReleaseItemSeriesExtendDo 释放ItemSeriesExtendDo
func ReleaseItemSeriesExtendDo(v *ItemSeriesExtendDo) {
	v.GroupItemList = v.GroupItemList[:0]
	v.Mode = ""
	v.SeriesName = ""
	v.SeriesDesc = ""
	v.CatId = 0
	poolItemSeriesExtendDo.Put(v)
}
