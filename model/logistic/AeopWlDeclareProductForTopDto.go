package logistic

import (
	"sync"
)

// AeopWlDeclareProductForTopDto 结构体
type AeopWlDeclareProductForTopDto struct {
	// 类目中文名称
	CategoryCnDesc string `json:"category_cn_desc,omitempty" xml:"category_cn_desc,omitempty"`
	// 类目英文名称
	CategoryEnDesc string `json:"category_en_desc,omitempty" xml:"category_en_desc,omitempty"`
	// 海关编码
	HsCode string `json:"hs_code,omitempty" xml:"hs_code,omitempty"`
	// 产品申报金额
	ProductDeclareAmount string `json:"product_declare_amount,omitempty" xml:"product_declare_amount,omitempty"`
	// 产品重量
	ProductWeight string `json:"product_weight,omitempty" xml:"product_weight,omitempty"`
	// scItem code
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// scItem name
	ScItemName string `json:"sc_item_name,omitempty" xml:"sc_item_name,omitempty"`
	// sku code
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// sku value
	SkuValue string `json:"sku_value,omitempty" xml:"sku_value,omitempty"`
	// 商品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 产品数量
	ProductNum int64 `json:"product_num,omitempty" xml:"product_num,omitempty"`
	// scItem id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 交易子单号
	ChildOrderId int64 `json:"child_order_id,omitempty" xml:"child_order_id,omitempty"`
	// 判断是否属于非液体化妆品
	AneroidMarkup bool `json:"aneroid_markup,omitempty" xml:"aneroid_markup,omitempty"`
	// 是否易碎
	Breakable bool `json:"breakable,omitempty" xml:"breakable,omitempty"`
	// 是否包含电池
	ContainsBattery bool `json:"contains_battery,omitempty" xml:"contains_battery,omitempty"`
	// 是否纯电池
	OnlyBattery bool `json:"only_battery,omitempty" xml:"only_battery,omitempty"`
	// 是否特货
	ContainsSpecialGoods bool `json:"contains_special_goods,omitempty" xml:"contains_special_goods,omitempty"`
}

var poolAeopWlDeclareProductForTopDto = sync.Pool{
	New: func() any {
		return new(AeopWlDeclareProductForTopDto)
	},
}

// GetAeopWlDeclareProductForTopDto() 从对象池中获取AeopWlDeclareProductForTopDto
func GetAeopWlDeclareProductForTopDto() *AeopWlDeclareProductForTopDto {
	return poolAeopWlDeclareProductForTopDto.Get().(*AeopWlDeclareProductForTopDto)
}

// ReleaseAeopWlDeclareProductForTopDto 释放AeopWlDeclareProductForTopDto
func ReleaseAeopWlDeclareProductForTopDto(v *AeopWlDeclareProductForTopDto) {
	v.CategoryCnDesc = ""
	v.CategoryEnDesc = ""
	v.HsCode = ""
	v.ProductDeclareAmount = ""
	v.ProductWeight = ""
	v.ScItemCode = ""
	v.ScItemName = ""
	v.SkuCode = ""
	v.SkuValue = ""
	v.ProductId = 0
	v.ProductNum = 0
	v.ScItemId = 0
	v.ChildOrderId = 0
	v.AneroidMarkup = false
	v.Breakable = false
	v.ContainsBattery = false
	v.OnlyBattery = false
	v.ContainsSpecialGoods = false
	poolAeopWlDeclareProductForTopDto.Put(v)
}
