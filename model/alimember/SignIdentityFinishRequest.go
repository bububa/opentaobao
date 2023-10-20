package alimember

import (
	"sync"
)

// SignIdentityFinishRequest 结构体
type SignIdentityFinishRequest struct {
	// 付费会员模板id
	IdentityTemplateId string `json:"identity_template_id,omitempty" xml:"identity_template_id,omitempty"`
	// 外部会员ID，识别商家会员的唯一身份标识
	OuterMemberId string `json:"outer_member_id,omitempty" xml:"outer_member_id,omitempty"`
	// 对外开放的merchantId
	OpenMerchantId string `json:"open_merchant_id,omitempty" xml:"open_merchant_id,omitempty"`
	// 时间戳
	TimeStamp int64 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// 付费会员时间信息
	IdentityModel *IdentityModel `json:"identity_model,omitempty" xml:"identity_model,omitempty"`
	// 签约单据信息
	OrderModel *OrderModel `json:"order_model,omitempty" xml:"order_model,omitempty"`
	// 处理结果，成功还是失败
	SendSuc bool `json:"send_suc,omitempty" xml:"send_suc,omitempty"`
}

var poolSignIdentityFinishRequest = sync.Pool{
	New: func() any {
		return new(SignIdentityFinishRequest)
	},
}

// GetSignIdentityFinishRequest() 从对象池中获取SignIdentityFinishRequest
func GetSignIdentityFinishRequest() *SignIdentityFinishRequest {
	return poolSignIdentityFinishRequest.Get().(*SignIdentityFinishRequest)
}

// ReleaseSignIdentityFinishRequest 释放SignIdentityFinishRequest
func ReleaseSignIdentityFinishRequest(v *SignIdentityFinishRequest) {
	v.IdentityTemplateId = ""
	v.OuterMemberId = ""
	v.OpenMerchantId = ""
	v.TimeStamp = 0
	v.IdentityModel = nil
	v.OrderModel = nil
	v.SendSuc = false
	poolSignIdentityFinishRequest.Put(v)
}
