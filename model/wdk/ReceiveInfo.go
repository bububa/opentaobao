package wdk

import (
	"sync"
)

// ReceiveInfo 结构体
type ReceiveInfo struct {
	// 收货人姓名, 格式: 刘**; (商家自配送场景给出)
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 收货人联系方式, 虚拟小号; (商家自配送场景给出)
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 收货人地址信息; (商家自配送场景给出)
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// 收货人下单备注
	ReceiverMemo string `json:"receiver_memo,omitempty" xml:"receiver_memo,omitempty"`
	// 收货人经纬度; (商家自配送场景给出)
	ReceiverPoi string `json:"receiver_poi,omitempty" xml:"receiver_poi,omitempty"`
	// 期望收货时间
	ExpectArriveTime string `json:"expect_arrive_time,omitempty" xml:"expect_arrive_time,omitempty"`
	// 收货人手机号
	ReceiverPrivacyPhone string `json:"receiver_privacy_phone,omitempty" xml:"receiver_privacy_phone,omitempty"`
}

var poolReceiveInfo = sync.Pool{
	New: func() any {
		return new(ReceiveInfo)
	},
}

// GetReceiveInfo() 从对象池中获取ReceiveInfo
func GetReceiveInfo() *ReceiveInfo {
	return poolReceiveInfo.Get().(*ReceiveInfo)
}

// ReleaseReceiveInfo 释放ReceiveInfo
func ReleaseReceiveInfo(v *ReceiveInfo) {
	v.ReceiverName = ""
	v.ReceiverPhone = ""
	v.ReceiverAddress = ""
	v.ReceiverMemo = ""
	v.ReceiverPoi = ""
	v.ExpectArriveTime = ""
	v.ReceiverPrivacyPhone = ""
	poolReceiveInfo.Put(v)
}
