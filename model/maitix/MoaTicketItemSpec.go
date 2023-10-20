package maitix

import (
	"sync"
)

// MoaTicketItemSpec 结构体
type MoaTicketItemSpec struct {
	// 是否套票 1=套票 0=普通票-必填
	IsPackage int64 `json:"is_package,omitempty" xml:"is_package,omitempty"`
	// 票价，单位分，必填
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 数量， 如果是套票,则是套票的套数(不等于实际票单数量)-必填
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 票品ID,如果是套票就是套票的票品id-必填
	TicketItemId int64 `json:"ticket_item_id,omitempty" xml:"ticket_item_id,omitempty"`
}

var poolMoaTicketItemSpec = sync.Pool{
	New: func() any {
		return new(MoaTicketItemSpec)
	},
}

// GetMoaTicketItemSpec() 从对象池中获取MoaTicketItemSpec
func GetMoaTicketItemSpec() *MoaTicketItemSpec {
	return poolMoaTicketItemSpec.Get().(*MoaTicketItemSpec)
}

// ReleaseMoaTicketItemSpec 释放MoaTicketItemSpec
func ReleaseMoaTicketItemSpec(v *MoaTicketItemSpec) {
	v.IsPackage = 0
	v.Price = 0
	v.Quantity = 0
	v.TicketItemId = 0
	poolMoaTicketItemSpec.Put(v)
}
