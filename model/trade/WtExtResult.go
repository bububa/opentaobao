package trade

// WtExtResult 
/* model for simplify = false
type WtExtResult struct {

    // 交易id
    
    Tid   int64 `json:"tid,omitempty"`
    

    // 用户手机号码
    
    PhoneNum   string `json:"phone_num,omitempty"`
    

    // 手机号码所在城市的区位码
    
    PhoneCityCode   string `json:"phone_city_code,omitempty"`
    

    // 手机号码所在省份的区位码
    
    PhoneProvinceCode   string `json:"phone_province_code,omitempty"`
    

    // 套餐名称
    
    PlanTitle   string `json:"plan_title,omitempty"`
    

    // 套餐商家编码
    
    OutPlanId   string `json:"out_plan_id,omitempty"`
    

    // 合约计划商家编码
    
    OutPackageId   string `json:"out_package_id,omitempty"`
    

    // 套餐开通规则
    
    EffectRule   int64 `json:"effect_rule,omitempty"`
    

    // 协议商家编码
    
    AgreementId   int64 `json:"agreement_id,omitempty"`
    

    // 机主姓名
    
    PhoneOwnerName   string `json:"phone_owner_name,omitempty"`
    

    // 证件类型
    
    CertType   int64 `json:"cert_type,omitempty"`
    

    // 证件号
    
    CertCardNum   string `json:"cert_card_num,omitempty"`
    

    // 号码预存款(单位是分)
    
    PhoneDeposit   int64 `json:"phone_deposit,omitempty"`
    

    // 减免 号码预存款(单位是分)
    
    PhoneFreeDeposit   int64 `json:"phone_free_deposit,omitempty"`
    

    // authType
    
    AuthType   string `json:"auth_type,omitempty"`
    

    // userType
    
    UserType   int64 `json:"user_type,omitempty"`
    

    // contractType
    
    ContractType   int64 `json:"contract_type,omitempty"`
    

    // attr
    
    Attr   string `json:"attr,omitempty"`
    

    // 安装地址
    
    Address   string `json:"address,omitempty"`
    

    // 联系人
    
    OwnerName   string `json:"owner_name,omitempty"`
    

    // 宽带账号
    
    Account   string `json:"account,omitempty"`
    

    // 描述信息
    
    PromotionDesc   string `json:"promotion_desc,omitempty"`
    

    // 活体流水id
    
    BiometricSeq   string `json:"biometric_seq,omitempty"`
    

}
*/

// WtExtResult 
type WtExtResult struct {

    // 交易id
    Tid   int64 `json:"tid,omitempty"`

    // 用户手机号码
    PhoneNum   string `json:"phone_num,omitempty"`

    // 手机号码所在城市的区位码
    PhoneCityCode   string `json:"phone_city_code,omitempty"`

    // 手机号码所在省份的区位码
    PhoneProvinceCode   string `json:"phone_province_code,omitempty"`

    // 套餐名称
    PlanTitle   string `json:"plan_title,omitempty"`

    // 套餐商家编码
    OutPlanId   string `json:"out_plan_id,omitempty"`

    // 合约计划商家编码
    OutPackageId   string `json:"out_package_id,omitempty"`

    // 套餐开通规则
    EffectRule   int64 `json:"effect_rule,omitempty"`

    // 协议商家编码
    AgreementId   int64 `json:"agreement_id,omitempty"`

    // 机主姓名
    PhoneOwnerName   string `json:"phone_owner_name,omitempty"`

    // 证件类型
    CertType   int64 `json:"cert_type,omitempty"`

    // 证件号
    CertCardNum   string `json:"cert_card_num,omitempty"`

    // 号码预存款(单位是分)
    PhoneDeposit   int64 `json:"phone_deposit,omitempty"`

    // 减免 号码预存款(单位是分)
    PhoneFreeDeposit   int64 `json:"phone_free_deposit,omitempty"`

    // authType
    AuthType   string `json:"auth_type,omitempty"`

    // userType
    UserType   int64 `json:"user_type,omitempty"`

    // contractType
    ContractType   int64 `json:"contract_type,omitempty"`

    // attr
    Attr   string `json:"attr,omitempty"`

    // 安装地址
    Address   string `json:"address,omitempty"`

    // 联系人
    OwnerName   string `json:"owner_name,omitempty"`

    // 宽带账号
    Account   string `json:"account,omitempty"`

    // 描述信息
    PromotionDesc   string `json:"promotion_desc,omitempty"`

    // 活体流水id
    BiometricSeq   string `json:"biometric_seq,omitempty"`

}
