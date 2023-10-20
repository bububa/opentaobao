package campus

import (
	"sync"
)

// VoucherDto 结构体
type VoucherDto struct {
	// 卡号
	VoucherNo string `json:"voucher_no,omitempty" xml:"voucher_no,omitempty"`
	// xxx
	VoucherTypeEnum string `json:"voucher_type_enum,omitempty" xml:"voucher_type_enum,omitempty"`
}

var poolVoucherDto = sync.Pool{
	New: func() any {
		return new(VoucherDto)
	},
}

// GetVoucherDto() 从对象池中获取VoucherDto
func GetVoucherDto() *VoucherDto {
	return poolVoucherDto.Get().(*VoucherDto)
}

// ReleaseVoucherDto 释放VoucherDto
func ReleaseVoucherDto(v *VoucherDto) {
	v.VoucherNo = ""
	v.VoucherTypeEnum = ""
	poolVoucherDto.Put(v)
}
