package tbk

// TaobaoTbkDgMaterialOptionalMapData 结构体
type TaobaoTbkDgMaterialOptionalMapData struct {
	// 商品信息-商品小图列表
	SmallImages []string `json:"small_images,omitempty" xml:"small_images>string,omitempty"`
	// 优惠券信息-优惠券开始时间
	CouponStartTime string `json:"coupon_start_time,omitempty" xml:"coupon_start_time,omitempty"`
	// 优惠券信息-优惠券结束时间
	CouponEndTime string `json:"coupon_end_time,omitempty" xml:"coupon_end_time,omitempty"`
	// 商品信息-定向计划信息
	InfoDxjh string `json:"info_dxjh,omitempty" xml:"info_dxjh,omitempty"`
	// 商品信息-淘客30天推广量
	TkTotalSales string `json:"tk_total_sales,omitempty" xml:"tk_total_sales,omitempty"`
	// 商品信息-月支出佣金(该字段废弃，请勿再用)
	TkTotalCommi string `json:"tk_total_commi,omitempty" xml:"tk_total_commi,omitempty"`
	// 优惠券信息-优惠券id
	CouponId string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
	// 商品信息-宝贝id(该字段废弃，请勿再用)
	NumIid string `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 商品信息-商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品信息-商品主图
	PictUrl string `json:"pict_url,omitempty" xml:"pict_url,omitempty"`
	// 商品信息-商品一口价格
	ReservePrice string `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
	// 折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价
	ZkFinalPrice string `json:"zk_final_price,omitempty" xml:"zk_final_price,omitempty"`
	// 商品信息-宝贝所在地
	Provcity string `json:"provcity,omitempty" xml:"provcity,omitempty"`
	// 链接-宝贝地址
	ItemUrl string `json:"item_url,omitempty" xml:"item_url,omitempty"`
	// 商品信息-是否包含营销计划
	IncludeMkt string `json:"include_mkt,omitempty" xml:"include_mkt,omitempty"`
	// 商品信息-是否包含定向计划
	IncludeDxjh string `json:"include_dxjh,omitempty" xml:"include_dxjh,omitempty"`
	// 商品信息-佣金比率。1550表示15.5%
	CommissionRate string `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
	// 店铺信息-店铺名称
	ShopTitle string `json:"shop_title,omitempty" xml:"shop_title,omitempty"`
	// 优惠券信息-优惠券满减信息
	CouponInfo string `json:"coupon_info,omitempty" xml:"coupon_info,omitempty"`
	// 商品信息-佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划
	CommissionType string `json:"commission_type,omitempty" xml:"commission_type,omitempty"`
	// 链接-宝贝+券二合一页面链接
	CouponShareUrl string `json:"coupon_share_url,omitempty" xml:"coupon_share_url,omitempty"`
	// 链接-宝贝推广链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 商品信息-一级类目名称
	LevelOneCategoryName string `json:"level_one_category_name,omitempty" xml:"level_one_category_name,omitempty"`
	// 商品信息-叶子类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 商品信息-商品短标题
	ShortTitle string `json:"short_title,omitempty" xml:"short_title,omitempty"`
	// 商品信息-商品白底图
	WhiteImage string `json:"white_image,omitempty" xml:"white_image,omitempty"`
	// 拼团专用-拼团结束时间
	Oetime string `json:"oetime,omitempty" xml:"oetime,omitempty"`
	// 拼团专用-拼团开始时间
	Ostime string `json:"ostime,omitempty" xml:"ostime,omitempty"`
	// 拼团专用-拼团拼成价，单位元
	JddPrice string `json:"jdd_price,omitempty" xml:"jdd_price,omitempty"`
	// 链接-物料块id(测试中请勿使用)
	XId string `json:"x_id,omitempty" xml:"x_id,omitempty"`
	// 优惠券信息-优惠券起用门槛，满X元可用。如：满299元减20元
	CouponStartFee string `json:"coupon_start_fee,omitempty" xml:"coupon_start_fee,omitempty"`
	// 优惠券（元） 若属于预售商品，该优惠券付尾款可用，付定金不可用
	CouponAmount string `json:"coupon_amount,omitempty" xml:"coupon_amount,omitempty"`
	// 商品信息-宝贝描述(推荐理由)
	ItemDescription string `json:"item_description,omitempty" xml:"item_description,omitempty"`
	// 店铺信息-卖家昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 拼团专用-拼团一人价（原价)，单位元
	OrigPrice string `json:"orig_price,omitempty" xml:"orig_price,omitempty"`
	// 营销-天猫营销玩法
	TmallPlayActivityInfo string `json:"tmall_play_activity_info,omitempty" xml:"tmall_play_activity_info,omitempty"`
	// 商品信息-宝贝id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品邮费
	RealPostFee string `json:"real_post_fee,omitempty" xml:"real_post_fee,omitempty"`
	// 锁住的佣金率
	LockRate string `json:"lock_rate,omitempty" xml:"lock_rate,omitempty"`
	// 预售商品-优惠
	PresaleDiscountFeeText string `json:"presale_discount_fee_text,omitempty" xml:"presale_discount_fee_text,omitempty"`
	// 预售商品-定金（元）
	PresaleDeposit string `json:"presale_deposit,omitempty" xml:"presale_deposit,omitempty"`
	// 预售有礼-淘礼金发放时间
	YsylTljSendTime string `json:"ysyl_tlj_send_time,omitempty" xml:"ysyl_tlj_send_time,omitempty"`
	// 预售有礼-佣金比例（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=9334376 ）
	YsylCommissionRate string `json:"ysyl_commission_rate,omitempty" xml:"ysyl_commission_rate,omitempty"`
	// 预售有礼-预估淘礼金（元）
	YsylTljFace string `json:"ysyl_tlj_face,omitempty" xml:"ysyl_tlj_face,omitempty"`
	// 预售有礼-推广链接
	YsylClickUrl string `json:"ysyl_click_url,omitempty" xml:"ysyl_click_url,omitempty"`
	// 预售有礼-淘礼金使用结束时间
	YsylTljUseEndTime string `json:"ysyl_tlj_use_end_time,omitempty" xml:"ysyl_tlj_use_end_time,omitempty"`
	// 预售有礼-淘礼金使用开始时间
	YsylTljUseStartTime string `json:"ysyl_tlj_use_start_time,omitempty" xml:"ysyl_tlj_use_start_time,omitempty"`
	// 本地化-销售开始时间
	SaleBeginTime string `json:"sale_begin_time,omitempty" xml:"sale_begin_time,omitempty"`
	// 本地化-销售结束时间
	SaleEndTime string `json:"sale_end_time,omitempty" xml:"sale_end_time,omitempty"`
	// 本地化-到门店距离（米）
	Distance string `json:"distance,omitempty" xml:"distance,omitempty"`
	// 本地化-可用店铺id
	UsableShopId string `json:"usable_shop_id,omitempty" xml:"usable_shop_id,omitempty"`
	// 本地化-可用店铺名称
	UsableShopName string `json:"usable_shop_name,omitempty" xml:"usable_shop_name,omitempty"`
	// 活动价
	SalePrice string `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 跨店满减信息
	KuadianPromotionInfo string `json:"kuadian_promotion_info,omitempty" xml:"kuadian_promotion_info,omitempty"`
	// 是否品牌精选，0不是，1是
	SuperiorBrand string `json:"superior_brand,omitempty" xml:"superior_brand,omitempty"`
	// 是否品牌快抢，0不是，1是
	IsBrandFlashSale string `json:"is_brand_flash_sale,omitempty" xml:"is_brand_flash_sale,omitempty"`
	// 本地化-扩展信息
	LocalizationExtend string `json:"localization_extend,omitempty" xml:"localization_extend,omitempty"`
	// 物料评估-匹配分
	MatchScore string `json:"match_score,omitempty" xml:"match_score,omitempty"`
	// 物料评估-收益分
	CommiScore string `json:"commi_score,omitempty" xml:"commi_score,omitempty"`
	// 是否是热门商品，0不是，1是
	HotFlag string `json:"hot_flag,omitempty" xml:"hot_flag,omitempty"`
	// 商品入驻淘特后产生的所有销量量级，不特指某段具体时间
	TtSoldCount string `json:"tt_sold_count,omitempty" xml:"tt_sold_count,omitempty"`
	// 额外奖励活动类型，如果一个商品有多个奖励类型，返回结果使用空格分割，0=预售单单奖励，1=618超级U选单单补
	CpaRewardType string `json:"cpa_reward_type,omitempty" xml:"cpa_reward_type,omitempty"`
	// 额外奖励活动金额，活动奖励金额的类型与cpa_reward_type字段对应，如果一个商品有多个奖励类型，返回结果使用空格分割
	CpaRewardAmount string `json:"cpa_reward_amount,omitempty" xml:"cpa_reward_amount,omitempty"`
	// 合作伙伴单单补ID，用作“年货节超级单单补”活动合作伙伴奖励统计依据
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 榜单url
	RankPageUrl string `json:"rank_page_url,omitempty" xml:"rank_page_url,omitempty"`
	// 搜索类型
	ItemSearchType string `json:"item_search_type,omitempty" xml:"item_search_type,omitempty"`
	// 店铺信息-卖家类型。0表示集市，1表示天猫
	UserType int64 `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 商品信息-30天销量（饿了么卡券信息-总销量）
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
	// 店铺信息-卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 优惠券信息-优惠券总量
	CouponTotalCount int64 `json:"coupon_total_count,omitempty" xml:"coupon_total_count,omitempty"`
	// 优惠券信息-优惠券剩余量
	CouponRemainCount int64 `json:"coupon_remain_count,omitempty" xml:"coupon_remain_count,omitempty"`
	// 店铺信息-店铺dsr评分
	ShopDsr int64 `json:"shop_dsr,omitempty" xml:"shop_dsr,omitempty"`
	// 商品信息-一级类目ID
	LevelOneCategoryId int64 `json:"level_one_category_id,omitempty" xml:"level_one_category_id,omitempty"`
	// 商品信息-叶子类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 拼团专用-拼团几人团
	JddNum int64 `json:"jdd_num,omitempty" xml:"jdd_num,omitempty"`
	// 预售专用-预售数量
	UvSumPreSale int64 `json:"uv_sum_pre_sale,omitempty" xml:"uv_sum_pre_sale,omitempty"`
	// 拼团专用-拼团库存数量
	TotalStock int64 `json:"total_stock,omitempty" xml:"total_stock,omitempty"`
	// 拼团专用-拼团已售数量
	SellNum int64 `json:"sell_num,omitempty" xml:"sell_num,omitempty"`
	// 拼团专用-拼团剩余库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
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
	// 比价场景专用，当系统检测到入参消费者ID购买当前商品会获得《天天开彩蛋》玩法的彩蛋时，该字段显示1，否则为0
	RewardInfo int64 `json:"reward_info,omitempty" xml:"reward_info,omitempty"`
	// 前N件佣金信息-前N件佣金生效或预热时透出以下字段
	TopnInfo *TopNInfoDto `json:"topn_info,omitempty" xml:"topn_info,omitempty"`
	// 百亿补贴信息
	BybtInfo *BybtInfoDto `json:"bybt_info,omitempty" xml:"bybt_info,omitempty"`
	// 猫超买返卡信息
	MaifanPromotion *MaifanPromotionDto `json:"maifan_promotion,omitempty" xml:"maifan_promotion,omitempty"`
}
