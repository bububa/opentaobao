package wdk

import (
	"sync"
)

// ReceiptDto 结构体
type ReceiptDto struct {
	// 商家/顾客联小票数据
	ReceiptInfo *ReceiptInfoDto `json:"receipt_info,omitempty" xml:"receipt_info,omitempty"`
}

var poolReceiptDto = sync.Pool{
	New: func() any {
		return new(ReceiptDto)
	},
}

// GetReceiptDto() 从对象池中获取ReceiptDto
func GetReceiptDto() *ReceiptDto {
	return poolReceiptDto.Get().(*ReceiptDto)
}

// ReleaseReceiptDto 释放ReceiptDto
func ReleaseReceiptDto(v *ReceiptDto) {
	v.ReceiptInfo = nil
	poolReceiptDto.Put(v)
}
