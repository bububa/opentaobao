package einvoice

import (
	"sync"
)

// InvoiceContactDto 结构体
type InvoiceContactDto struct {
	// 联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人电话
	ContactMobile string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
	// 联系人地址
	ContactAddr string `json:"contact_addr,omitempty" xml:"contact_addr,omitempty"`
	// 联系人邮件
	ContactMail string `json:"contact_mail,omitempty" xml:"contact_mail,omitempty"`
}

var poolInvoiceContactDto = sync.Pool{
	New: func() any {
		return new(InvoiceContactDto)
	},
}

// GetInvoiceContactDto() 从对象池中获取InvoiceContactDto
func GetInvoiceContactDto() *InvoiceContactDto {
	return poolInvoiceContactDto.Get().(*InvoiceContactDto)
}

// ReleaseInvoiceContactDto 释放InvoiceContactDto
func ReleaseInvoiceContactDto(v *InvoiceContactDto) {
	v.ContactName = ""
	v.ContactMobile = ""
	v.ContactAddr = ""
	v.ContactMail = ""
	poolInvoiceContactDto.Put(v)
}
