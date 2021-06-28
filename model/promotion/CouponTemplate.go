package promotion

// CouponTemplate 
type CouponTemplate struct {

    // 领取限制
    
    ApplyLimitConfig   *CouponTemplateApplyLimitConfig `json:"apply_limit_config,omitempty" xml:"apply_limit_config,omitempty"`
    

    // 基础信息
    
    CommonConfig   *CouponTemplateCommonConfig `json:"common_config,omitempty" xml:"common_config,omitempty"`
    

    // 生效条件
    
    ConditionConfig   *CouponTemplateConditionConfig `json:"condition_config,omitempty" xml:"condition_config,omitempty"`
    

    // 优惠效果
    
    DiscountConfig   *CouponTemplateDiscountConfig `json:"discount_config,omitempty" xml:"discount_config,omitempty"`
    

    // 模板表主键 创建时为空
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 出资人配置
    
    InvestmentConfig   *CouponTemplateInvestmentConfig `json:"investment_config,omitempty" xml:"investment_config,omitempty"`
    

    // 其他配置--拓展信息
    
    OptionConfig   *CouponTemplateOptionConfig `json:"option_config,omitempty" xml:"option_config,omitempty"`
    

    // 参与者
    
    ParticipateConfig   *CouponTemplateParticipateConfig `json:"participate_config,omitempty" xml:"participate_config,omitempty"`
    

    // ump模板ID
    
    SourceId   int64 `json:"source_id,omitempty" xml:"source_id,omitempty"`
    

    // 实例有效时间配置
    
    TimeLimitConfig   *CouponTemplateTimeLimitConfig `json:"time_limit_config,omitempty" xml:"time_limit_config,omitempty"`
    

    // 优惠券模版uuid
    
    Uuid   string `json:"uuid,omitempty" xml:"uuid,omitempty"`
    

    // 兼容历史逻辑配置
    
    CompatibleConfig   *CouponTemplateCompatibleConfig `json:"compatible_config,omitempty" xml:"compatible_config,omitempty"`
    

    // 幂等id，外部透传
    
    UniqueId   string `json:"unique_id,omitempty" xml:"unique_id,omitempty"`
    

}
