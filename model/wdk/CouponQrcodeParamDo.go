package wdk

// CouponQrcodeParamDo 结构体
type CouponQrcodeParamDo struct {
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 导购员ID，需要保证唯一性
	GuideId string `json:"guide_id,omitempty" xml:"guide_id,omitempty"`
	// 推广开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 推广结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}
