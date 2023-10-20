package alsc

import (
	"sync"
)

// PrizeItemInfo 结构体
type PrizeItemInfo struct {
	// 券id
	VoucherIds []string `json:"voucher_ids,omitempty" xml:"voucher_ids>string,omitempty"`
	// 优惠金额
	Denomination string `json:"denomination,omitempty" xml:"denomination,omitempty"`
	// 奖品名称
	PrizeName string `json:"prize_name,omitempty" xml:"prize_name,omitempty"`
	// 券名称
	VoucherName string `json:"voucher_name,omitempty" xml:"voucher_name,omitempty"`
	// 券类型
	VoucherType string `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	// 几等奖
	PrizeLevel int64 `json:"prize_level,omitempty" xml:"prize_level,omitempty"`
}

var poolPrizeItemInfo = sync.Pool{
	New: func() any {
		return new(PrizeItemInfo)
	},
}

// GetPrizeItemInfo() 从对象池中获取PrizeItemInfo
func GetPrizeItemInfo() *PrizeItemInfo {
	return poolPrizeItemInfo.Get().(*PrizeItemInfo)
}

// ReleasePrizeItemInfo 释放PrizeItemInfo
func ReleasePrizeItemInfo(v *PrizeItemInfo) {
	v.VoucherIds = v.VoucherIds[:0]
	v.Denomination = ""
	v.PrizeName = ""
	v.VoucherName = ""
	v.VoucherType = ""
	v.PrizeLevel = 0
	poolPrizeItemInfo.Put(v)
}
