package alicom

// PhoneRecommendRes 结构体
type PhoneRecommendRes struct {
	// 商品列表
	List []PhoneDistributionPhoneItemVo `json:"list,omitempty" xml:"list>phone_distribution_phone_item_vo,omitempty"`
	// 公告
	Notice string `json:"notice,omitempty" xml:"notice,omitempty"`
	// 号码信息
	CatInfo *CatInfoVo `json:"cat_info,omitempty" xml:"cat_info,omitempty"`
}
