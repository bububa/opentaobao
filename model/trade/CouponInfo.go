package trade

// CouponInfo 结构体
type CouponInfo struct {
	// 优惠名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 优惠金额，单位人民币：分
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 优惠标识，编号
	OptionId string `json:"option_id,omitempty" xml:"option_id,omitempty"`
}
