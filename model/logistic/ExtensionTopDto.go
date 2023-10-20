package logistic

import (
	"sync"
)

// ExtensionTopDto 结构体
type ExtensionTopDto struct {
	// Whether a seller has added invoice to transaction order. If a seller doesn&#39;t provide invoice, Correios will be the only available shipment option. This is mainly to remind sellers of adding invoice before shipment
	InvoiceAdded bool `json:"invoice_added,omitempty" xml:"invoice_added,omitempty"`
}

var poolExtensionTopDto = sync.Pool{
	New: func() any {
		return new(ExtensionTopDto)
	},
}

// GetExtensionTopDto() 从对象池中获取ExtensionTopDto
func GetExtensionTopDto() *ExtensionTopDto {
	return poolExtensionTopDto.Get().(*ExtensionTopDto)
}

// ReleaseExtensionTopDto 释放ExtensionTopDto
func ReleaseExtensionTopDto(v *ExtensionTopDto) {
	v.InvoiceAdded = false
	poolExtensionTopDto.Put(v)
}
