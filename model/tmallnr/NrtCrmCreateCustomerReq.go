package tmallnr

import (
	"sync"
)

// NrtCrmCreateCustomerReq 结构体
type NrtCrmCreateCustomerReq struct {
	// 手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 业务身份：居然/红星
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 会员openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 微信昵称
	WechatName string `json:"wechat_name,omitempty" xml:"wechat_name,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 摊位id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 导购员id
	GuiderId int64 `json:"guider_id,omitempty" xml:"guider_id,omitempty"`
}

var poolNrtCrmCreateCustomerReq = sync.Pool{
	New: func() any {
		return new(NrtCrmCreateCustomerReq)
	},
}

// GetNrtCrmCreateCustomerReq() 从对象池中获取NrtCrmCreateCustomerReq
func GetNrtCrmCreateCustomerReq() *NrtCrmCreateCustomerReq {
	return poolNrtCrmCreateCustomerReq.Get().(*NrtCrmCreateCustomerReq)
}

// ReleaseNrtCrmCreateCustomerReq 释放NrtCrmCreateCustomerReq
func ReleaseNrtCrmCreateCustomerReq(v *NrtCrmCreateCustomerReq) {
	v.Phone = ""
	v.BizCode = ""
	v.OpenId = ""
	v.Name = ""
	v.WechatName = ""
	v.ActivityId = 0
	v.StoreId = 0
	v.GuiderId = 0
	poolNrtCrmCreateCustomerReq.Put(v)
}
