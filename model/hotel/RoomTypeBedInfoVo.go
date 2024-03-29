package hotel

import (
	"sync"
)

// RoomTypeBedInfoVo 结构体
type RoomTypeBedInfoVo struct {
	// 或关系床型集合
	BedInfoGroups []BedInfoGroupVo `json:"bed_info_groups,omitempty" xml:"bed_info_groups>bed_info_group_vo,omitempty"`
	// 分类，大类，用于搜索的筛选项
	Classifications []string `json:"classifications,omitempty" xml:"classifications>string,omitempty"`
	// 简短描述，用于详情页报价前面的床型展示。
	BriefDesc string `json:"brief_desc,omitempty" xml:"brief_desc,omitempty"`
	// 分类描述
	ClassificationDesc string `json:"classification_desc,omitempty" xml:"classification_desc,omitempty"`
	// 描述，长描述，描述为最详细的内容，用于详情页房型详情Dialog，下单页，订单页展示
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 模糊描述，最精简的描述信息，大床/双床/单人床/多床/床位简单描述，较长描述省略床宽
	FuzzyDesc string `json:"fuzzy_desc,omitempty" xml:"fuzzy_desc,omitempty"`
	// 简单描述，较长描述省略床宽，但依然会描述具体的床型信息，用于详情页标准房型床型展示
	SimpleDesc string `json:"simple_desc,omitempty" xml:"simple_desc,omitempty"`
}

var poolRoomTypeBedInfoVo = sync.Pool{
	New: func() any {
		return new(RoomTypeBedInfoVo)
	},
}

// GetRoomTypeBedInfoVo() 从对象池中获取RoomTypeBedInfoVo
func GetRoomTypeBedInfoVo() *RoomTypeBedInfoVo {
	return poolRoomTypeBedInfoVo.Get().(*RoomTypeBedInfoVo)
}

// ReleaseRoomTypeBedInfoVo 释放RoomTypeBedInfoVo
func ReleaseRoomTypeBedInfoVo(v *RoomTypeBedInfoVo) {
	v.BedInfoGroups = v.BedInfoGroups[:0]
	v.Classifications = v.Classifications[:0]
	v.BriefDesc = ""
	v.ClassificationDesc = ""
	v.Desc = ""
	v.FuzzyDesc = ""
	v.SimpleDesc = ""
	poolRoomTypeBedInfoVo.Put(v)
}
