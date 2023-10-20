package fpm

import (
	"sync"
)

// InvoiceDetails 结构体
type InvoiceDetails struct {
	// 单价
	UnitPrice string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 不含税金额
	AmountWithoutTax string `json:"amount_without_tax,omitempty" xml:"amount_without_tax,omitempty"`
	// 规格型号
	ItemSpec string `json:"item_spec,omitempty" xml:"item_spec,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 货物或应税劳务名称
	CargoName string `json:"cargo_name,omitempty" xml:"cargo_name,omitempty"`
	// 单位
	QuantityUnit string `json:"quantity_unit,omitempty" xml:"quantity_unit,omitempty"`
	// 税额
	TaxAmount string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
}

var poolInvoiceDetails = sync.Pool{
	New: func() any {
		return new(InvoiceDetails)
	},
}

// GetInvoiceDetails() 从对象池中获取InvoiceDetails
func GetInvoiceDetails() *InvoiceDetails {
	return poolInvoiceDetails.Get().(*InvoiceDetails)
}

// ReleaseInvoiceDetails 释放InvoiceDetails
func ReleaseInvoiceDetails(v *InvoiceDetails) {
	v.UnitPrice = ""
	v.AmountWithoutTax = ""
	v.ItemSpec = ""
	v.TaxRate = ""
	v.Quantity = ""
	v.CargoName = ""
	v.QuantityUnit = ""
	v.TaxAmount = ""
	poolInvoiceDetails.Put(v)
}
