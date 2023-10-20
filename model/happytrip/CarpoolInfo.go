package happytrip

import (
	"sync"
)

// CarpoolInfo 结构体
type CarpoolInfo struct {
	// 乘客订单信息
	PassengerOrders []PassengerOrderInfo `json:"passenger_orders,omitempty" xml:"passenger_orders>passenger_order_info,omitempty"`
	// 拼车单id
	CarpoolOrderId string `json:"carpool_order_id,omitempty" xml:"carpool_order_id,omitempty"`
	// 0: 未拼成 1:拼成
	CarpoolFlag int64 `json:"carpool_flag,omitempty" xml:"carpool_flag,omitempty"`
}

var poolCarpoolInfo = sync.Pool{
	New: func() any {
		return new(CarpoolInfo)
	},
}

// GetCarpoolInfo() 从对象池中获取CarpoolInfo
func GetCarpoolInfo() *CarpoolInfo {
	return poolCarpoolInfo.Get().(*CarpoolInfo)
}

// ReleaseCarpoolInfo 释放CarpoolInfo
func ReleaseCarpoolInfo(v *CarpoolInfo) {
	v.PassengerOrders = v.PassengerOrders[:0]
	v.CarpoolOrderId = ""
	v.CarpoolFlag = 0
	poolCarpoolInfo.Put(v)
}
