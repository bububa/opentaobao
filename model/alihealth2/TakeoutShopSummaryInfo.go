package alihealth2

import (
	"sync"
)

// TakeoutShopSummaryInfo 结构体
type TakeoutShopSummaryInfo struct {
	// 店铺分店名
	SubName string `json:"sub_name,omitempty" xml:"sub_name,omitempty"`
	// 外卖店铺名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 店铺地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 电话号码
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 店铺与ISV的关联ID
	Shopoutid string `json:"shopoutid,omitempty" xml:"shopoutid,omitempty"`
	// 外卖店铺id
	Shopid int64 `json:"shopid,omitempty" xml:"shopid,omitempty"`
	// 等待确认配送的订单笔数
	WaitConfirm int64 `json:"wait_confirm,omitempty" xml:"wait_confirm,omitempty"`
	// 等待确认的兑换券的订单笔数
	DigitalWaitConfirm int64 `json:"digital_wait_confirm,omitempty" xml:"digital_wait_confirm,omitempty"`
	// 等待退款的订单笔数
	WaitRefund int64 `json:"wait_refund,omitempty" xml:"wait_refund,omitempty"`
	// 店铺营业状态，歇业：0，营业：1
	IsOpen int64 `json:"is_open,omitempty" xml:"is_open,omitempty"`
}

var poolTakeoutShopSummaryInfo = sync.Pool{
	New: func() any {
		return new(TakeoutShopSummaryInfo)
	},
}

// GetTakeoutShopSummaryInfo() 从对象池中获取TakeoutShopSummaryInfo
func GetTakeoutShopSummaryInfo() *TakeoutShopSummaryInfo {
	return poolTakeoutShopSummaryInfo.Get().(*TakeoutShopSummaryInfo)
}

// ReleaseTakeoutShopSummaryInfo 释放TakeoutShopSummaryInfo
func ReleaseTakeoutShopSummaryInfo(v *TakeoutShopSummaryInfo) {
	v.SubName = ""
	v.Name = ""
	v.City = ""
	v.Address = ""
	v.Phone = ""
	v.Shopoutid = ""
	v.Shopid = 0
	v.WaitConfirm = 0
	v.DigitalWaitConfirm = 0
	v.WaitRefund = 0
	v.IsOpen = 0
	poolTakeoutShopSummaryInfo.Put(v)
}
