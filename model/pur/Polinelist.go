package pur

import (
	"sync"
)

// Polinelist 结构体
type Polinelist struct {
	// BOM节点列表，节点之间使用父节点id维护层级关系
	ItemList []Polinestructureitemdtolist `json:"item_list,omitempty" xml:"item_list>polinestructureitemdtolist,omitempty"`
	// 物品名称
	GoodsName string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
	// 规格
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 单位
	Uom string `json:"uom,omitempty" xml:"uom,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 是否涉及预提所得税
	IsWhtRelated string `json:"is_wht_related,omitempty" xml:"is_wht_related,omitempty"`
	// 含税单价
	TaxPrice string `json:"tax_price,omitempty" xml:"tax_price,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 含税总价
	TaxAmount string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// 收货地址
	DeliveryLocation string `json:"delivery_location,omitempty" xml:"delivery_location,omitempty"`
	// 期望交付日期
	NeedByDate string `json:"need_by_date,omitempty" xml:"need_by_date,omitempty"`
	// 期望交付日期
	NeedByDateEnd string `json:"need_by_date_end,omitempty" xml:"need_by_date_end,omitempty"`
	// 收货人
	Receiver string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// 收货人电话
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 收货人手机
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// additionalInfo
	AdditionalInfo string `json:"additional_info,omitempty" xml:"additional_info,omitempty"`
	// 订单行商品报价币种和订单头币种之间的汇率
	ExchangeRate string `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
	// 订单行商品报价币种
	GoodsCurrencyCode string `json:"goods_currency_code,omitempty" xml:"goods_currency_code,omitempty"`
	// lineNum
	LineNum int64 `json:"line_num,omitempty" xml:"line_num,omitempty"`
}

var poolPolinelist = sync.Pool{
	New: func() any {
		return new(Polinelist)
	},
}

// GetPolinelist() 从对象池中获取Polinelist
func GetPolinelist() *Polinelist {
	return poolPolinelist.Get().(*Polinelist)
}

// ReleasePolinelist 释放Polinelist
func ReleasePolinelist(v *Polinelist) {
	v.ItemList = v.ItemList[:0]
	v.GoodsName = ""
	v.Description = ""
	v.Uom = ""
	v.Quantity = ""
	v.TaxRate = ""
	v.IsWhtRelated = ""
	v.TaxPrice = ""
	v.CurrencyCode = ""
	v.TaxAmount = ""
	v.DeliveryLocation = ""
	v.NeedByDate = ""
	v.NeedByDateEnd = ""
	v.Receiver = ""
	v.ReceiverPhone = ""
	v.ReceiverMobile = ""
	v.AdditionalInfo = ""
	v.ExchangeRate = ""
	v.GoodsCurrencyCode = ""
	v.LineNum = 0
	poolPolinelist.Put(v)
}
