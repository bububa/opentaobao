package tmallitem

// TmallExtendSearchItem 结构体
type TmallExtendSearchItem struct {
	// 搜索宝贝的标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 搜索宝贝的月数量
	Sold string `json:"sold,omitempty" xml:"sold,omitempty"`
	// 搜索宝贝url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 搜索宝贝的图片url
	PicPath string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
	// 搜索宝贝的卖家昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 搜索宝贝的宝贝所在地
	Location string `json:"location,omitempty" xml:"location,omitempty"`
	// 搜索宝贝的卖家所在地
	SellerLoc string `json:"seller_loc,omitempty" xml:"seller_loc,omitempty"`
	// 搜索宝贝的sku最低价
	Price float64 `json:"price,omitempty" xml:"price,omitempty"`
	// 搜索宝贝的数字id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 是否免邮
	Shipping int64 `json:"shipping,omitempty" xml:"shipping,omitempty"`
	// 搜索宝贝的spuid
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// 是否货到付款
	IsCod int64 `json:"is_cod,omitempty" xml:"is_cod,omitempty"`
	// 邮费
	FastPostFee float64 `json:"fast_post_fee,omitempty" xml:"fast_post_fee,omitempty"`
	// 搜索宝贝的卖家数字id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 搜索宝贝的sku最低价的折扣价
	PriceWithRate float64 `json:"price_with_rate,omitempty" xml:"price_with_rate,omitempty"`
	// 是否折扣
	IsPromotion bool `json:"is_promotion,omitempty" xml:"is_promotion,omitempty"`
}
