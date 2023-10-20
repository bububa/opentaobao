package tbk

// TaobaotbkdgoptimuspromotionMapData 结构体
type TaobaotbkdgoptimuspromotionMapData struct {
	// 权益信息
	Promotionlist []PromotionList `json:"promotion_list,omitempty" xml:"promotion_list>promotion_list,omitempty"`
	// 权益类型。1 有价券（需要购买的店铺券或单品券） 2 公开券（直接领取的店铺券或单品券） 3 妈妈券（妈妈渠道领取的店铺券或单品券） 4.品类券 （跨店可用券，可与1，2，3叠加）
	Promotiontype string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 优惠门槛类型： 1 满元 2 满件
	Conditiontype string `json:"condition_type,omitempty" xml:"condition_type,omitempty"`
	// 优惠类型： 1 减钱 2 打折
	Discounttype string `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	// 权益信息展示开始时间，精确到毫秒时间戳
	Displaystarttime string `json:"display_start_time,omitempty" xml:"display_start_time,omitempty"`
	// 权益信息展示结束时间，精确到毫秒时间戳
	Displayendtime string `json:"display_end_time,omitempty" xml:"display_end_time,omitempty"`
	// 店铺信息-卖家ID
	Sellerid string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 店铺信息-卖家昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 店铺信息-店铺名称
	Shoptitle string `json:"shop_title,omitempty" xml:"shop_title,omitempty"`
	// 店铺信息-店铺logo
	Shoppictureurl string `json:"shop_picture_url,omitempty" xml:"shop_picture_url,omitempty"`
	// 权益信息-总量（权益初始库存量）
	Totalcount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 权益信息-剩余库存（权益剩余库存量）
	Remaincount int64 `json:"remain_count,omitempty" xml:"remain_count,omitempty"`
	// 权益扩展信息
	Promotionextend *PromotionExtend `json:"promotion_extend,omitempty" xml:"promotion_extend,omitempty"`
}
