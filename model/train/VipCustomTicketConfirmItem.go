package train

import (
	"sync"
)

// VipCustomTicketConfirmItem 结构体
type VipCustomTicketConfirmItem struct {
	// 定制信息
	CustomType string `json:"custom_type,omitempty" xml:"custom_type,omitempty"`
}

var poolVipCustomTicketConfirmItem = sync.Pool{
	New: func() any {
		return new(VipCustomTicketConfirmItem)
	},
}

// GetVipCustomTicketConfirmItem() 从对象池中获取VipCustomTicketConfirmItem
func GetVipCustomTicketConfirmItem() *VipCustomTicketConfirmItem {
	return poolVipCustomTicketConfirmItem.Get().(*VipCustomTicketConfirmItem)
}

// ReleaseVipCustomTicketConfirmItem 释放VipCustomTicketConfirmItem
func ReleaseVipCustomTicketConfirmItem(v *VipCustomTicketConfirmItem) {
	v.CustomType = ""
	poolVipCustomTicketConfirmItem.Put(v)
}
