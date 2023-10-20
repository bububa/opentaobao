package idle

import (
	"sync"
)

// AgreementPayParam 结构体
type AgreementPayParam struct {
	// 商户订单号（唯一建）
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 代扣计划
	PlanId string `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	// R1:闲鱼回收代扣（V1版本）
	PayBizCode string `json:"pay_biz_code,omitempty" xml:"pay_biz_code,omitempty"`
}

var poolAgreementPayParam = sync.Pool{
	New: func() any {
		return new(AgreementPayParam)
	},
}

// GetAgreementPayParam() 从对象池中获取AgreementPayParam
func GetAgreementPayParam() *AgreementPayParam {
	return poolAgreementPayParam.Get().(*AgreementPayParam)
}

// ReleaseAgreementPayParam 释放AgreementPayParam
func ReleaseAgreementPayParam(v *AgreementPayParam) {
	v.BizOrderId = ""
	v.PlanId = ""
	v.PayBizCode = ""
	poolAgreementPayParam.Put(v)
}
