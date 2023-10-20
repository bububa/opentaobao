package promotion

import (
	"sync"
)

// WdkCouponApplyParam 结构体
type WdkCouponApplyParam struct {
	// 申请渠道
	ApplyChannel string `json:"apply_channel,omitempty" xml:"apply_channel,omitempty"`
	// 券领取来源，非必填，用于区分不同的业务来源，默认传空为领取普通优惠券
	ApplySource string `json:"apply_source,omitempty" xml:"apply_source,omitempty"`
	// 卡券来源  [ump:ump卡券; koubei:口碑券]
	CouponSource string `json:"coupon_source,omitempty" xml:"coupon_source,omitempty"`
	// 拓展属性
	Features string `json:"features,omitempty" xml:"features,omitempty"`
	// 外部业务id
	OutBizNo string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty"`
	// 优惠券模版id
	TemplateId string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 优惠券模版uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 幂等键
	IdempotentKey string `json:"idempotent_key,omitempty" xml:"idempotent_key,omitempty"`
	// 领取淘系id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolWdkCouponApplyParam = sync.Pool{
	New: func() any {
		return new(WdkCouponApplyParam)
	},
}

// GetWdkCouponApplyParam() 从对象池中获取WdkCouponApplyParam
func GetWdkCouponApplyParam() *WdkCouponApplyParam {
	return poolWdkCouponApplyParam.Get().(*WdkCouponApplyParam)
}

// ReleaseWdkCouponApplyParam 释放WdkCouponApplyParam
func ReleaseWdkCouponApplyParam(v *WdkCouponApplyParam) {
	v.ApplyChannel = ""
	v.ApplySource = ""
	v.CouponSource = ""
	v.Features = ""
	v.OutBizNo = ""
	v.TemplateId = ""
	v.Uuid = ""
	v.IdempotentKey = ""
	v.UserId = 0
	poolWdkCouponApplyParam.Put(v)
}
