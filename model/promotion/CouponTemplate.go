package promotion

// CouponTemplate 
/* model for simplify = false
type CouponTemplate struct {

    // 领取限制
    
    ApplyLimitConfig  *struct {
        CouponTemplateApplyLimitConfig  *CouponTemplateApplyLimitConfig `json:"coupon_template_apply_limit_config,omitempty"`
    } `json:"apply_limit_config,omitempty"`
    

    // 基础信息
    
    CommonConfig  *struct {
        CouponTemplateCommonConfig  *CouponTemplateCommonConfig `json:"coupon_template_common_config,omitempty"`
    } `json:"common_config,omitempty"`
    

    // 生效条件
    
    ConditionConfig  *struct {
        CouponTemplateConditionConfig  *CouponTemplateConditionConfig `json:"coupon_template_condition_config,omitempty"`
    } `json:"condition_config,omitempty"`
    

    // 优惠效果
    
    DiscountConfig  *struct {
        CouponTemplateDiscountConfig  *CouponTemplateDiscountConfig `json:"coupon_template_discount_config,omitempty"`
    } `json:"discount_config,omitempty"`
    

    // 模板表主键 创建时为空
    
    Id   int64 `json:"id,omitempty"`
    

    // 出资人配置
    
    InvestmentConfig  *struct {
        CouponTemplateInvestmentConfig  *CouponTemplateInvestmentConfig `json:"coupon_template_investment_config,omitempty"`
    } `json:"investment_config,omitempty"`
    

    // 其他配置--拓展信息
    
    OptionConfig  *struct {
        CouponTemplateOptionConfig  *CouponTemplateOptionConfig `json:"coupon_template_option_config,omitempty"`
    } `json:"option_config,omitempty"`
    

    // 参与者
    
    ParticipateConfig  *struct {
        CouponTemplateParticipateConfig  *CouponTemplateParticipateConfig `json:"coupon_template_participate_config,omitempty"`
    } `json:"participate_config,omitempty"`
    

    // ump模板ID
    
    SourceId   int64 `json:"source_id,omitempty"`
    

    // 实例有效时间配置
    
    TimeLimitConfig  *struct {
        CouponTemplateTimeLimitConfig  *CouponTemplateTimeLimitConfig `json:"coupon_template_time_limit_config,omitempty"`
    } `json:"time_limit_config,omitempty"`
    

    // 优惠券模版uuid
    
    Uuid   string `json:"uuid,omitempty"`
    

    // 兼容历史逻辑配置
    
    CompatibleConfig  *struct {
        CouponTemplateCompatibleConfig  *CouponTemplateCompatibleConfig `json:"coupon_template_compatible_config,omitempty"`
    } `json:"compatible_config,omitempty"`
    

    // 幂等id，外部透传
    
    UniqueId   string `json:"unique_id,omitempty"`
    

}
*/

// CouponTemplate 
type CouponTemplate struct {

    // 领取限制
    ApplyLimitConfig   *CouponTemplateApplyLimitConfig `json:"apply_limit_config,omitempty"`

    // 基础信息
    CommonConfig   *CouponTemplateCommonConfig `json:"common_config,omitempty"`

    // 生效条件
    ConditionConfig   *CouponTemplateConditionConfig `json:"condition_config,omitempty"`

    // 优惠效果
    DiscountConfig   *CouponTemplateDiscountConfig `json:"discount_config,omitempty"`

    // 模板表主键 创建时为空
    Id   int64 `json:"id,omitempty"`

    // 出资人配置
    InvestmentConfig   *CouponTemplateInvestmentConfig `json:"investment_config,omitempty"`

    // 其他配置--拓展信息
    OptionConfig   *CouponTemplateOptionConfig `json:"option_config,omitempty"`

    // 参与者
    ParticipateConfig   *CouponTemplateParticipateConfig `json:"participate_config,omitempty"`

    // ump模板ID
    SourceId   int64 `json:"source_id,omitempty"`

    // 实例有效时间配置
    TimeLimitConfig   *CouponTemplateTimeLimitConfig `json:"time_limit_config,omitempty"`

    // 优惠券模版uuid
    Uuid   string `json:"uuid,omitempty"`

    // 兼容历史逻辑配置
    CompatibleConfig   *CouponTemplateCompatibleConfig `json:"compatible_config,omitempty"`

    // 幂等id，外部透传
    UniqueId   string `json:"unique_id,omitempty"`

}
