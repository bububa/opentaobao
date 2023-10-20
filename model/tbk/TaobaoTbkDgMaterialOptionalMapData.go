package tbk

// TaobaotbkdgmaterialoptionalMapData 结构体
type TaobaotbkdgmaterialoptionalMapData struct {
	// 商品信息-商品小图列表
	Smallimages []string `json:"small_images,omitempty" xml:"small_images>string,omitempty"`
	// 定向计划集合
	Spcampaignlist []SpCampaign `json:"sp_campaign_list,omitempty" xml:"sp_campaign_list>sp_campaign,omitempty"`
	// 优惠券信息-优惠券开始时间
	Couponstarttime string `json:"coupon_start_time,omitempty" xml:"coupon_start_time,omitempty"`
	// 优惠券信息-优惠券结束时间
	Couponendtime string `json:"coupon_end_time,omitempty" xml:"coupon_end_time,omitempty"`
	// 商品信息-定向计划信息
	Infodxjh string `json:"info_dxjh,omitempty" xml:"info_dxjh,omitempty"`
	// 商品信息-淘客30天推广量
	Tktotalsales string `json:"tk_total_sales,omitempty" xml:"tk_total_sales,omitempty"`
	// 商品信息-月支出佣金(该字段废弃，请勿再用)
	Tktotalcommi string `json:"tk_total_commi,omitempty" xml:"tk_total_commi,omitempty"`
	// 优惠券信息-优惠券id
	Couponid string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
	// 商品信息-宝贝id(该字段废弃，请勿再用)
	Numiid string `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 商品信息-商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品信息-商品主图
	Picturl string `json:"pict_url,omitempty" xml:"pict_url,omitempty"`
	// 商品信息-商品一口价格
	Reserveprice string `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
	// 折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价
	Zkfinalprice string `json:"zk_final_price,omitempty" xml:"zk_final_price,omitempty"`
	// 商品信息-宝贝所在地
	Provcity string `json:"provcity,omitempty" xml:"provcity,omitempty"`
	// 链接-宝贝地址
	Itemurl string `json:"item_url,omitempty" xml:"item_url,omitempty"`
	// 商品信息-是否包含营销计划
	Includemkt string `json:"include_mkt,omitempty" xml:"include_mkt,omitempty"`
	// 商品信息-是否包含定向计划
	Includedxjh string `json:"include_dxjh,omitempty" xml:"include_dxjh,omitempty"`
	// 商品信息-佣金比率。1550表示15.5%
	Commissionrate string `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
	// 店铺信息-店铺名称
	Shoptitle string `json:"shop_title,omitempty" xml:"shop_title,omitempty"`
	// 优惠券信息-优惠券满减信息
	Couponinfo string `json:"coupon_info,omitempty" xml:"coupon_info,omitempty"`
	// 商品信息-佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划
	Commissiontype string `json:"commission_type,omitempty" xml:"commission_type,omitempty"`
	// 链接-宝贝+券二合一页面链接
	Couponshareurl string `json:"coupon_share_url,omitempty" xml:"coupon_share_url,omitempty"`
	// 链接-宝贝推广链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 商品信息-一级类目名称
	Levelonecategoryname string `json:"level_one_category_name,omitempty" xml:"level_one_category_name,omitempty"`
	// 商品信息-叶子类目名称
	Categoryname string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 商品信息-商品短标题
	Shorttitle string `json:"short_title,omitempty" xml:"short_title,omitempty"`
	// 商品信息-商品白底图
	Whiteimage string `json:"white_image,omitempty" xml:"white_image,omitempty"`
	// 拼团专用-拼团结束时间
	Oetime string `json:"oetime,omitempty" xml:"oetime,omitempty"`
	// 拼团专用-拼团开始时间
	Ostime string `json:"ostime,omitempty" xml:"ostime,omitempty"`
	// 拼团专用-拼团拼成价，单位元
	Jddprice string `json:"jdd_price,omitempty" xml:"jdd_price,omitempty"`
	// 链接-物料块id(测试中请勿使用)
	Xid string `json:"x_id,omitempty" xml:"x_id,omitempty"`
	// 优惠券信息-优惠券起用门槛，满X元可用。如：满299元减20元
	Couponstartfee string `json:"coupon_start_fee,omitempty" xml:"coupon_start_fee,omitempty"`
	// 优惠券（元） 若属于预售商品，该优惠券付尾款可用，付定金不可用
	Couponamount string `json:"coupon_amount,omitempty" xml:"coupon_amount,omitempty"`
	// 商品信息-宝贝描述(推荐理由)
	Itemdescription string `json:"item_description,omitempty" xml:"item_description,omitempty"`
	// 店铺信息-卖家昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 拼团专用-拼团一人价（原价)，单位元
	Origprice string `json:"orig_price,omitempty" xml:"orig_price,omitempty"`
	// 营销-天猫营销玩法
	Tmallplayactivityinfo string `json:"tmall_play_activity_info,omitempty" xml:"tmall_play_activity_info,omitempty"`
	// 商品信息-宝贝id
	Itemid string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品邮费
	Realpostfee string `json:"real_post_fee,omitempty" xml:"real_post_fee,omitempty"`
	// 锁住的佣金率
	Lockrate string `json:"lock_rate,omitempty" xml:"lock_rate,omitempty"`
	// 预售商品-优惠
	Presalediscountfeetext string `json:"presale_discount_fee_text,omitempty" xml:"presale_discount_fee_text,omitempty"`
	// 预售商品-定金（元）
	Presaledeposit string `json:"presale_deposit,omitempty" xml:"presale_deposit,omitempty"`
	// 预售有礼-淘礼金发放时间
	Ysyltljsendtime string `json:"ysyl_tlj_send_time,omitempty" xml:"ysyl_tlj_send_time,omitempty"`
	// 预售有礼-佣金比例（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&amp;postId=9334376 ）
	Ysylcommissionrate string `json:"ysyl_commission_rate,omitempty" xml:"ysyl_commission_rate,omitempty"`
	// 预售有礼-预估淘礼金（元）
	Ysyltljface string `json:"ysyl_tlj_face,omitempty" xml:"ysyl_tlj_face,omitempty"`
	// 预售有礼-推广链接
	Ysylclickurl string `json:"ysyl_click_url,omitempty" xml:"ysyl_click_url,omitempty"`
	// 预售有礼-淘礼金使用结束时间
	Ysyltljuseendtime string `json:"ysyl_tlj_use_end_time,omitempty" xml:"ysyl_tlj_use_end_time,omitempty"`
	// 预售有礼-淘礼金使用开始时间
	Ysyltljusestarttime string `json:"ysyl_tlj_use_start_time,omitempty" xml:"ysyl_tlj_use_start_time,omitempty"`
	// 本地化-销售开始时间
	Salebegintime string `json:"sale_begin_time,omitempty" xml:"sale_begin_time,omitempty"`
	// 本地化-销售结束时间
	Saleendtime string `json:"sale_end_time,omitempty" xml:"sale_end_time,omitempty"`
	// 本地化-到门店距离（米）
	Distance string `json:"distance,omitempty" xml:"distance,omitempty"`
	// 本地化-可用店铺id
	Usableshopid string `json:"usable_shop_id,omitempty" xml:"usable_shop_id,omitempty"`
	// 本地化-可用店铺名称
	Usableshopname string `json:"usable_shop_name,omitempty" xml:"usable_shop_name,omitempty"`
	// 活动价
	Saleprice string `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 跨店满减信息
	Kuadianpromotioninfo string `json:"kuadian_promotion_info,omitempty" xml:"kuadian_promotion_info,omitempty"`
	// 是否品牌精选，0不是，1是
	Superiorbrand string `json:"superior_brand,omitempty" xml:"superior_brand,omitempty"`
	// 是否品牌快抢，0不是，1是
	Isbrandflashsale string `json:"is_brand_flash_sale,omitempty" xml:"is_brand_flash_sale,omitempty"`
	// 本地化-扩展信息
	Localizationextend string `json:"localization_extend,omitempty" xml:"localization_extend,omitempty"`
	// 物料评估-匹配分
	Matchscore string `json:"match_score,omitempty" xml:"match_score,omitempty"`
	// 物料评估-收益分
	Commiscore string `json:"commi_score,omitempty" xml:"commi_score,omitempty"`
	// 是否是热门商品，0不是，1是
	Hotflag string `json:"hot_flag,omitempty" xml:"hot_flag,omitempty"`
	// 商品入驻淘特后产生的所有销量量级，不特指某段具体时间
	Ttsoldcount string `json:"tt_sold_count,omitempty" xml:"tt_sold_count,omitempty"`
	// 额外奖励活动类型，如果一个商品有多个奖励类型，返回结果使用空格分割，0=预售单单奖励，1=618超级U选单单补
	Cparewardtype string `json:"cpa_reward_type,omitempty" xml:"cpa_reward_type,omitempty"`
	// 额外奖励活动金额，活动奖励金额的类型与cpa_reward_type字段对应，如果一个商品有多个奖励类型，返回结果使用空格分割
	Cparewardamount string `json:"cpa_reward_amount,omitempty" xml:"cpa_reward_amount,omitempty"`
	// 合作伙伴单单补ID，用作“年货节超级单单补”活动合作伙伴奖励统计依据
	Activityid string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 榜单url
	Rankpageurl string `json:"rank_page_url,omitempty" xml:"rank_page_url,omitempty"`
	// 搜索类型
	Itemsearchtype string `json:"item_search_type,omitempty" xml:"item_search_type,omitempty"`
	// 店铺信息-卖家类型。0表示集市，1表示天猫
	Usertype int64 `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 商品信息-30天销量（饿了么卡券信息-总销量）
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
	// 店铺信息-卖家id
	Sellerid int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 优惠券信息-优惠券总量
	Coupontotalcount int64 `json:"coupon_total_count,omitempty" xml:"coupon_total_count,omitempty"`
	// 优惠券信息-优惠券剩余量
	Couponremaincount int64 `json:"coupon_remain_count,omitempty" xml:"coupon_remain_count,omitempty"`
	// 店铺信息-店铺dsr评分
	Shopdsr int64 `json:"shop_dsr,omitempty" xml:"shop_dsr,omitempty"`
	// 商品信息-一级类目ID
	Levelonecategoryid int64 `json:"level_one_category_id,omitempty" xml:"level_one_category_id,omitempty"`
	// 商品信息-叶子类目id
	Categoryid int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 拼团专用-拼团几人团
	Jddnum int64 `json:"jdd_num,omitempty" xml:"jdd_num,omitempty"`
	// 预售专用-预售数量
	Uvsumpresale int64 `json:"uv_sum_pre_sale,omitempty" xml:"uv_sum_pre_sale,omitempty"`
	// 拼团专用-拼团库存数量
	Totalstock int64 `json:"total_stock,omitempty" xml:"total_stock,omitempty"`
	// 拼团专用-拼团已售数量
	Sellnum int64 `json:"sell_num,omitempty" xml:"sell_num,omitempty"`
	// 拼团专用-拼团剩余库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
	// 锁佣结束时间
	Lockrateendtime int64 `json:"lock_rate_end_time,omitempty" xml:"lock_rate_end_time,omitempty"`
	// 锁佣开始时间
	Lockratestarttime int64 `json:"lock_rate_start_time,omitempty" xml:"lock_rate_start_time,omitempty"`
	// 预售商品-付尾款结束时间（毫秒）
	Presaletailendtime int64 `json:"presale_tail_end_time,omitempty" xml:"presale_tail_end_time,omitempty"`
	// 预售商品-付尾款开始时间（毫秒）
	Presaletailstarttime int64 `json:"presale_tail_start_time,omitempty" xml:"presale_tail_start_time,omitempty"`
	// 预售商品-付定金结束时间（毫秒）
	Presaleendtime int64 `json:"presale_end_time,omitempty" xml:"presale_end_time,omitempty"`
	// 预售商品-付定金开始时间（毫秒）
	Presalestarttime int64 `json:"presale_start_time,omitempty" xml:"presale_start_time,omitempty"`
	// 比价场景专用，当系统检测到入参消费者ID购买当前商品会获得《天天开彩蛋》玩法的彩蛋时，该字段显示1，否则为0
	Rewardinfo int64 `json:"reward_info,omitempty" xml:"reward_info,omitempty"`
	// 前N件佣金信息-前N件佣金生效或预热时透出以下字段
	Topninfo *TopNInfoDto `json:"topn_info,omitempty" xml:"topn_info,omitempty"`
	// 百亿补贴信息
	Bybtinfo *BybtInfoDto `json:"bybt_info,omitempty" xml:"bybt_info,omitempty"`
	// 猫超买返卡信息
	Maifanpromotion *MaifanPromotionDto `json:"maifan_promotion,omitempty" xml:"maifan_promotion,omitempty"`
}
