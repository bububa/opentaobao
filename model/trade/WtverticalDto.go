package trade

// WtverticalDto 
type WtverticalDto struct {
    // 手机号码
    PhoneNo   string `json:"phone_no,omitempty" xml:"phone_no,omitempty"`
    // 订单号
    BizOrderId   string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    // 城市区号
    PhoneCityCode   string `json:"phone_city_code,omitempty" xml:"phone_city_code,omitempty"`
    // 省区号
    PhoneProvinceCode   string `json:"phone_province_code,omitempty" xml:"phone_province_code,omitempty"`
    // 属性
    Attr   string `json:"attr,omitempty" xml:"attr,omitempty"`
    // 号码归属地城市
    PhoneCityName   string `json:"phone_city_name,omitempty" xml:"phone_city_name,omitempty"`
    // 号码归属地省
    PhoneProvinceName   string `json:"phone_province_name,omitempty" xml:"phone_province_name,omitempty"`
    // 套餐名称
    PlanTitle   string `json:"plan_title,omitempty" xml:"plan_title,omitempty"`
    // 套餐商家编码
    OutPlanId   string `json:"out_plan_id,omitempty" xml:"out_plan_id,omitempty"`
    // 套餐开通规则
    EffectRule   string `json:"effect_rule,omitempty" xml:"effect_rule,omitempty"`
    // 协议商家编码
    AgreementId   string `json:"agreement_id,omitempty" xml:"agreement_id,omitempty"`
}
