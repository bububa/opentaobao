package alitripmerchant

import (
	"sync"
)

// DerbyVoucherCountVo 结构体
type DerbyVoucherCountVo struct {
	// 1
	VoucherIds []string `json:"voucher_ids,omitempty" xml:"voucher_ids>string,omitempty"`
	// 1
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// aa
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// aa
	TypeUrl string `json:"type_url,omitempty" xml:"type_url,omitempty"`
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
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 1
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 单次或无限次
	UseTimes string `json:"use_times,omitempty" xml:"use_times,omitempty"`
	// 1
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 1
	DiscountOff int64 `json:"discount_off,omitempty" xml:"discount_off,omitempty"`
	// 1
	Onoffline bool `json:"onoffline,omitempty" xml:"onoffline,omitempty"`
}

var poolDerbyVoucherCountVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherCountVo)
	},
}

// GetDerbyVoucherCountVo() 从对象池中获取DerbyVoucherCountVo
func GetDerbyVoucherCountVo() *DerbyVoucherCountVo {
	return poolDerbyVoucherCountVo.Get().(*DerbyVoucherCountVo)
}

// ReleaseDerbyVoucherCountVo 释放DerbyVoucherCountVo
func ReleaseDerbyVoucherCountVo(v *DerbyVoucherCountVo) {
	v.VoucherIds = v.VoucherIds[:0]
	v.Type = ""
	v.TypeName = ""
	v.TypeUrl = ""
	v.Name = ""
	v.Category = ""
	v.Code = ""
	v.VoucherPic = ""
	v.ShortDes = ""
	v.LongDes = ""
	v.EndDate = ""
	v.Status = ""
	v.UseTimes = ""
	v.Count = 0
	v.DiscountOff = 0
	v.Onoffline = false
	poolDerbyVoucherCountVo.Put(v)
}
