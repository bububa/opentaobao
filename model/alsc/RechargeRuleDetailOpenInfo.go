package alsc

// RechargeRuleDetailOpenInfo 
type RechargeRuleDetailOpenInfo struct {
    // 赠送积分
    GiftPoint   int64 `json:"gift_point,omitempty" xml:"gift_point,omitempty"`
    // 赠送金额
    GiftValue   int64 `json:"gift_value,omitempty" xml:"gift_value,omitempty"`
    // 储值金额
    RealValue   int64 `json:"real_value,omitempty" xml:"real_value,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
}
