package pur

import (
	"sync"
)

// Polinestructureitemdtolist 结构体
type Polinestructureitemdtolist struct {
	// 节点名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// ITEM_PRICE：物料价格，BOM_COST：BOM COST，OTHER_COST_TOTAL：其他费用小计，PROFIT_COST：利润，OEM_COST：代工费用，MAINTENANCE_COST：维保费用，LOGISTICS_COST：物流费用，INSTALLATION_COST：上架安装费用，TARIFF_TAX：关税税费，SALES_TAX：销售税税费，VAT_TAX：增值税税费
	CostType string `json:"cost_type,omitempty" xml:"cost_type,omitempty"`
	// 计价单位
	ChargeUnit string `json:"charge_unit,omitempty" xml:"charge_unit,omitempty"`
	// 单价
	UnitPrice string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 单价币种
	UnitCurrency string `json:"unit_currency,omitempty" xml:"unit_currency,omitempty"`
	// 金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 金额币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 汇率
	ExchangeRate string `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 参数
	ExpenseParam string `json:"expense_param,omitempty" xml:"expense_param,omitempty"`
	// 节点id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 父节点id，为空则表示该节点是根节点
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
}

var poolPolinestructureitemdtolist = sync.Pool{
	New: func() any {
		return new(Polinestructureitemdtolist)
	},
}

// GetPolinestructureitemdtolist() 从对象池中获取Polinestructureitemdtolist
func GetPolinestructureitemdtolist() *Polinestructureitemdtolist {
	return poolPolinestructureitemdtolist.Get().(*Polinestructureitemdtolist)
}

// ReleasePolinestructureitemdtolist 释放Polinestructureitemdtolist
func ReleasePolinestructureitemdtolist(v *Polinestructureitemdtolist) {
	v.ItemName = ""
	v.CostType = ""
	v.ChargeUnit = ""
	v.UnitPrice = ""
	v.UnitCurrency = ""
	v.Amount = ""
	v.Currency = ""
	v.ExchangeRate = ""
	v.Quantity = ""
	v.ExpenseParam = ""
	v.Id = 0
	v.ParentId = 0
	poolPolinestructureitemdtolist.Put(v)
}
