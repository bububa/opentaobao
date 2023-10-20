package tuanhotel

import (
	"sync"
)

// TuanItemOnlineSkuManagerVo 结构体
type TuanItemOnlineSkuManagerVo struct {
	// 套餐名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// sku id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 间夜
	Nights int64 `json:"nights,omitempty" xml:"nights,omitempty"`
}

var poolTuanItemOnlineSkuManagerVo = sync.Pool{
	New: func() any {
		return new(TuanItemOnlineSkuManagerVo)
	},
}

// GetTuanItemOnlineSkuManagerVo() 从对象池中获取TuanItemOnlineSkuManagerVo
func GetTuanItemOnlineSkuManagerVo() *TuanItemOnlineSkuManagerVo {
	return poolTuanItemOnlineSkuManagerVo.Get().(*TuanItemOnlineSkuManagerVo)
}

// ReleaseTuanItemOnlineSkuManagerVo 释放TuanItemOnlineSkuManagerVo
func ReleaseTuanItemOnlineSkuManagerVo(v *TuanItemOnlineSkuManagerVo) {
	v.SkuName = ""
	v.Price = ""
	v.SkuId = 0
	v.Nights = 0
	poolTuanItemOnlineSkuManagerVo.Put(v)
}
