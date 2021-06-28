package promotion

// CouponTemplateCommonConfig 
type CouponTemplateCommonConfig struct {

    // 申请渠道 anonymousOffline
    
    ApplyChannels   []string `json:"apply_channels,omitempty" xml:"apply_channels>string,omitempty"`
    

    // 优惠券活动描述
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // 模板状态 NORMAL(1, "有效"), DELETE(-1, "删除"), ENDING(0, "结束领取"), NOUSE(-2, "无效"), WDK_COUPON_DRAFT(-999, "草稿"),
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 优惠券名称
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 优惠券类型 UNIT_PRICE(10,"unitPrice","单品定价券"), FULL_AMOUNT_REDUCE(11, "fullAmountReduce", "满元减券"), FULL_AMOUNT_DISCOUNT(12, "fullAmountDiscount", "满元折券"), FULL_COUNT_REDUCE(13, "fullCountReduce", "满件减券"), FULL_COUNT_DISCOUNT(14, "fullCountDiscount", "满件折券"), VOUCHER(15, "voucher", "抵用券"),
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

    // 优惠券发放方式 ANONYMOUS("anonymous","匿名券"), REGISTERED("registered","记名券"),
    
    SendType   string `json:"send_type,omitempty" xml:"send_type,omitempty"`
    

}
