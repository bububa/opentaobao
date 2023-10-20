package bus

import (
	"sync"
)

// OfflineRefundTicketPriceRq 结构体
type OfflineRefundTicketPriceRq struct {
	// 退票信息集合
	ListRefundInfos []SingleRefundInfo `json:"list_refund_infos,omitempty" xml:"list_refund_infos>single_refund_info,omitempty"`
	// top平台用户唯一识别
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 代理商店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
}

var poolOfflineRefundTicketPriceRq = sync.Pool{
	New: func() any {
		return new(OfflineRefundTicketPriceRq)
	},
}

// GetOfflineRefundTicketPriceRq() 从对象池中获取OfflineRefundTicketPriceRq
func GetOfflineRefundTicketPriceRq() *OfflineRefundTicketPriceRq {
	return poolOfflineRefundTicketPriceRq.Get().(*OfflineRefundTicketPriceRq)
}

// ReleaseOfflineRefundTicketPriceRq 释放OfflineRefundTicketPriceRq
func ReleaseOfflineRefundTicketPriceRq(v *OfflineRefundTicketPriceRq) {
	v.ListRefundInfos = v.ListRefundInfos[:0]
	v.AppKey = ""
	v.AgentId = 0
	poolOfflineRefundTicketPriceRq.Put(v)
}
