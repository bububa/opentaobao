package trade

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
	// authType
	AuthType string `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
	// attr
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
	// userType
	UserType int64 `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// contractType
	ContractType int64 `json:"contract_type,omitempty" xml:"contract_type,omitempty"`
}
