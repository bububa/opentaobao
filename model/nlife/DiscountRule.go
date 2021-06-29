package nlife

// DiscountRule 
type DiscountRule struct {
    // 抵扣人民币1分钱所需要的积分数量，比如10积分一分钱
    PointsValue   string `json:"points_value,omitempty" xml:"points_value,omitempty"`
    // 百分比，抵扣上限。15代表积分最多能抵扣订单总额的15%
    DeductionLimit   int64 `json:"deduction_limit,omitempty" xml:"deduction_limit,omitempty"`
    // 金额上限，单位人民币分
    MoneyLimit   int64 `json:"money_limit,omitempty" xml:"money_limit,omitempty"`
}
