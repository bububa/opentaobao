package alsc

import (
	"sync"
)

// Item 结构体
type Item struct {
	// 商品名称
	GoodsName string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
	// 商品规格，需保证唯一
	Sku string `json:"sku,omitempty" xml:"sku,omitempty"`
	// 商品码
	Spu string `json:"spu,omitempty" xml:"spu,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 一级分类
	PrimaryClass string `json:"primary_class,omitempty" xml:"primary_class,omitempty"`
	// 二级分类
	SecondaryClass string `json:"secondary_class,omitempty" xml:"secondary_class,omitempty"`
	// 实际费用
	ActualFee int64 `json:"actual_fee,omitempty" xml:"actual_fee,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 优惠金额
	PromFee int64 `json:"prom_fee,omitempty" xml:"prom_fee,omitempty"`
	// 总金额
	TotalFee int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
}

var poolItem = sync.Pool{
	New: func() any {
		return new(Item)
	},
}

// GetItem() 从对象池中获取Item
func GetItem() *Item {
	return poolItem.Get().(*Item)
}

// ReleaseItem 释放Item
func ReleaseItem(v *Item) {
	v.GoodsName = ""
	v.Sku = ""
	v.Spu = ""
	v.Unit = ""
	v.Quantity = ""
	v.PrimaryClass = ""
	v.SecondaryClass = ""
	v.ActualFee = 0
	v.Price = 0
	v.PromFee = 0
	v.TotalFee = 0
	poolItem.Put(v)
}
