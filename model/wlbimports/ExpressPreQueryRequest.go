package wlbimports

import (
	"sync"
)

// ExpressPreQueryRequest 结构体
type ExpressPreQueryRequest struct {
	// packages
	Packages []HandoverPackageInfo `json:"packages,omitempty" xml:"packages>handover_package_info,omitempty"`
	// 期望揽收日期以及时间
	PlannedShippingDateAndTime string `json:"planned_shipping_date_and_time,omitempty" xml:"planned_shipping_date_and_time,omitempty"`
	// scItemInfo
	ScitemInfo string `json:"scitem_info,omitempty" xml:"scitem_info,omitempty"`
	// 揽收cp资源code
	ReceiveCpCode string `json:"receive_cp_code,omitempty" xml:"receive_cp_code,omitempty"`
	// senderInfo
	SenderInfo *ContactInfoRequest `json:"sender_info,omitempty" xml:"sender_info,omitempty"`
	// receiverInfo
	ReceiverInfo *ContactInfoRequest `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}

var poolExpressPreQueryRequest = sync.Pool{
	New: func() any {
		return new(ExpressPreQueryRequest)
	},
}

// GetExpressPreQueryRequest() 从对象池中获取ExpressPreQueryRequest
func GetExpressPreQueryRequest() *ExpressPreQueryRequest {
	return poolExpressPreQueryRequest.Get().(*ExpressPreQueryRequest)
}

// ReleaseExpressPreQueryRequest 释放ExpressPreQueryRequest
func ReleaseExpressPreQueryRequest(v *ExpressPreQueryRequest) {
	v.Packages = v.Packages[:0]
	v.PlannedShippingDateAndTime = ""
	v.ScitemInfo = ""
	v.ReceiveCpCode = ""
	v.SenderInfo = nil
	v.ReceiverInfo = nil
	v.SellerId = 0
	poolExpressPreQueryRequest.Put(v)
}
