package lstlogistics2

import (
	"sync"
)

// SubOrderParam 结构体
type SubOrderParam struct {
	// 外部erp子单号
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// erp商品编码
	OutItemCode string `json:"out_item_code,omitempty" xml:"out_item_code,omitempty"`
	// 商品条码
	ItemBarCode string `json:"item_bar_code,omitempty" xml:"item_bar_code,omitempty"`
	// 计量单位/规格 箱/个
	ItemUnit string `json:"item_unit,omitempty" xml:"item_unit,omitempty"`
	// lst商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// lst货品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 商品单价，单位：分
	ItemPrice int64 `json:"item_price,omitempty" xml:"item_price,omitempty"`
	// 商品数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
	// 支付金额，单位：分
	PayAmount int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// 长 单位: mm
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 宽 单位: mm
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 高 单位: mm
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 重量 单位：克
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 体积 单位：立方毫米
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
}

var poolSubOrderParam = sync.Pool{
	New: func() any {
		return new(SubOrderParam)
	},
}

// GetSubOrderParam() 从对象池中获取SubOrderParam
func GetSubOrderParam() *SubOrderParam {
	return poolSubOrderParam.Get().(*SubOrderParam)
}

// ReleaseSubOrderParam 释放SubOrderParam
func ReleaseSubOrderParam(v *SubOrderParam) {
	v.SubOrderId = ""
	v.ItemName = ""
	v.OutItemCode = ""
	v.ItemBarCode = ""
	v.ItemUnit = ""
	v.ItemId = 0
	v.ScItemId = 0
	v.ItemPrice = 0
	v.ItemQuantity = 0
	v.PayAmount = 0
	v.Length = 0
	v.Width = 0
	v.Height = 0
	v.Weight = 0
	v.Volume = 0
	poolSubOrderParam.Put(v)
}
