package train

import (
	"sync"
)

// FreeChildrenTicketListRs 结构体
type FreeChildrenTicketListRs struct {
	// 待处理列表
	FreeChildrenTicketNeedDealList []string `json:"free_children_ticket_need_deal_list,omitempty" xml:"free_children_ticket_need_deal_list>string,omitempty"`
}

var poolFreeChildrenTicketListRs = sync.Pool{
	New: func() any {
		return new(FreeChildrenTicketListRs)
	},
}

// GetFreeChildrenTicketListRs() 从对象池中获取FreeChildrenTicketListRs
func GetFreeChildrenTicketListRs() *FreeChildrenTicketListRs {
	return poolFreeChildrenTicketListRs.Get().(*FreeChildrenTicketListRs)
}

// ReleaseFreeChildrenTicketListRs 释放FreeChildrenTicketListRs
func ReleaseFreeChildrenTicketListRs(v *FreeChildrenTicketListRs) {
	v.FreeChildrenTicketNeedDealList = v.FreeChildrenTicketNeedDealList[:0]
	poolFreeChildrenTicketListRs.Put(v)
}
