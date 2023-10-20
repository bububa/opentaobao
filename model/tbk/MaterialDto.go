package tbk

// MaterialDto 结构体
type MaterialDto struct {
	// 优惠券面额
	Couponinfo string `json:"coupon_info,omitempty" xml:"coupon_info,omitempty"`
	// 优惠券结束时间
	Couponendtime string `json:"coupon_end_time,omitempty" xml:"coupon_end_time,omitempty"`
	// 优惠券开始时间
	Couponstarttime string `json:"coupon_start_time,omitempty" xml:"coupon_start_time,omitempty"`
	// 商品优惠券推广链接
	Couponclickurl string `json:"coupon_click_url,omitempty" xml:"coupon_click_url,omitempty"`
	// 当不入参special_id、relation_id、external_id时，展示常规佣金率(%)
	Maxcommissionrate string `json:"max_commission_rate,omitempty" xml:"max_commission_rate,omitempty"`
	// 商品淘客链接
	Itemurl string `json:"item_url,omitempty" xml:"item_url,omitempty"`
	// 预售有礼-推广链接
	Ysylclickurl string `json:"ysyl_click_url,omitempty" xml:"ysyl_click_url,omitempty"`
	// 预售有礼-预估淘礼金（元）
	Ysyltljface string `json:"ysyl_tlj_face,omitempty" xml:"ysyl_tlj_face,omitempty"`
	// 预售有礼-淘礼金发放时间
	Ysyltljsendtime string `json:"ysyl_tlj_send_time,omitempty" xml:"ysyl_tlj_send_time,omitempty"`
	// 预售有礼-淘礼金使用开始时间
	Ysyltljusestarttime string `json:"ysyl_tlj_use_start_time,omitempty" xml:"ysyl_tlj_use_start_time,omitempty"`
	// 预售有礼-佣金比例（%）（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&amp;postId=9334376 ）
	Ysylcommissionrate string `json:"ysyl_commission_rate,omitempty" xml:"ysyl_commission_rate,omitempty"`
	// 预售有礼-淘礼金使用结束时间
	Ysyltljuseendtime string `json:"ysyl_tlj_use_end_time,omitempty" xml:"ysyl_tlj_use_end_time,omitempty"`
	// 当入参special_id、relation_id、external_id时，该字段展示预估最低佣金率(%)
	Mincommissionrate string `json:"min_commission_rate,omitempty" xml:"min_commission_rate,omitempty"`
	// 商品ID
	Itemid string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 转链成功的场景下，需要补充说明的信息
	Extrainfo string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	// 优惠券剩余量
	Couponremaincount int64 `json:"coupon_remain_count,omitempty" xml:"coupon_remain_count,omitempty"`
	// 优惠券总量
	Coupontotalcount int64 `json:"coupon_total_count,omitempty" xml:"coupon_total_count,omitempty"`
	// 优惠券(商品优惠券推广链接中的券)类型，1 公开券，2 私有券，3 妈妈券
	Coupontype int64 `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
	// 后台一级类目
	Categoryid int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 比价场景专用，当系统检测到入参消费者ID购买当前商品会获得《天天开彩蛋》玩法的彩蛋时，该字段显示1，否则为0
	Rewardinfo int64 `json:"reward_info,omitempty" xml:"reward_info,omitempty"`
	// 前N件佣金信息-当入参get_topn_rate=1，前N件佣金生效且最高，透出该组字段
	Topninfo *StepRateDto `json:"topn_info,omitempty" xml:"topn_info,omitempty"`
	// 小程序链接(暂未对外开放)
	Miniprogram *MiniProgramDto `json:"mini_program,omitempty" xml:"mini_program,omitempty"`
	// 计划类型，0代表通用计划，1代表定向计划，2代表鹊桥计划，3代表营销计划
	Campaigntype int64 `json:"campaign_type,omitempty" xml:"campaign_type,omitempty"`
}
