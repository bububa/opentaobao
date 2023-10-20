package tbk

import (
	"sync"
)

// TaobaoTbkScMembergroupOptionalMapData 结构体
type TaobaoTbkScMembergroupOptionalMapData struct {
	// 组内的成员ID
	TbNumIds string `json:"tb_num_ids,omitempty" xml:"tb_num_ids,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 组ID
	MemberGroupId int64 `json:"member_group_id,omitempty" xml:"member_group_id,omitempty"`
}

var poolTaobaoTbkScMembergroupOptionalMapData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScMembergroupOptionalMapData)
	},
}

// GetTaobaoTbkScMembergroupOptionalMapData() 从对象池中获取TaobaoTbkScMembergroupOptionalMapData
func GetTaobaoTbkScMembergroupOptionalMapData() *TaobaoTbkScMembergroupOptionalMapData {
	return poolTaobaoTbkScMembergroupOptionalMapData.Get().(*TaobaoTbkScMembergroupOptionalMapData)
}

// ReleaseTaobaoTbkScMembergroupOptionalMapData 释放TaobaoTbkScMembergroupOptionalMapData
func ReleaseTaobaoTbkScMembergroupOptionalMapData(v *TaobaoTbkScMembergroupOptionalMapData) {
	v.TbNumIds = ""
	v.CreateTime = ""
	v.UpdateTime = ""
	v.MemberGroupId = 0
	poolTaobaoTbkScMembergroupOptionalMapData.Put(v)
}
