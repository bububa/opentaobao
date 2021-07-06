package fenxiao

// DiscountDetail 结构体
type DiscountDetail struct {
	// 折扣类型:GRADE（按会员等级优惠）、DISTRIBUTOR（按分销商优惠）
	TargetType string `json:"target_type,omitempty" xml:"target_type,omitempty"`
	// 等级名称或者分销商名称
	TargetName string `json:"target_name,omitempty" xml:"target_name,omitempty"`
	// 优惠方式:PERCENT（按折扣优惠）、PRICE（按减价优惠）
	DiscountType string `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	// 创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 折扣详情ID
	DetailId int64 `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
	// 会员等级的id或者分销商id
	TargetId int64 `json:"target_id,omitempty" xml:"target_id,omitempty"`
	// 优惠比率或者优惠价格 10%或10
	DiscountValue int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty"`
}
