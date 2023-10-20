package ascpchannel

import (
	"sync"
)

// Refunderinfo 结构体
type Refunderinfo struct {
	// 国家
	RefunderCountry string `json:"refunder_country,omitempty" xml:"refunder_country,omitempty"`
	// 省
	RefunderProvince string `json:"refunder_province,omitempty" xml:"refunder_province,omitempty"`
	// 市
	RefunderCity string `json:"refunder_city,omitempty" xml:"refunder_city,omitempty"`
	// 区
	RefunderArea string `json:"refunder_area,omitempty" xml:"refunder_area,omitempty"`
	// 村镇
	RefunderTown string `json:"refunder_town,omitempty" xml:"refunder_town,omitempty"`
	// 详细地址
	RefunderDetailAddress string `json:"refunder_detail_address,omitempty" xml:"refunder_detail_address,omitempty"`
	// 退货人
	RefunderName string `json:"refunder_name,omitempty" xml:"refunder_name,omitempty"`
	// 手机号
	RefunderMobile string `json:"refunder_mobile,omitempty" xml:"refunder_mobile,omitempty"`
	// 固定电话
	RefunderPhone string `json:"refunder_phone,omitempty" xml:"refunder_phone,omitempty"`
	// 邮编
	RefunderZipCode string `json:"refunder_zip_code,omitempty" xml:"refunder_zip_code,omitempty"`
}

var poolRefunderinfo = sync.Pool{
	New: func() any {
		return new(Refunderinfo)
	},
}

// GetRefunderinfo() 从对象池中获取Refunderinfo
func GetRefunderinfo() *Refunderinfo {
	return poolRefunderinfo.Get().(*Refunderinfo)
}

// ReleaseRefunderinfo 释放Refunderinfo
func ReleaseRefunderinfo(v *Refunderinfo) {
	v.RefunderCountry = ""
	v.RefunderProvince = ""
	v.RefunderCity = ""
	v.RefunderArea = ""
	v.RefunderTown = ""
	v.RefunderDetailAddress = ""
	v.RefunderName = ""
	v.RefunderMobile = ""
	v.RefunderPhone = ""
	v.RefunderZipCode = ""
	poolRefunderinfo.Put(v)
}
