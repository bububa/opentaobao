package flight

import (
	"sync"
)

// PushAuxProductsRq 结构体
type PushAuxProductsRq struct {
	// 廉航辅营产品
	ProductItems []AuxProductItemApiBean `json:"product_items,omitempty" xml:"product_items>aux_product_item_api_bean,omitempty"`
	// 接口身份标识用户名（渠道唯一标识）
	Cid string `json:"cid,omitempty" xml:"cid,omitempty"`
	// 渠道id
	ChannelId int64 `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 代理商ID
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
}

var poolPushAuxProductsRq = sync.Pool{
	New: func() any {
		return new(PushAuxProductsRq)
	},
}

// GetPushAuxProductsRq() 从对象池中获取PushAuxProductsRq
func GetPushAuxProductsRq() *PushAuxProductsRq {
	return poolPushAuxProductsRq.Get().(*PushAuxProductsRq)
}

// ReleasePushAuxProductsRq 释放PushAuxProductsRq
func ReleasePushAuxProductsRq(v *PushAuxProductsRq) {
	v.ProductItems = v.ProductItems[:0]
	v.Cid = ""
	v.ChannelId = 0
	v.AgentId = 0
	poolPushAuxProductsRq.Put(v)
}
