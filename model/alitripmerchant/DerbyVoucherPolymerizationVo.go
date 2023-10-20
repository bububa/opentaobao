package alitripmerchant

import (
	"sync"
)

// DerbyVoucherPolymerizationVo 结构体
type DerbyVoucherPolymerizationVo struct {
	// 1
	MemberVoucherID string `json:"member_voucher_i_d,omitempty" xml:"member_voucher_i_d,omitempty"`
	// 1
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 1
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 1
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 1
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 1
	VoucherPic string `json:"voucher_pic,omitempty" xml:"voucher_pic,omitempty"`
	// 1
	ShortDes string `json:"short_des,omitempty" xml:"short_des,omitempty"`
	// 1
	LongDes string `json:"long_des,omitempty" xml:"long_des,omitempty"`
	// 1
	UsedDate string `json:"used_date,omitempty" xml:"used_date,omitempty"`
	// 1
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 1
	QRCodeImage string `json:"q_r_code_image,omitempty" xml:"q_r_code_image,omitempty"`
	// 1
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 1
	VoucherCode string `json:"voucher_code,omitempty" xml:"voucher_code,omitempty"`
	// 单次/无限次
	UseTimes string `json:"use_times,omitempty" xml:"use_times,omitempty"`
	// 1
	DiscountOff int64 `json:"discount_off,omitempty" xml:"discount_off,omitempty"`
	// 1
	Onoffline bool `json:"onoffline,omitempty" xml:"onoffline,omitempty"`
}

var poolDerbyVoucherPolymerizationVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherPolymerizationVo)
	},
}

// GetDerbyVoucherPolymerizationVo() 从对象池中获取DerbyVoucherPolymerizationVo
func GetDerbyVoucherPolymerizationVo() *DerbyVoucherPolymerizationVo {
	return poolDerbyVoucherPolymerizationVo.Get().(*DerbyVoucherPolymerizationVo)
}

// ReleaseDerbyVoucherPolymerizationVo 释放DerbyVoucherPolymerizationVo
func ReleaseDerbyVoucherPolymerizationVo(v *DerbyVoucherPolymerizationVo) {
	v.MemberVoucherID = ""
	v.Type = ""
	v.Name = ""
	v.Category = ""
	v.Code = ""
	v.VoucherPic = ""
	v.ShortDes = ""
	v.LongDes = ""
	v.UsedDate = ""
	v.Status = ""
	v.QRCodeImage = ""
	v.EndDate = ""
	v.VoucherCode = ""
	v.UseTimes = ""
	v.DiscountOff = 0
	v.Onoffline = false
	poolDerbyVoucherPolymerizationVo.Put(v)
}
