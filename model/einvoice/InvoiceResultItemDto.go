package einvoice

import (
	"sync"
)

// InvoiceResultItemDto 结构体
type InvoiceResultItemDto struct {
	// 数量，最多6位小数。  折扣行此参数不能传，非折扣行必传。存在则需&gt;0
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 单价（不含税），格式为2位小数。最大支持6位小数，不足2位小数时需转化为2位小数格式。折扣行此参数不能传，非折扣行必传
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 税率。格式为2位小数，如：0.00, 0.03, 0.13等等
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 发票项目编号（或商品编号）
	ItemNo string `json:"item_no,omitempty" xml:"item_no,omitempty"`
	// 发票项目名称（或商品名称）
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 含税总金额 (等于sum_price和tax之和)，  单位：元，格式为2位小数，精度2位小数
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 单位。折扣行不能传，非折扣行必传
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 规格型号
	Specification string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 0税率标识，只有税率为0的情况才有值，0=出口零税率，1=免税，2=不征收，3=普通零税率
	ZeroRateFlag string `json:"zero_rate_flag,omitempty" xml:"zero_rate_flag,omitempty"`
	// 税额， 格式为2位小数
	Tax string `json:"tax,omitempty" xml:"tax,omitempty"`
	// 不含税总金额，格式为2位小数  单位：元，精度2位小数
	SumPrice string `json:"sum_price,omitempty" xml:"sum_price,omitempty"`
	// 发票行性质。0表示正常行，1表示折扣行，2表示被折扣行。  比如充电器单价100元，折扣10元，则明细为2行，充电器行性质为2，折扣行性质为1。如果充电器没有折扣，则值应为0
	RowType int64 `json:"row_type,omitempty" xml:"row_type,omitempty"`
}

var poolInvoiceResultItemDto = sync.Pool{
	New: func() any {
		return new(InvoiceResultItemDto)
	},
}

// GetInvoiceResultItemDto() 从对象池中获取InvoiceResultItemDto
func GetInvoiceResultItemDto() *InvoiceResultItemDto {
	return poolInvoiceResultItemDto.Get().(*InvoiceResultItemDto)
}

// ReleaseInvoiceResultItemDto 释放InvoiceResultItemDto
func ReleaseInvoiceResultItemDto(v *InvoiceResultItemDto) {
	v.Quantity = ""
	v.Price = ""
	v.TaxRate = ""
	v.ItemNo = ""
	v.ItemName = ""
	v.Amount = ""
	v.Unit = ""
	v.Specification = ""
	v.ZeroRateFlag = ""
	v.Tax = ""
	v.SumPrice = ""
	v.RowType = 0
	poolInvoiceResultItemDto.Put(v)
}
