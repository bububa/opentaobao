package tbtrade

import (
	"sync"
)

// BillDetailInfo 结构体
type BillDetailInfo struct {
	// 店铺名称
	ShopNick string `json:"shop_nick,omitempty" xml:"shop_nick,omitempty"`
	// 账单计费
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
	// 已弃用,查询账单详情请使用taobao.top.secret.bill.detail接口。
	SecretNo string `json:"secret_no,omitempty" xml:"secret_no,omitempty"`
	// 已弃用,查询账单详情请使用taobao.top.secret.bill.detail接口。
	Start string `json:"start,omitempty" xml:"start,omitempty"`
	// 已弃用,查询账单详情请使用taobao.top.secret.bill.detail接口。
	End string `json:"end,omitempty" xml:"end,omitempty"`
	// 商家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 0-号租费，1-通话费
	BillItem int64 `json:"bill_item,omitempty" xml:"bill_item,omitempty"`
	// 账单持续时间，号租费时单位为天，通话费单位为分钟
	BillDuration int64 `json:"bill_duration,omitempty" xml:"bill_duration,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolBillDetailInfo = sync.Pool{
	New: func() any {
		return new(BillDetailInfo)
	},
}

// GetBillDetailInfo() 从对象池中获取BillDetailInfo
func GetBillDetailInfo() *BillDetailInfo {
	return poolBillDetailInfo.Get().(*BillDetailInfo)
}

// ReleaseBillDetailInfo 释放BillDetailInfo
func ReleaseBillDetailInfo(v *BillDetailInfo) {
	v.ShopNick = ""
	v.Cost = ""
	v.SecretNo = ""
	v.Start = ""
	v.End = ""
	v.SellerNick = ""
	v.BillItem = 0
	v.BillDuration = 0
	v.OrderId = 0
	poolBillDetailInfo.Put(v)
}
