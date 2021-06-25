package alsc

// FeeDetail 
type FeeDetail struct {

    // 金额
    Amount   int64 `json:"amount,omitempty"`

    // 计算方式 ：  加-PLUS   加-MINUS
    Operator   string `json:"operator,omitempty"`

    // MERCHANT_ADJUST  商家调整  PLATFORM_PROM  平台优惠  MERCHANT_PROM  商家优惠  CHARGE  储值  POINT  积分  OTHER  其他
    Type   string `json:"type,omitempty"`

}
