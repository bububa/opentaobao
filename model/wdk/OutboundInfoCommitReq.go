package wdk

import (
	"sync"
)

// OutboundInfoCommitReq 结构体
type OutboundInfoCommitReq struct {
	// 商品列表
	OutboundItemInfos []OutboundItemInfo `json:"outbound_item_infos,omitempty" xml:"outbound_item_infos>outbound_item_info,omitempty"`
	// 预计到货时间
	EstimatedArrivalAt string `json:"estimated_arrival_at,omitempty" xml:"estimated_arrival_at,omitempty"`
	// 出货时间
	OutboundAt string `json:"outbound_at,omitempty" xml:"outbound_at,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 供应商编码
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// asn单号
	AsnOrderNo string `json:"asn_order_no,omitempty" xml:"asn_order_no,omitempty"`
	// 商家编码，已经废弃
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 收货类型(信任收货、非信任收货)
	TrustedInbound bool `json:"trusted_inbound,omitempty" xml:"trusted_inbound,omitempty"`
}

var poolOutboundInfoCommitReq = sync.Pool{
	New: func() any {
		return new(OutboundInfoCommitReq)
	},
}

// GetOutboundInfoCommitReq() 从对象池中获取OutboundInfoCommitReq
func GetOutboundInfoCommitReq() *OutboundInfoCommitReq {
	return poolOutboundInfoCommitReq.Get().(*OutboundInfoCommitReq)
}

// ReleaseOutboundInfoCommitReq 释放OutboundInfoCommitReq
func ReleaseOutboundInfoCommitReq(v *OutboundInfoCommitReq) {
	v.OutboundItemInfos = v.OutboundItemInfos[:0]
	v.EstimatedArrivalAt = ""
	v.OutboundAt = ""
	v.SupplierName = ""
	v.SupplierCode = ""
	v.AsnOrderNo = ""
	v.MerchantCode = ""
	v.TrustedInbound = false
	poolOutboundInfoCommitReq.Put(v)
}
