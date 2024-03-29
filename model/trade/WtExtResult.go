package trade

import (
	"sync"
)

// WtExtResult 结构体
type WtExtResult struct {
	// 用户手机号码
	PhoneNum string `json:"phone_num,omitempty" xml:"phone_num,omitempty"`
	// 手机号码所在城市的区位码
	PhoneCityCode string `json:"phone_city_code,omitempty" xml:"phone_city_code,omitempty"`
	// 手机号码所在省份的区位码
	PhoneProvinceCode string `json:"phone_province_code,omitempty" xml:"phone_province_code,omitempty"`
	// 套餐名称
	PlanTitle string `json:"plan_title,omitempty" xml:"plan_title,omitempty"`
	// 套餐商家编码
	OutPlanId string `json:"out_plan_id,omitempty" xml:"out_plan_id,omitempty"`
	// 合约计划商家编码
	OutPackageId string `json:"out_package_id,omitempty" xml:"out_package_id,omitempty"`
	// 机主姓名
	PhoneOwnerName string `json:"phone_owner_name,omitempty" xml:"phone_owner_name,omitempty"`
	// 证件号
	CertCardNum string `json:"cert_card_num,omitempty" xml:"cert_card_num,omitempty"`
	// 实人认证方式，如：V1,V2，默认V1
	AuthType string `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
	// 预留属性字符串
	Attr string `json:"attr,omitempty" xml:"attr,omitempty"`
	// 安装地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 联系人
	OwnerName string `json:"owner_name,omitempty" xml:"owner_name,omitempty"`
	// 宽带账号
	Account string `json:"account,omitempty" xml:"account,omitempty"`
	// 描述信息
	PromotionDesc string `json:"promotion_desc,omitempty" xml:"promotion_desc,omitempty"`
	// 活体流水id
	BiometricSeq string `json:"biometric_seq,omitempty" xml:"biometric_seq,omitempty"`
	// 交易id
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 套餐开通规则
	EffectRule int64 `json:"effect_rule,omitempty" xml:"effect_rule,omitempty"`
	// 协议商家编码
	AgreementId int64 `json:"agreement_id,omitempty" xml:"agreement_id,omitempty"`
	// 证件类型
	CertType int64 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// 号码预存款(单位是分)
	PhoneDeposit int64 `json:"phone_deposit,omitempty" xml:"phone_deposit,omitempty"`
	// 减免 号码预存款(单位是分)
	PhoneFreeDeposit int64 `json:"phone_free_deposit,omitempty" xml:"phone_free_deposit,omitempty"`
	// 合约用户类型，0=老用户合约机，1=新用户合约机，2=新用户合约号卡，99=其它（默认）
	UserType int64 `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 合约类型，合约类目， 0=机卡合约，1001=购物送，1002=阿里通信号卡合约，1003=信用购，1007=话费充值卡，1005=游戏点卡，1006=QQ点卡，99=未知（默认）
	ContractType int64 `json:"contract_type,omitempty" xml:"contract_type,omitempty"`
}

var poolWtExtResult = sync.Pool{
	New: func() any {
		return new(WtExtResult)
	},
}

// GetWtExtResult() 从对象池中获取WtExtResult
func GetWtExtResult() *WtExtResult {
	return poolWtExtResult.Get().(*WtExtResult)
}

// ReleaseWtExtResult 释放WtExtResult
func ReleaseWtExtResult(v *WtExtResult) {
	v.PhoneNum = ""
	v.PhoneCityCode = ""
	v.PhoneProvinceCode = ""
	v.PlanTitle = ""
	v.OutPlanId = ""
	v.OutPackageId = ""
	v.PhoneOwnerName = ""
	v.CertCardNum = ""
	v.AuthType = ""
	v.Attr = ""
	v.Address = ""
	v.OwnerName = ""
	v.Account = ""
	v.PromotionDesc = ""
	v.BiometricSeq = ""
	v.Tid = 0
	v.EffectRule = 0
	v.AgreementId = 0
	v.CertType = 0
	v.PhoneDeposit = 0
	v.PhoneFreeDeposit = 0
	v.UserType = 0
	v.ContractType = 0
	poolWtExtResult.Put(v)
}
