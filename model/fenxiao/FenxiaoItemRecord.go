package fenxiao

import (
	"sync"
)

// FenxiaoItemRecord 结构体
type FenxiaoItemRecord struct {
	// 下载时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 更新时间（系统时间，无业务意义）
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）
	TradeType string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	// 分销商ID
	DistributorId int64 `json:"distributor_id,omitempty" xml:"distributor_id,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 产品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolFenxiaoItemRecord = sync.Pool{
	New: func() any {
		return new(FenxiaoItemRecord)
	},
}

// GetFenxiaoItemRecord() 从对象池中获取FenxiaoItemRecord
func GetFenxiaoItemRecord() *FenxiaoItemRecord {
	return poolFenxiaoItemRecord.Get().(*FenxiaoItemRecord)
}

// ReleaseFenxiaoItemRecord 释放FenxiaoItemRecord
func ReleaseFenxiaoItemRecord(v *FenxiaoItemRecord) {
	v.Created = ""
	v.Modified = ""
	v.TradeType = ""
	v.DistributorId = 0
	v.ItemId = 0
	v.ProductId = 0
	poolFenxiaoItemRecord.Put(v)
}
