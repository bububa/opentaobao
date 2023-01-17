package tbk

// TaobaoTbkDgOptimusMaterialMapData 结构体
type TaobaoTbkDgOptimusMaterialMapData struct {
	// 商品信息-商品小图列表
	SmallImages []string `json:"small_images,omitempty" xml:"small_images>string,omitempty"`
	// 商品信息-商品关联词
	WordList []WordMapData `json:"word_list,omitempty" xml:"word_list>word_map_data,omitempty"`
	// 定向计划集合
	SpCampaignList []SpCampaign `json:"sp_campaign_list,omitempty" xml:"sp_campaign_list>sp_campaign,omitempty"`
	// 店铺信息-店铺名称
	ShopTitle string `json:"shop_title,omitempty" xml:"shop_title,omitempty"`
	// 优惠券信息-优惠券起用门槛，满X元可用。如：满299元减20元
	CouponStartFee string `json:"coupon_start_fee,omitempty" xml:"coupon_start_fee,omitempty"`
	// 商品信息-宝贝id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价
	ZkFinalPrice string `json:"zk_final_price,omitempty" xml:"zk_final_price,omitempty"`
	// 商品信息-佣金比率(%)
	CommissionRate string `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
	// 优惠券信息-优惠券开始时间
	CouponStartTime string `json:"coupon_start_time,omitempty" xml:"coupon_start_time,omitempty"`
	// 商品信息-商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品信息-宝贝描述（推荐理由,不一定有）
	ItemDescription string `json:"item_description,omitempty" xml:"item_description,omitempty"`
	// 优惠券信息-优惠券结束时间
	CouponEndTime string `json:"coupon_end_time,omitempty" xml:"coupon_end_time,omitempty"`
	// 链接-宝贝+券二合一页面链接(该字段废弃，请勿再用)
	CouponClickUrl string `json:"coupon_click_url,omitempty" xml:"coupon_click_url,omitempty"`
	// 商品信息-商品主图
	PictUrl string `json:"pict_url,omitempty" xml:"pict_url,omitempty"`
	// 链接-宝贝推广链接
	ClickUrl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 拼团专用-拼团结束时间
	Oetime string `json:"oetime,omitempty" xml:"oetime,omitempty"`
	// 拼团专用-拼团开始时间
	Ostime string `json:"ostime,omitempty" xml:"ostime,omitempty"`
	// 拼团专用-拼团拼成价，单位元
	JddPrice string `json:"jdd_price,omitempty" xml:"jdd_price,omitempty"`
	// 拼团专用-拼团一人价（原价)，单位元
	OrigPrice string `json:"orig_price,omitempty" xml:"orig_price,omitempty"`
	// 商品信息-一级类目名称
	LevelOneCategoryName string `json:"level_one_category_name,omitempty" xml:"level_one_category_name,omitempty"`
	// 商品信息-叶子类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 商品信息-商品白底图
	WhiteImage string `json:"white_image,omitempty" xml:"white_image,omitempty"`
	// 商品信息-商品短标题
	ShortTitle string `json:"short_title,omitempty" xml:"short_title,omitempty"`
	// 营销-天猫营销玩法
	TmallPlayActivityInfo string `json:"tmall_play_activity_info,omitempty" xml:"tmall_play_activity_info,omitempty"`
	// 物料块id(测试中请勿使用)
	XId string `json:"x_id,omitempty" xml:"x_id,omitempty"`
	// 商品信息-新人价
	NewUserPrice string `json:"new_user_price,omitempty" xml:"new_user_price,omitempty"`
	// 优惠券信息-优惠券满减信息
	CouponInfo string `json:"coupon_info,omitempty" xml:"coupon_info,omitempty"`
	// 链接-宝贝+券二合一页面链接
	CouponShareUrl string `json:"coupon_share_url,omitempty" xml:"coupon_share_url,omitempty"`
	// 店铺信息-卖家昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 商品信息-一口价
	ReservePrice string `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
	// 聚划算信息-聚淘结束时间
	JuOnlineEndTime string `json:"ju_online_end_time,omitempty" xml:"ju_online_end_time,omitempty"`
	// 聚划算信息-聚淘开始时间
	JuOnlineStartTime string `json:"ju_online_start_time,omitempty" xml:"ju_online_start_time,omitempty"`
	// 猫超玩法信息-活动结束时间，精确到毫秒
	MaochaoPlayEndTime string `json:"maochao_play_end_time,omitempty" xml:"maochao_play_end_time,omitempty"`
	// 猫超玩法信息-活动开始时间，精确到毫秒
	MaochaoPlayStartTime string `json:"maochao_play_start_time,omitempty" xml:"maochao_play_start_time,omitempty"`
	// 猫超玩法信息-折扣条件，价格百分数存储，件数按数量存储。可以有多个折扣条件，与折扣字段对应，&#39;,&#39;分割
	MaochaoPlayConditions string `json:"maochao_play_conditions,omitempty" xml:"maochao_play_conditions,omitempty"`
	// 猫超玩法信息-折扣，折扣按照百分数存储，其余按照单位分存储。可以有多个折扣，&#39;,&#39;分割
	MaochaoPlayDiscounts string `json:"maochao_play_discounts,omitempty" xml:"maochao_play_discounts,omitempty"`
	// 猫超玩法信息-玩法类型，2:折扣(满n件折扣),5:减钱(满n元减m元)
	MaochaoPlayDiscountType string `json:"maochao_play_discount_type,omitempty" xml:"maochao_play_discount_type,omitempty"`
	// 猫超玩法信息-当前是否包邮，1:是，0:否
	MaochaoPlayFreePostFee string `json:"maochao_play_free_post_fee,omitempty" xml:"maochao_play_free_post_fee,omitempty"`
	// 多件券优惠比例
	MultiCouponZkRate string `json:"multi_coupon_zk_rate,omitempty" xml:"multi_coupon_zk_rate,omitempty"`
	// 多件券件单价
	PriceAfterMultiCoupon string `json:"price_after_multi_coupon,omitempty" xml:"price_after_multi_coupon,omitempty"`
	// 多件券单品件数
	MultiCouponItemCount string `json:"multi_coupon_item_count,omitempty" xml:"multi_coupon_item_count,omitempty"`
	// 锁住的佣金率
	LockRate string `json:"lock_rate,omitempty" xml:"lock_rate,omitempty"`
	// 满减满折的类型（1. 满减 2. 满折）
	PromotionType string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 满减满折信息
	PromotionInfo string `json:"promotion_info,omitempty" xml:"promotion_info,omitempty"`
	// 满减满折门槛（满2件打5折中值为2；满300减20中值为300）
	PromotionDiscount string `json:"promotion_discount,omitempty" xml:"promotion_discount,omitempty"`
	// 满减满折优惠（满2件打5折中值为5；满300减20中值为20）
	PromotionCondition string `json:"promotion_condition,omitempty" xml:"promotion_condition,omitempty"`
	// 预售商品-优惠信息
	PresaleDiscountFeeText string `json:"presale_discount_fee_text,omitempty" xml:"presale_discount_fee_text,omitempty"`
	// 预售商品-定金（元）
	PresaleDeposit string `json:"presale_deposit,omitempty" xml:"presale_deposit,omitempty"`
	// 预售有礼-淘礼金使用开始时间
	YsylTljUseStartTime string `json:"ysyl_tlj_use_start_time,omitempty" xml:"ysyl_tlj_use_start_time,omitempty"`
	// 预售有礼-佣金比例（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&amp;postId=9334376 ）
	YsylCommissionRate string `json:"ysyl_commission_rate,omitempty" xml:"ysyl_commission_rate,omitempty"`
	// 预售有礼-淘礼金发放时间
	YsylTljSendTime string `json:"ysyl_tlj_send_time,omitempty" xml:"ysyl_tlj_send_time,omitempty"`
	// 预售有礼-预估淘礼金（元）
	YsylTljFace string `json:"ysyl_tlj_face,omitempty" xml:"ysyl_tlj_face,omitempty"`
	// 预售有礼-推广链接
	YsylClickUrl string `json:"ysyl_click_url,omitempty" xml:"ysyl_click_url,omitempty"`
	// 预售有礼-淘礼金使用结束时间
	YsylTljUseEndTime string `json:"ysyl_tlj_use_end_time,omitempty" xml:"ysyl_tlj_use_end_time,omitempty"`
	// 1聚划算满减：满N件减X元，满N件X折，满N件X元）  2天猫限时抢：前N分钟每件X元，前N分钟满N件每件X元，前N件每件X元）
	PlayInfo string `json:"play_info,omitempty" xml:"play_info,omitempty"`
	// 聚划算信息-商品预热开始时间（毫秒）
	JuPreShowEndTime string `json:"ju_pre_show_end_time,omitempty" xml:"ju_pre_show_end_time,omitempty"`
	// 聚划算信息-商品预热结束时间（毫秒）
	JuPreShowStartTime string `json:"ju_pre_show_start_time,omitempty" xml:"ju_pre_show_start_time,omitempty"`
	// 活动价
	SalePrice string `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 跨店满减信息
	KuadianPromotionInfo string `json:"kuadian_promotion_info,omitempty" xml:"kuadian_promotion_info,omitempty"`
	// 商品子标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 聚划算商品价格卖点描述
	JhsPriceUspList string `json:"jhs_price_usp_list,omitempty" xml:"jhs_price_usp_list,omitempty"`
	// 淘抢购商品专用-结束时间
	TqgOnlineEndTime string `json:"tqg_online_end_time,omitempty" xml:"tqg_online_end_time,omitempty"`
	// 淘抢购商品专用-开团时间
	TqgOnlineStartTime string `json:"tqg_online_start_time,omitempty" xml:"tqg_online_start_time,omitempty"`
	// 是否品牌精选，0不是，1是
	SuperiorBrand string `json:"superior_brand,omitempty" xml:"superior_brand,omitempty"`
	// 是否品牌快抢，0不是，1是
	IsBrandFlashSale string `json:"is_brand_flash_sale,omitempty" xml:"is_brand_flash_sale,omitempty"`
	// 是否是热门商品，0不是，1是
	HotFlag string `json:"hot_flag,omitempty" xml:"hot_flag,omitempty"`
	// 商品入驻淘特后产生的所有销量量级，不特指某段具体时间
	TtSoldCount string `json:"tt_sold_count,omitempty" xml:"tt_sold_count,omitempty"`
	// 额外奖励活动类型，如果一个商品有多个奖励类型，返回结果使用空格分割，0=单单奖励(已失效)，1=超级单单奖励(已失效)，2=年货节单单奖励
	CpaRewardType string `json:"cpa_reward_type,omitempty" xml:"cpa_reward_type,omitempty"`
	// 额外奖励活动金额，活动奖励金额的类型与cpa_reward_type字段对应，如果一个商品有多个奖励类型，返回结果使用空格分割
	CpaRewardAmount string `json:"cpa_reward_amount,omitempty" xml:"cpa_reward_amount,omitempty"`
	// 合作伙伴单单补ID，用作“年货节超级单单补”活动合作伙伴奖励统计依据
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 优惠券（元） 若属于预售商品，该优惠券付尾款可用，付定金不可用
	CouponAmount int64 `json:"coupon_amount,omitempty" xml:"coupon_amount,omitempty"`
	// 商品信息-叶子类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 优惠券信息-优惠券总量
	CouponTotalCount int64 `json:"coupon_total_count,omitempty" xml:"coupon_total_count,omitempty"`
	// 店铺信息-卖家类型，0表示淘宝，1表示天猫，3表示特价版
	UserType int64 `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 优惠券信息-优惠券剩余量
	CouponRemainCount int64 `json:"coupon_remain_count,omitempty" xml:"coupon_remain_count,omitempty"`
	// 店铺信息-卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 商品信息-30天销量
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
	// 拼团专用-拼团剩余库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
	// 拼团专用-拼团已售数量
	SellNum int64 `json:"sell_num,omitempty" xml:"sell_num,omitempty"`
	// 拼团专用-拼团库存数量
	TotalStock int64 `json:"total_stock,omitempty" xml:"total_stock,omitempty"`
	// 拼团专用-拼团几人团
	JddNum int64 `json:"jdd_num,omitempty" xml:"jdd_num,omitempty"`
	// 商品信息-一级类目ID
	LevelOneCategoryId int64 `json:"level_one_category_id,omitempty" xml:"level_one_category_id,omitempty"`
	// 商品信息-预售数量
	UvSumPreSale int64 `json:"uv_sum_pre_sale,omitempty" xml:"uv_sum_pre_sale,omitempty"`
	// 锁佣结束时间
	LockRateEndTime int64 `json:"lock_rate_end_time,omitempty" xml:"lock_rate_end_time,omitempty"`
	// 锁佣开始时间
	LockRateStartTime int64 `json:"lock_rate_start_time,omitempty" xml:"lock_rate_start_time,omitempty"`
	// 预售商品-付尾款结束时间（毫秒）
	PresaleTailEndTime int64 `json:"presale_tail_end_time,omitempty" xml:"presale_tail_end_time,omitempty"`
	// 预售商品-付尾款开始时间（毫秒）
	PresaleTailStartTime int64 `json:"presale_tail_start_time,omitempty" xml:"presale_tail_start_time,omitempty"`
	// 预售商品-付定金结束时间（毫秒）
	PresaleEndTime int64 `json:"presale_end_time,omitempty" xml:"presale_end_time,omitempty"`
	// 预售商品-付定金开始时间（毫秒）
	PresaleStartTime int64 `json:"presale_start_time,omitempty" xml:"presale_start_time,omitempty"`
	// 聚划算满减  -结束时间（毫秒）
	JuPlayEndTime int64 `json:"ju_play_end_time,omitempty" xml:"ju_play_end_time,omitempty"`
	// 聚划算满减  -开始时间（毫秒）
	JuPlayStartTime int64 `json:"ju_play_start_time,omitempty" xml:"ju_play_start_time,omitempty"`
	// 天猫限时抢可售  -结束时间（毫秒）
	TmallPlayActivityEndTime int64 `json:"tmall_play_activity_end_time,omitempty" xml:"tmall_play_activity_end_time,omitempty"`
	// 天猫限时抢可售  -开始时间（毫秒）
	TmallPlayActivityStartTime int64 `json:"tmall_play_activity_start_time,omitempty" xml:"tmall_play_activity_start_time,omitempty"`
	// 选品库信息
	FavoritesInfo *FavoritesInfo `json:"favorites_info,omitempty" xml:"favorites_info,omitempty"`
	// 淘抢购商品专用-已抢购数量
	TqgSoldCount int64 `json:"tqg_sold_count,omitempty" xml:"tqg_sold_count,omitempty"`
	// 淘抢购商品专用-总库存
	TqgTotalCount int64 `json:"tqg_total_count,omitempty" xml:"tqg_total_count,omitempty"`
	// 前N件佣金信息-前N件佣金生效或预热时透出以下字段
	TopnInfo *TopNInfoDto `json:"topn_info,omitempty" xml:"topn_info,omitempty"`
	// 百亿补贴信息
	BybtInfo *BybtInfoDto `json:"bybt_info,omitempty" xml:"bybt_info,omitempty"`
	// 猫超买返卡信息
	MaifanPromotion *MaifanPromotionDto `json:"maifan_promotion,omitempty" xml:"maifan_promotion,omitempty"`
}
