package einvoice

import (
	"sync"
)

// UserInvoiceItemDto 结构体
type UserInvoiceItemDto struct {
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 订单日
	BillDate string `json:"bill_date,omitempty" xml:"bill_date,omitempty"`
	// 商品类型：1. 税控单元，2.开票日账单
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
}

var poolUserInvoiceItemDto = sync.Pool{
	New: func() any {
		return new(UserInvoiceItemDto)
	},
}

// GetUserInvoiceItemDto() 从对象池中获取UserInvoiceItemDto
func GetUserInvoiceItemDto() *UserInvoiceItemDto {
	return poolUserInvoiceItemDto.Get().(*UserInvoiceItemDto)
}

// ReleaseUserInvoiceItemDto 释放UserInvoiceItemDto
func ReleaseUserInvoiceItemDto(v *UserInvoiceItemDto) {
	v.ItemName = ""
	v.Amount = ""
	v.BillDate = ""
	v.ItemType = 0
	poolUserInvoiceItemDto.Put(v)
}
