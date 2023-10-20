package alimember

import (
	"sync"
)

// SaasCustomerInfo 结构体
type SaasCustomerInfo struct {
	// 本地生活会员id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// yyyy-MM-dd HH:mm:ss 生日
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// MALE-男 FEMALE-女
	Sex string `json:"sex,omitempty" xml:"sex,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 积分
	Point string `json:"point,omitempty" xml:"point,omitempty"`
	// 入会门店
	MemberShopId string `json:"member_shop_id,omitempty" xml:"member_shop_id,omitempty"`
	// yyyy-MM-dd HH:mm:ss 入会时间
	MemberTime string `json:"member_time,omitempty" xml:"member_time,omitempty"`
	// isv会员id
	IsvMemberId string `json:"isv_member_id,omitempty" xml:"isv_member_id,omitempty"`
}

var poolSaasCustomerInfo = sync.Pool{
	New: func() any {
		return new(SaasCustomerInfo)
	},
}

// GetSaasCustomerInfo() 从对象池中获取SaasCustomerInfo
func GetSaasCustomerInfo() *SaasCustomerInfo {
	return poolSaasCustomerInfo.Get().(*SaasCustomerInfo)
}

// ReleaseSaasCustomerInfo 释放SaasCustomerInfo
func ReleaseSaasCustomerInfo(v *SaasCustomerInfo) {
	v.CustomerId = ""
	v.Birthday = ""
	v.Mobile = ""
	v.Sex = ""
	v.Name = ""
	v.Point = ""
	v.MemberShopId = ""
	v.MemberTime = ""
	v.IsvMemberId = ""
	poolSaasCustomerInfo.Put(v)
}
