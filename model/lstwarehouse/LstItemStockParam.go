package lstwarehouse

import (
	"sync"
)

// LstItemStockParam 结构体
type LstItemStockParam struct {
	// 库存参数列表
	StockList []Stocklist `json:"stock_list,omitempty" xml:"stock_list>stocklist,omitempty"`
}

var poolLstItemStockParam = sync.Pool{
	New: func() any {
		return new(LstItemStockParam)
	},
}

// GetLstItemStockParam() 从对象池中获取LstItemStockParam
func GetLstItemStockParam() *LstItemStockParam {
	return poolLstItemStockParam.Get().(*LstItemStockParam)
}

// ReleaseLstItemStockParam 释放LstItemStockParam
func ReleaseLstItemStockParam(v *LstItemStockParam) {
	v.StockList = v.StockList[:0]
	poolLstItemStockParam.Put(v)
}
