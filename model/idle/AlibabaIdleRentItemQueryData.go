package idle

import (
	"sync"
)

// AlibabaIdleRentItemQueryData 结构体
type AlibabaIdleRentItemQueryData struct {
	// 商品sku信息
	ItemSkuList []ItemSkuDto `json:"item_sku_list,omitempty" xml:"item_sku_list>item_sku_dto,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 新旧水平
	UsedLevel int64 `json:"used_level,omitempty" xml:"used_level,omitempty"`
	// 库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 地址信息
	Address *AddressDto `json:"address,omitempty" xml:"address,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 价格信息
	PriceInfo *PriceDto `json:"price_info,omitempty" xml:"price_info,omitempty"`
	// 商品状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 运费模板id
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

var poolAlibabaIdleRentItemQueryData = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentItemQueryData)
	},
}

// GetAlibabaIdleRentItemQueryData() 从对象池中获取AlibabaIdleRentItemQueryData
func GetAlibabaIdleRentItemQueryData() *AlibabaIdleRentItemQueryData {
	return poolAlibabaIdleRentItemQueryData.Get().(*AlibabaIdleRentItemQueryData)
}

// ReleaseAlibabaIdleRentItemQueryData 释放AlibabaIdleRentItemQueryData
func ReleaseAlibabaIdleRentItemQueryData(v *AlibabaIdleRentItemQueryData) {
	v.ItemSkuList = v.ItemSkuList[:0]
	v.Title = ""
	v.Desc = ""
	v.UsedLevel = 0
	v.Quantity = 0
	v.Address = nil
	v.ItemId = 0
	v.PriceInfo = nil
	v.Status = 0
	v.TemplateId = 0
	poolAlibabaIdleRentItemQueryData.Put(v)
}
