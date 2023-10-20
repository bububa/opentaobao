package tbk

// TaobaotbkscoptimusmaterialMapData 结构体
type TaobaotbkscoptimusmaterialMapData struct {
	// 商品信息-商品小图列表
	Smallimages []string `json:"small_images,omitempty" xml:"small_images>string,omitempty"`
	// 商品信息-商品关联词
	Wordlist []WordMapData `json:"word_list,omitempty" xml:"word_list>word_map_data,omitempty"`
	// 定向计划集合
	Spcampaignlist []SpCampaign `json:"sp_campaign_list,omitempty" xml:"sp_campaign_list>sp_campaign,omitempty"`
	// 店铺信息-店铺名称
	Shoptitle string `json:"shop_title,omitempty" xml:"shop_title,omitempty"`
	// 优惠券信息-优惠券起用门槛，满X元可用。如：满299元减20元
	Couponstartfee string `json:"coupon_start_fee,omitempty" xml:"coupon_start_fee,omitempty"`
	// 商品信息-宝贝id
	Itemid string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品信息-在线售卖价(元)。若属于预售商品，付定金时间内，在线售卖价=预售价
	Zkfinalprice string `json:"zk_final_price,omitempty" xml:"zk_final_price,omitempty"`
	// 商品信息-佣金比率(%)
	Commissionrate string `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
	// 优惠券信息-优惠券开始时间
	Couponstarttime string `json:"coupon_start_time,omitempty" xml:"coupon_start_time,omitempty"`
	// 商品信息-商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品信息-宝贝描述(推荐理由,不一定有)
	Itemdescription string `json:"item_description,omitempty" xml:"item_description,omitempty"`
	// 优惠券信息-优惠券结束时间
	Couponendtime string `json:"coupon_end_time,omitempty" xml:"coupon_end_time,omitempty"`
	// 链接-宝贝+券二合一页面链接(该字段废弃，请勿再用)
	Couponclickurl string `json:"coupon_click_url,omitempty" xml:"coupon_click_url,omitempty"`
	// 商品信息-商品主图
	Picturl string `json:"pict_url,omitempty" xml:"pict_url,omitempty"`
	// 链接-宝贝推广链接
	Clickurl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 拼团专用-拼团结束时间
	Oetime string `json:"oetime,omitempty" xml:"oetime,omitempty"`
	// 拼团专用-拼团开始时间
	Ostime string `json:"ostime,omitempty" xml:"ostime,omitempty"`
	// 拼团专用-拼团拼成价，单位元
	Jddprice string `json:"jdd_price,omitempty" xml:"jdd_price,omitempty"`
	// 拼团专用-拼团一人价(原价)，单位元
	Origprice string `json:"orig_price,omitempty" xml:"orig_price,omitempty"`
	// 商品信息-一级类目名称
	Levelonecategoryname string `json:"level_one_category_name,omitempty" xml:"level_one_category_name,omitempty"`
	// 商品信息-叶子类目名称
	Categoryname string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 商品信息-商品短标题
	Shorttitle string `json:"short_title,omitempty" xml:"short_title,omitempty"`
	// 商品信息-商品白底图
	Whiteimage string `json:"white_image,omitempty" xml:"white_image,omitempty"`
	// 营销-天猫营销玩法
	Tmallplayactivityinfo string `json:"tmall_play_activity_info,omitempty" xml:"tmall_play_activity_info,omitempty"`
	// 商品信息-新人价
	Newuserprice string `json:"new_user_price,omitempty" xml:"new_user_price,omitempty"`
	// 优惠券信息-优惠券满减信息
	Couponinfo string `json:"coupon_info,omitempty" xml:"coupon_info,omitempty"`
	// 链接-宝贝+券二合一页面链接
	Couponshareurl string `json:"coupon_share_url,omitempty" xml:"coupon_share_url,omitempty"`
	// 店铺信息-卖家昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 商品信息-一口价通常显示为划线价
	Reserveprice string `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
	// 营销-聚划算聚淘结束时间
	Juonlineendtime string `json:"ju_online_end_time,omitempty" xml:"ju_online_end_time,omitempty"`
	// 营销-聚划算聚淘开始时间
	Juonlinestarttime string `json:"ju_online_start_time,omitempty" xml:"ju_online_start_time,omitempty"`
	// 猫超玩法信息-当前是否包邮，1:是，0:否
	Maochaoplayfreepostfee string `json:"maochao_play_free_post_fee,omitempty" xml:"maochao_play_free_post_fee,omitempty"`
	// 猫超玩法信息-活动结束时间，精确到毫秒
	Maochaoplayendtime string `json:"maochao_play_end_time,omitempty" xml:"maochao_play_end_time,omitempty"`
	// 猫超玩法信息-活动开始时间，精确到毫秒
	Maochaoplaystarttime string `json:"maochao_play_start_time,omitempty" xml:"maochao_play_start_time,omitempty"`
	// 猫超玩法信息-折扣条件，价格百分数存储，件数按数量存储。可以有多个折扣条件，与折扣字段对应，&#39;,&#39;分割
	Maochaoplayconditions string `json:"maochao_play_conditions,omitempty" xml:"maochao_play_conditions,omitempty"`
	// 猫超玩法信息-折扣，折扣按照百分数存储，其余按照单位分存储。可以有多个折扣，&#39;,&#39;分割
	Maochaoplaydiscounts string `json:"maochao_play_discounts,omitempty" xml:"maochao_play_discounts,omitempty"`
	// 猫超玩法信息-玩法类型，2:折扣(满n件折扣),5:减钱(满n元减m元)
	Maochaoplaydiscounttype string `json:"maochao_play_discount_type,omitempty" xml:"maochao_play_discount_type,omitempty"`
	// 营销-多件券优惠比例
	Multicouponzkrate string `json:"multi_coupon_zk_rate,omitempty" xml:"multi_coupon_zk_rate,omitempty"`
	// 营销-多件券件单价
	Priceaftermulticoupon string `json:"price_after_multi_coupon,omitempty" xml:"price_after_multi_coupon,omitempty"`
	// 营销-多件券件数
	Multicouponitemcount string `json:"multi_coupon_item_count,omitempty" xml:"multi_coupon_item_count,omitempty"`
	// 商品信息-锁住的佣金率
	Lockrate string `json:"lock_rate,omitempty" xml:"lock_rate,omitempty"`
	// 预售商品-优惠信息
	Presalediscountfeetext string `json:"presale_discount_fee_text,omitempty" xml:"presale_discount_fee_text,omitempty"`
	// 预售商品-定金(元)
	Presaledeposit string `json:"presale_deposit,omitempty" xml:"presale_deposit,omitempty"`
	// 预售有礼-淘礼金发放时间
	Ysyltljsendtime string `json:"ysyl_tlj_send_time,omitempty" xml:"ysyl_tlj_send_time,omitempty"`
	// 预售有礼-推广链接
	Ysylclickurl string `json:"ysyl_click_url,omitempty" xml:"ysyl_click_url,omitempty"`
	// 预售有礼-预估淘礼金(元)
	Ysyltljface string `json:"ysyl_tlj_face,omitempty" xml:"ysyl_tlj_face,omitempty"`
	// 预售有礼-佣金比例(预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&amp;postId=9334376)
	Ysylcommissionrate string `json:"ysyl_commission_rate,omitempty" xml:"ysyl_commission_rate,omitempty"`
	// 预售有礼-淘礼金使用结束时间
	Ysyltljuseendtime string `json:"ysyl_tlj_use_end_time,omitempty" xml:"ysyl_tlj_use_end_time,omitempty"`
	// 预售有礼-淘礼金使用开始时间
	Ysyltljusestarttime string `json:"ysyl_tlj_use_start_time,omitempty" xml:"ysyl_tlj_use_start_time,omitempty"`
	// 营销-1、聚划算满减满折：满N件减X元、满N件X折、满N件X元；2、天猫限时抢：前N分钟每件X元、前N分钟满N件每件X元、前N件每件X元
	Playinfo string `json:"play_info,omitempty" xml:"play_info,omitempty"`
	// 营销-满减满折的类型(1. 满减 2. 满折)
	Promotiontype string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 营销-满减满折信息
	Promotioninfo string `json:"promotion_info,omitempty" xml:"promotion_info,omitempty"`
	// 营销-满减满折门槛(满2件打5折中值为2，满300减20中值为300)
	Promotiondiscount string `json:"promotion_discount,omitempty" xml:"promotion_discount,omitempty"`
	// 营销-满减满折优惠(满2件打5折中值为5，满300减20中值为20)
	Promotioncondition string `json:"promotion_condition,omitempty" xml:"promotion_condition,omitempty"`
	// 营销-聚划算商品预热开始时间(毫秒)
	Jupreshowendtime string `json:"ju_pre_show_end_time,omitempty" xml:"ju_pre_show_end_time,omitempty"`
	// 营销-聚划算商品预热结束时间(毫秒)
	Jupreshowstarttime string `json:"ju_pre_show_start_time,omitempty" xml:"ju_pre_show_start_time,omitempty"`
	// 商品信息-大促活动预热价
	Saleprice string `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 营销-跨店满减信息
	Kuadianpromotioninfo string `json:"kuadian_promotion_info,omitempty" xml:"kuadian_promotion_info,omitempty"`
	// 商品信息-商品子标题
	Subtitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 营销-聚划算商品价格卖点描述
	Jhspriceusplist string `json:"jhs_price_usp_list,omitempty" xml:"jhs_price_usp_list,omitempty"`
	// 淘抢购商品专用-总库存
	Tqgtotalcount string `json:"tqg_total_count,omitempty" xml:"tqg_total_count,omitempty"`
	// 淘抢购商品专用-已抢购数量
	Tqgsoldcount string `json:"tqg_sold_count,omitempty" xml:"tqg_sold_count,omitempty"`
	// 淘抢购商品专用-开团时间
	Tqgonlinestarttime string `json:"tqg_online_start_time,omitempty" xml:"tqg_online_start_time,omitempty"`
	// 淘抢购商品专用-结束时间
	Tqgonlineendtime string `json:"tqg_online_end_time,omitempty" xml:"tqg_online_end_time,omitempty"`
	// 是否品牌精选，0不是，1是
	Superiorbrand string `json:"superior_brand,omitempty" xml:"superior_brand,omitempty"`
	// 是否品牌快抢，0不是，1是
	Isbrandflashsale string `json:"is_brand_flash_sale,omitempty" xml:"is_brand_flash_sale,omitempty"`
	// 是否是热门商品，0不是，1是
	Hotflag string `json:"hot_flag,omitempty" xml:"hot_flag,omitempty"`
	// 商品入驻淘特后产生的所有销量量级，不特指某段具体时间
	Ttsoldcount string `json:"tt_sold_count,omitempty" xml:"tt_sold_count,omitempty"`
	// 额外奖励活动类型，如果一个商品有多个奖励类型，返回结果使用空格分割，0=单单奖励(已失效)，1=超级单单奖励(已失效)，2=年货节单单奖励
	Cparewardtype string `json:"cpa_reward_type,omitempty" xml:"cpa_reward_type,omitempty"`
	// 额外奖励活动金额，活动奖励金额的类型与cpa_reward_type字段对应，如果一个商品有多个奖励类型，返回结果使用空格分割
	Cparewardamount string `json:"cpa_reward_amount,omitempty" xml:"cpa_reward_amount,omitempty"`
	// 合作伙伴单单补ID，用作“年货节超级单单补”活动合作伙伴奖励统计依据
	Activityid string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 优惠券信息-优惠券(元)。若属于预售商品，该优惠券付尾款可用，付定金不可用
	Couponamount int64 `json:"coupon_amount,omitempty" xml:"coupon_amount,omitempty"`
	// 商品信息-叶子类目id
	Categoryid int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 优惠券信息-优惠券总量
	Coupontotalcount int64 `json:"coupon_total_count,omitempty" xml:"coupon_total_count,omitempty"`
	// 店铺信息-卖家类型，0表示淘宝，1表示天猫，3表示特价版
	Usertype int64 `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 优惠券信息-优惠券剩余量
	Couponremaincount int64 `json:"coupon_remain_count,omitempty" xml:"coupon_remain_count,omitempty"`
	// 店铺信息-卖家id
	Sellerid int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 商品信息-30天销量
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
	// 拼团专用-拼团剩余库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
	// 拼团专用-拼团已售数量
	Sellnum int64 `json:"sell_num,omitempty" xml:"sell_num,omitempty"`
	// 拼团专用-拼团库存数量
	Totalstock int64 `json:"total_stock,omitempty" xml:"total_stock,omitempty"`
	// 拼团专用-拼团几人团
	Jddnum int64 `json:"jdd_num,omitempty" xml:"jdd_num,omitempty"`
	// 商品信息-一级类目ID
	Levelonecategoryid int64 `json:"level_one_category_id,omitempty" xml:"level_one_category_id,omitempty"`
	// 商品信息-预售数量
	Uvsumpresale int64 `json:"uv_sum_pre_sale,omitempty" xml:"uv_sum_pre_sale,omitempty"`
	// 商品信息-锁佣结束时间
	Lockrateendtime int64 `json:"lock_rate_end_time,omitempty" xml:"lock_rate_end_time,omitempty"`
	// 商品信息-锁佣开始时间
	Lockratestarttime int64 `json:"lock_rate_start_time,omitempty" xml:"lock_rate_start_time,omitempty"`
	// 预售商品-付尾款结束时间(毫秒)
	Presaletailendtime int64 `json:"presale_tail_end_time,omitempty" xml:"presale_tail_end_time,omitempty"`
	// 预售商品-付尾款开始时间(毫秒)
	Presaletailstarttime int64 `json:"presale_tail_start_time,omitempty" xml:"presale_tail_start_time,omitempty"`
	// 预售商品-付定金结束时间(毫秒)
	Presaleendtime int64 `json:"presale_end_time,omitempty" xml:"presale_end_time,omitempty"`
	// 预售商品-付定金开始时间(毫秒)
	Presalestarttime int64 `json:"presale_start_time,omitempty" xml:"presale_start_time,omitempty"`
	// 天猫限时抢可售  -结束时间(毫秒)
	Tmallplayactivityendtime int64 `json:"tmall_play_activity_end_time,omitempty" xml:"tmall_play_activity_end_time,omitempty"`
	// 天猫限时抢可售  -开始时间(毫秒)
	Tmallplayactivitystarttime int64 `json:"tmall_play_activity_start_time,omitempty" xml:"tmall_play_activity_start_time,omitempty"`
	// 营销-聚划算满减结束时间(毫秒)
	Juplayendtime int64 `json:"ju_play_end_time,omitempty" xml:"ju_play_end_time,omitempty"`
	// 营销-聚划算满减开始时间(毫秒)
	Juplaystarttime int64 `json:"ju_play_start_time,omitempty" xml:"ju_play_start_time,omitempty"`
	// 前N件佣金信息-前N件佣金生效或预热时透出以下字段
	Topninfo *TopNInfoDto `json:"topn_info,omitempty" xml:"topn_info,omitempty"`
	// 百亿补贴信息
	Bybtinfo *BybtInfoDto `json:"bybt_info,omitempty" xml:"bybt_info,omitempty"`
	// 猫超买返卡信息
	Maifanpromotion *MaifanPromotionDto `json:"maifan_promotion,omitempty" xml:"maifan_promotion,omitempty"`
}
