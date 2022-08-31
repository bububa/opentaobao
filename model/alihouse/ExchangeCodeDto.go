package alihouse

// ExchangeCodeDto 结构体
type ExchangeCodeDto struct {
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 券码
	CouponCode string `json:"coupon_code,omitempty" xml:"coupon_code,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
