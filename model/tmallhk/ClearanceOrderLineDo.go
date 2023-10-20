package tmallhk

import (
	"sync"
)

// ClearanceOrderLineDo 结构体
type ClearanceOrderLineDo struct {
	// 税费封装，示例：&#34;declaration&#34;: {             &#34;品牌&#34;: &#34;1&#34;,             &#34;用途&#34;: &#34;1&#34;,             &#34;品名&#34;: &#34;1&#34;,             &#34;包装规格&#34;: &#34;1&#34;           }
	Declaration string `json:"declaration,omitempty" xml:"declaration,omitempty"`
	// 销售属性
	SellProperty string `json:"sell_property,omitempty" xml:"sell_property,omitempty"`
	// 销售单位
	SaleUnit string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
	// 税费封装
	TaxDO *ClearanceTaxDo `json:"tax_d_o,omitempty" xml:"tax_d_o,omitempty"`
	// 货款
	ActualValue int64 `json:"actual_value,omitempty" xml:"actual_value,omitempty"`
	// 计量单位封装
	UnitDO *ClearanceTaxDo `json:"unit_d_o,omitempty" xml:"unit_d_o,omitempty"`
	// 货品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 淘系商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolClearanceOrderLineDo = sync.Pool{
	New: func() any {
		return new(ClearanceOrderLineDo)
	},
}

// GetClearanceOrderLineDo() 从对象池中获取ClearanceOrderLineDo
func GetClearanceOrderLineDo() *ClearanceOrderLineDo {
	return poolClearanceOrderLineDo.Get().(*ClearanceOrderLineDo)
}

// ReleaseClearanceOrderLineDo 释放ClearanceOrderLineDo
func ReleaseClearanceOrderLineDo(v *ClearanceOrderLineDo) {
	v.Declaration = ""
	v.SellProperty = ""
	v.SaleUnit = ""
	v.TaxDO = nil
	v.ActualValue = 0
	v.UnitDO = nil
	v.ScItemId = 0
	v.ItemId = 0
	poolClearanceOrderLineDo.Put(v)
}
