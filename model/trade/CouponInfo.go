package trade

// CouponInfo 
type CouponInfo struct {

    // 优惠名称
    Name   string `json:"name,omitempty"`

    // 优惠金额，单位人民币：分
    Discount   int64 `json:"discount,omitempty"`

    // 优惠标识，编号
    OptionId   string `json:"option_id,omitempty"`

}
