package ascp

import (
	"sync"
)

// SubScItemModel 结构体
type SubScItemModel struct {
	// 货品商家编码
	ScitemCode string `json:"scitem_code,omitempty" xml:"scitem_code,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 子货品id
	ScitemId string `json:"scitem_id,omitempty" xml:"scitem_id,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 子货品单价(人命币则为分)
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 是否固定价格 1是 0否
	FixedPrice int64 `json:"fixed_price,omitempty" xml:"fixed_price,omitempty"`
}

var poolSubScItemModel = sync.Pool{
	New: func() any {
		return new(SubScItemModel)
	},
}

// GetSubScItemModel() 从对象池中获取SubScItemModel
func GetSubScItemModel() *SubScItemModel {
	return poolSubScItemModel.Get().(*SubScItemModel)
}

// ReleaseSubScItemModel 释放SubScItemModel
func ReleaseSubScItemModel(v *SubScItemModel) {
	v.ScitemCode = ""
	v.Currency = ""
	v.ScitemId = ""
	v.Quantity = 0
	v.RetailPrice = 0
	v.FixedPrice = 0
	poolSubScItemModel.Put(v)
}
