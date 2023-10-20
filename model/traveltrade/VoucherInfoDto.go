package traveltrade

import (
	"sync"
)

// VoucherInfoDto 结构体
type VoucherInfoDto struct {
	// 用户短信会收到的确认号
	ConfirmCode string `json:"confirm_code,omitempty" xml:"confirm_code,omitempty"`
	// 凭证使用时间，格式:yyyy-MM-dd HH:mm:ss
	UsedDate string `json:"used_date,omitempty" xml:"used_date,omitempty"`
	// 凭证使用次数
	UsedQuantity int64 `json:"used_quantity,omitempty" xml:"used_quantity,omitempty"`
}

var poolVoucherInfoDto = sync.Pool{
	New: func() any {
		return new(VoucherInfoDto)
	},
}

// GetVoucherInfoDto() 从对象池中获取VoucherInfoDto
func GetVoucherInfoDto() *VoucherInfoDto {
	return poolVoucherInfoDto.Get().(*VoucherInfoDto)
}

// ReleaseVoucherInfoDto 释放VoucherInfoDto
func ReleaseVoucherInfoDto(v *VoucherInfoDto) {
	v.ConfirmCode = ""
	v.UsedDate = ""
	v.UsedQuantity = 0
	poolVoucherInfoDto.Put(v)
}
