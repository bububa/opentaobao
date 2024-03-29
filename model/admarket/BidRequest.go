package admarket

import (
	"sync"
)

// BidRequest 结构体
type BidRequest struct {
	// 广告位列表
	AdSlots []AdSlot `json:"ad_slots,omitempty" xml:"ad_slots>ad_slot,omitempty"`
	// 渠道id
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 版本号
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// 批次id
	BatchId string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 应用信息
	AppInfo *AppInfo `json:"app_info,omitempty" xml:"app_info,omitempty"`
	// sdk信息
	SdkInfo *SdkInfo `json:"sdk_info,omitempty" xml:"sdk_info,omitempty"`
	// 设备信息
	DeviceInfo *DeviceInfo `json:"device_info,omitempty" xml:"device_info,omitempty"`
	// 设备id
	Udid *Udid `json:"udid,omitempty" xml:"udid,omitempty"`
	// 用户信息
	UserInfo *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
	// 网络信息
	Network *Network `json:"network,omitempty" xml:"network,omitempty"`
	// 定位信息
	Location *Location `json:"location,omitempty" xml:"location,omitempty"`
}

var poolBidRequest = sync.Pool{
	New: func() any {
		return new(BidRequest)
	},
}

// GetBidRequest() 从对象池中获取BidRequest
func GetBidRequest() *BidRequest {
	return poolBidRequest.Get().(*BidRequest)
}

// ReleaseBidRequest 释放BidRequest
func ReleaseBidRequest(v *BidRequest) {
	v.AdSlots = v.AdSlots[:0]
	v.Channel = ""
	v.Version = ""
	v.BatchId = ""
	v.AppInfo = nil
	v.SdkInfo = nil
	v.DeviceInfo = nil
	v.Udid = nil
	v.UserInfo = nil
	v.Network = nil
	v.Location = nil
	poolBidRequest.Put(v)
}
