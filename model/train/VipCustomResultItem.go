package train

import (
	"sync"
)

// VipCustomResultItem 结构体
type VipCustomResultItem struct {
	// 定制车厢号
	CarriageCustom []string `json:"carriage_custom,omitempty" xml:"carriage_custom>string,omitempty"`
	// 定制类型
	CustomType string `json:"custom_type,omitempty" xml:"custom_type,omitempty"`
	// 定制坐席号
	SeatCustom string `json:"seat_custom,omitempty" xml:"seat_custom,omitempty"`
	// 定制数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 是否接受非定制,0不接受,1接受
	AcceptNoVipCustom int64 `json:"accept_no_vip_custom,omitempty" xml:"accept_no_vip_custom,omitempty"`
	// 联程订单定制序号
	SegmentIndex int64 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

var poolVipCustomResultItem = sync.Pool{
	New: func() any {
		return new(VipCustomResultItem)
	},
}

// GetVipCustomResultItem() 从对象池中获取VipCustomResultItem
func GetVipCustomResultItem() *VipCustomResultItem {
	return poolVipCustomResultItem.Get().(*VipCustomResultItem)
}

// ReleaseVipCustomResultItem 释放VipCustomResultItem
func ReleaseVipCustomResultItem(v *VipCustomResultItem) {
	v.CarriageCustom = v.CarriageCustom[:0]
	v.CustomType = ""
	v.SeatCustom = ""
	v.Count = 0
	v.AcceptNoVipCustom = 0
	v.SegmentIndex = 0
	poolVipCustomResultItem.Put(v)
}
