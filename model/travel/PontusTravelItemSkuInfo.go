package travel

import (
	"sync"
)

// PontusTravelItemSkuInfo 结构体
type PontusTravelItemSkuInfo struct {
	// 套餐的日历价格库存。如果是预约商品，只需要填写一个Price，并且，不需要填写Price中的date字段不填，且预约商品只有成人价格和库存。
	Prices []PontusTravelPrices `json:"prices,omitempty" xml:"prices>pontus_travel_prices,omitempty"`
	// 套餐关联的产品元素信息
	Products []PontusTravelProduct `json:"products,omitempty" xml:"products>pontus_travel_product,omitempty"`
	// 映射到具体日历价格套餐的外部商家编码
	OuterSkuId string `json:"outer_sku_id,omitempty" xml:"outer_sku_id,omitempty"`
	// 套餐描述
	PackageDesc string `json:"package_desc,omitempty" xml:"package_desc,omitempty"`
	// 套餐名称
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// 套餐下面对应商品元素信息 仅针对新版商品
	Combos string `json:"combos,omitempty" xml:"combos,omitempty"`
	// 邮轮房型名称
	RoomTypeName string `json:"room_type_name,omitempty" xml:"room_type_name,omitempty"`
	// 邮轮房型ID，新版本有值
	RoomTypeId int64 `json:"room_type_id,omitempty" xml:"room_type_id,omitempty"`
	// 邮轮房型类型
	RoomType int64 `json:"room_type,omitempty" xml:"room_type,omitempty"`
	// 人数
	PeopleNumber int64 `json:"people_number,omitempty" xml:"people_number,omitempty"`
	// 下单人数是否与房型人数一致
	OrderCountMatch bool `json:"order_count_match,omitempty" xml:"order_count_match,omitempty"`
}

var poolPontusTravelItemSkuInfo = sync.Pool{
	New: func() any {
		return new(PontusTravelItemSkuInfo)
	},
}

// GetPontusTravelItemSkuInfo() 从对象池中获取PontusTravelItemSkuInfo
func GetPontusTravelItemSkuInfo() *PontusTravelItemSkuInfo {
	return poolPontusTravelItemSkuInfo.Get().(*PontusTravelItemSkuInfo)
}

// ReleasePontusTravelItemSkuInfo 释放PontusTravelItemSkuInfo
func ReleasePontusTravelItemSkuInfo(v *PontusTravelItemSkuInfo) {
	v.Prices = v.Prices[:0]
	v.Products = v.Products[:0]
	v.OuterSkuId = ""
	v.PackageDesc = ""
	v.PackageName = ""
	v.Combos = ""
	v.RoomTypeName = ""
	v.RoomTypeId = 0
	v.RoomType = 0
	v.PeopleNumber = 0
	v.OrderCountMatch = false
	poolPontusTravelItemSkuInfo.Put(v)
}
