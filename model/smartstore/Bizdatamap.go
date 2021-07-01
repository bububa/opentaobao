package smartstore

// Bizdatamap 结构体
type Bizdatamap struct {
	// 商品ID，item_id 在action为ITEM_CLICK、ITEM_SENSOR、BUY_CLICK必须传入； 必须使用淘宝商品id，否则失败。
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// "例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
	CouponId string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
}
