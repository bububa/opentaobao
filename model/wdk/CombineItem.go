package wdk

import (
	"sync"
)

// CombineItem 结构体
type CombineItem struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品名称
	AuctionTitle string `json:"auction_title,omitempty" xml:"auction_title,omitempty"`
	// 商品条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 库存单位购买数量
	BuyAmountStock string `json:"buy_amount_stock,omitempty" xml:"buy_amount_stock,omitempty"`
	// 销售单位
	SaleUnit string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
	// 库存单位
	StockUnit string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
	// 销售单位购买数量
	BuyAmountSale int64 `json:"buy_amount_sale,omitempty" xml:"buy_amount_sale,omitempty"`
	// 商品价格
	AuctionPrice int64 `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
}

var poolCombineItem = sync.Pool{
	New: func() any {
		return new(CombineItem)
	},
}

// GetCombineItem() 从对象池中获取CombineItem
func GetCombineItem() *CombineItem {
	return poolCombineItem.Get().(*CombineItem)
}

// ReleaseCombineItem 释放CombineItem
func ReleaseCombineItem(v *CombineItem) {
	v.SkuCode = ""
	v.AuctionTitle = ""
	v.BarCode = ""
	v.BuyAmountStock = ""
	v.SaleUnit = ""
	v.StockUnit = ""
	v.BuyAmountSale = 0
	v.AuctionPrice = 0
	poolCombineItem.Put(v)
}
