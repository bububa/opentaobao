package fundplatform

import (
	"sync"
)

// InputInvoiceLineLedgerDto 结构体
type InputInvoiceLineLedgerDto struct {
	// 总金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 不含税金额
	ExcludingTaxAmount string `json:"excluding_tax_amount,omitempty" xml:"excluding_tax_amount,omitempty"`
	// 货物名称
	GoodsDesc string `json:"goods_desc,omitempty" xml:"goods_desc,omitempty"`
	// 规格型号
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 单位
	QuantityUnit string `json:"quantity_unit,omitempty" xml:"quantity_unit,omitempty"`
	// 不含税金额
	TaxAmount string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// 单价
	UnitPrice string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
}

var poolInputInvoiceLineLedgerDto = sync.Pool{
	New: func() any {
		return new(InputInvoiceLineLedgerDto)
	},
}

// GetInputInvoiceLineLedgerDto() 从对象池中获取InputInvoiceLineLedgerDto
func GetInputInvoiceLineLedgerDto() *InputInvoiceLineLedgerDto {
	return poolInputInvoiceLineLedgerDto.Get().(*InputInvoiceLineLedgerDto)
}

// ReleaseInputInvoiceLineLedgerDto 释放InputInvoiceLineLedgerDto
func ReleaseInputInvoiceLineLedgerDto(v *InputInvoiceLineLedgerDto) {
	v.Amount = ""
	v.ExcludingTaxAmount = ""
	v.GoodsDesc = ""
	v.Model = ""
	v.Quantity = ""
	v.QuantityUnit = ""
	v.TaxAmount = ""
	v.UnitPrice = ""
	v.TaxRate = ""
	poolInputInvoiceLineLedgerDto.Put(v)
}
