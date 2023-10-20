package tmallnr

import (
	"sync"
)

// CrmAppointActivityReq 结构体
type CrmAppointActivityReq struct {
	// 微信名称
	WechatName string `json:"wechat_name,omitempty" xml:"wechat_name,omitempty"`
	// 业务code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 会员ID
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 导购员ID
	GuiderId int64 `json:"guider_id,omitempty" xml:"guider_id,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

var poolCrmAppointActivityReq = sync.Pool{
	New: func() any {
		return new(CrmAppointActivityReq)
	},
}

// GetCrmAppointActivityReq() 从对象池中获取CrmAppointActivityReq
func GetCrmAppointActivityReq() *CrmAppointActivityReq {
	return poolCrmAppointActivityReq.Get().(*CrmAppointActivityReq)
}

// ReleaseCrmAppointActivityReq 释放CrmAppointActivityReq
func ReleaseCrmAppointActivityReq(v *CrmAppointActivityReq) {
	v.WechatName = ""
	v.BizCode = ""
	v.OpenId = ""
	v.GuiderId = 0
	v.StoreId = 0
	v.ActivityId = 0
	poolCrmAppointActivityReq.Put(v)
}
