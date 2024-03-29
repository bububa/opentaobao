package damai

import (
	"sync"
)

// PushTicketItemPushOpenParam 结构体
type PushTicketItemPushOpenParam struct {
	// 票品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 票品描述
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 票品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 票品价格
	ItemPrice int64 `json:"item_price,omitempty" xml:"item_price,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 商户id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
}

var poolPushTicketItemPushOpenParam = sync.Pool{
	New: func() any {
		return new(PushTicketItemPushOpenParam)
	},
}

// GetPushTicketItemPushOpenParam() 从对象池中获取PushTicketItemPushOpenParam
func GetPushTicketItemPushOpenParam() *PushTicketItemPushOpenParam {
	return poolPushTicketItemPushOpenParam.Get().(*PushTicketItemPushOpenParam)
}

// ReleasePushTicketItemPushOpenParam 释放PushTicketItemPushOpenParam
func ReleasePushTicketItemPushOpenParam(v *PushTicketItemPushOpenParam) {
	v.ItemName = ""
	v.PushTime = ""
	v.SupplierSecret = ""
	v.Remark = ""
	v.ItemId = 0
	v.ItemPrice = 0
	v.PerformId = 0
	v.SystemId = 0
	poolPushTicketItemPushOpenParam.Put(v)
}
