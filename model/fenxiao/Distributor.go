package fenxiao

import (
	"sync"
)

// Distributor 结构体
type Distributor struct {
	// 分销商姓名
	DistributorName string `json:"distributor_name,omitempty" xml:"distributor_name,omitempty"`
	// 联系人
	ContactPerson string `json:"contact_person,omitempty" xml:"contact_person,omitempty"`
	// 分销商的手机号
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 分销商的电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 分销商的email
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 分销商的支付宝帐户
	AlipayAccount string `json:"alipay_account,omitempty" xml:"alipay_account,omitempty"`
	// 分销商的网店链接
	ShopWebLink string `json:"shop_web_link,omitempty" xml:"shop_web_link,omitempty"`
	// 分销商的真实姓名，认证姓名
	FullName string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// 分销商卖家的开店时间
	Starts string `json:"starts,omitempty" xml:"starts,omitempty"`
	// 分销商Id
	DistributorId int64 `json:"distributor_id,omitempty" xml:"distributor_id,omitempty"`
	// 分销商店铺主营类目
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 店铺等级
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// 分销商的淘宝卖家评价
	Appraise int64 `json:"appraise,omitempty" xml:"appraise,omitempty"`
}

var poolDistributor = sync.Pool{
	New: func() any {
		return new(Distributor)
	},
}

// GetDistributor() 从对象池中获取Distributor
func GetDistributor() *Distributor {
	return poolDistributor.Get().(*Distributor)
}

// ReleaseDistributor 释放Distributor
func ReleaseDistributor(v *Distributor) {
	v.DistributorName = ""
	v.ContactPerson = ""
	v.MobilePhone = ""
	v.Phone = ""
	v.Email = ""
	v.AlipayAccount = ""
	v.ShopWebLink = ""
	v.FullName = ""
	v.Starts = ""
	v.DistributorId = 0
	v.CategoryId = 0
	v.Level = 0
	v.Appraise = 0
	poolDistributor.Put(v)
}
