package usergrowth2

// TaobaoUsergrowthDeliveryAskData 结构体
type TaobaoUsergrowthDeliveryAskData struct {
	// 点击跳转h5链接
	H5Url string `json:"h5_url,omitempty" xml:"h5_url,omitempty"`
	// 图片链接
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 优惠信息描述
	Coupon string `json:"coupon,omitempty" xml:"coupon,omitempty"`
	// 出价信息
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 点击拉端deeplink链接
	DeeplinkUrl string `json:"deeplink_url,omitempty" xml:"deeplink_url,omitempty"`
	// 智能出价信息
	SmartBid string `json:"smart_bid,omitempty" xml:"smart_bid,omitempty"`
	// 追踪id, 需要将其带到曝光反馈和点击反馈接口中
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
}
