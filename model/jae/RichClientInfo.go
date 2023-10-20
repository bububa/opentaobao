package jae

import (
	"sync"
)

// RichClientInfo 结构体
type RichClientInfo struct {
	// appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 用户昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 设备号
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 来源ip
	Ip string `json:"ip,omitempty" xml:"ip,omitempty"`
	// 用户id
	BuyerId int64 `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
}

var poolRichClientInfo = sync.Pool{
	New: func() any {
		return new(RichClientInfo)
	},
}

// GetRichClientInfo() 从对象池中获取RichClientInfo
func GetRichClientInfo() *RichClientInfo {
	return poolRichClientInfo.Get().(*RichClientInfo)
}

// ReleaseRichClientInfo 释放RichClientInfo
func ReleaseRichClientInfo(v *RichClientInfo) {
	v.Appkey = ""
	v.BuyerNick = ""
	v.Lng = ""
	v.Lat = ""
	v.DeviceId = ""
	v.Ip = ""
	v.BuyerId = 0
	poolRichClientInfo.Put(v)
}
