package alicom

// PhoneDistributionPhoneItemVo 结构体
type PhoneDistributionPhoneItemVo struct {
	// 商品面额
	Face string `json:"face,omitempty" xml:"face,omitempty"`
	// 商品价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品标识
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
