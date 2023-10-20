package fenxiao

import (
	"sync"
)

// DpBuyerDetail 结构体
type DpBuyerDetail struct {
	// 地区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 收货地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 移动电话
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 邮编
	Post string `json:"post,omitempty" xml:"post,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 固定电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 收货人全名
	FullName string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// 省份
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
	// 收件人ID (Open Addressee ID)，长度在128个字符之内。
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
}

var poolDpBuyerDetail = sync.Pool{
	New: func() any {
		return new(DpBuyerDetail)
	},
}

// GetDpBuyerDetail() 从对象池中获取DpBuyerDetail
func GetDpBuyerDetail() *DpBuyerDetail {
	return poolDpBuyerDetail.Get().(*DpBuyerDetail)
}

// ReleaseDpBuyerDetail 释放DpBuyerDetail
func ReleaseDpBuyerDetail(v *DpBuyerDetail) {
	v.Area = ""
	v.Address = ""
	v.MobilePhone = ""
	v.Post = ""
	v.City = ""
	v.Phone = ""
	v.FullName = ""
	v.Prov = ""
	v.Oaid = ""
	poolDpBuyerDetail.Put(v)
}
