package tbtrade

import (
	"sync"
)

// UDSmartServiceRequest 结构体
type UDSmartServiceRequest struct {
	// 服务商标识，可自定义
	TecAgent string `json:"tec_agent,omitempty" xml:"tec_agent,omitempty"`
	// 点击监测归因click_id，与直达唤端归因clickid_id 二选一必填
	CmClickId string `json:"cm_click_id,omitempty" xml:"cm_click_id,omitempty"`
	// 点击监测归媒体替换的requestid
	CmRequestId string `json:"cm_request_id,omitempty" xml:"cm_request_id,omitempty"`
	// 字节2.0product_id
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 子订单ID，必填
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 必填；12-支付，28-退货，25-落地页点击（仅微博）
	EventType string `json:"event_type,omitempty" xml:"event_type,omitempty"`
	// 直达唤端归因clickid_id，与点击监测归因click_id 二选一必填
	CkClickId string `json:"ck_click_id,omitempty" xml:"ck_click_id,omitempty"`
	// app_access_token,流量通Pro需回传群峰token
	AppAccessToken string `json:"app_access_token,omitempty" xml:"app_access_token,omitempty"`
	// 支付金额（单位分），必填
	PurchaseAmount string `json:"purchase_amount,omitempty" xml:"purchase_amount,omitempty"`
	// 购买件数，默认值1，必填
	ItemCount string `json:"item_count,omitempty" xml:"item_count,omitempty"`
	// 直达唤端媒体替换的requestid
	CkRequestId string `json:"ck_request_id,omitempty" xml:"ck_request_id,omitempty"`
	// 商品ID，必填
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品标题
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 转化时间戳（秒），必填
	EventTime string `json:"event_time,omitempty" xml:"event_time,omitempty"`
	// 媒体渠道ID，对应查询接口返回的channel_id，必填。1-巨量2.0，104-腾讯3.0，105-快手，108-B站，109-微博
	ChannelId string `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 默认值1，必填
	CallbackType string `json:"callback_type,omitempty" xml:"callback_type,omitempty"`
	// 媒体账户ID
	AdvertiserId string `json:"advertiser_id,omitempty" xml:"advertiser_id,omitempty"`
	// 店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 曝光回传使用
	IsDirectMatch string `json:"is_direct_match,omitempty" xml:"is_direct_match,omitempty"`
	// 曝光回传使用
	CallBackUrl string `json:"call_back_url,omitempty" xml:"call_back_url,omitempty"`
	// 曝光回传使用
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 曝光回传使用
	AdSrc string `json:"ad_src,omitempty" xml:"ad_src,omitempty"`
	// 曝光回传使用
	BusinessType string `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// ts
	Ts int64 `json:"ts,omitempty" xml:"ts,omitempty"`
}

var poolUDSmartServiceRequest = sync.Pool{
	New: func() any {
		return new(UDSmartServiceRequest)
	},
}

// GetUDSmartServiceRequest() 从对象池中获取UDSmartServiceRequest
func GetUDSmartServiceRequest() *UDSmartServiceRequest {
	return poolUDSmartServiceRequest.Get().(*UDSmartServiceRequest)
}

// ReleaseUDSmartServiceRequest 释放UDSmartServiceRequest
func ReleaseUDSmartServiceRequest(v *UDSmartServiceRequest) {
	v.TecAgent = ""
	v.CmClickId = ""
	v.CmRequestId = ""
	v.ProductId = ""
	v.OrderId = ""
	v.EventType = ""
	v.CkClickId = ""
	v.AppAccessToken = ""
	v.PurchaseAmount = ""
	v.ItemCount = ""
	v.CkRequestId = ""
	v.ItemId = ""
	v.ItemName = ""
	v.EventTime = ""
	v.ChannelId = ""
	v.CallbackType = ""
	v.AdvertiserId = ""
	v.ShopName = ""
	v.IsDirectMatch = ""
	v.CallBackUrl = ""
	v.TraceId = ""
	v.AdSrc = ""
	v.BusinessType = ""
	v.Ts = 0
	poolUDSmartServiceRequest.Put(v)
}
