package tbk

// MapData 
type MapData struct {

    // 商品ID对应的sku集合
    SkuIdList   []Number `json:"sku_id_list,omitempty"`

    // 商品ID
    ItemId   int64 `json:"item_id,omitempty"`

    // 优惠券门槛金额
    CouponStartFee   string `json:"coupon_start_fee,omitempty"`

    // 优惠券剩余量
    CouponRemainCount   int64 `json:"coupon_remain_count,omitempty"`

    // 优惠券总量
    CouponTotalCount   int64 `json:"coupon_total_count,omitempty"`

    // 优惠券结束时间
    CouponEndTime   string `json:"coupon_end_time,omitempty"`

    // 优惠券开始时间
    CouponStartTime   string `json:"coupon_start_time,omitempty"`

    // 优惠券金额
    CouponAmount   string `json:"coupon_amount,omitempty"`

    // 券类型，1 表示全网公开券，4 表示妈妈渠道券
    CouponSrcScene   int64 `json:"coupon_src_scene,omitempty"`

    // 券属性，0表示店铺券，1表示单品券
    CouponType   int64 `json:"coupon_type,omitempty"`

    // 券ID
    CouponActivityId   string `json:"coupon_activity_id,omitempty"`

    // 商品信息-定向计划信息
    InfoDxjh   string `json:"info_dxjh,omitempty"`

    // 商品信息-淘客30天推广量
    TkTotalSales   string `json:"tk_total_sales,omitempty"`

    // 商品信息-月支出佣金(该字段废弃，请勿再用)
    TkTotalCommi   string `json:"tk_total_commi,omitempty"`

    // 优惠券信息-优惠券id
    CouponId   string `json:"coupon_id,omitempty"`

    // 商品信息-宝贝id(该字段废弃，请勿再用)
    NumIid   int64 `json:"num_iid,omitempty"`

    // 商品信息-商品标题
    Title   string `json:"title,omitempty"`

    // 商品信息-商品主图
    PictUrl   string `json:"pict_url,omitempty"`

    // 商品信息-商品小图列表
    SmallImages   []String `json:"small_images,omitempty"`

    // 商品信息-商品一口价格
    ReservePrice   string `json:"reserve_price,omitempty"`

    // 折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价
    ZkFinalPrice   string `json:"zk_final_price,omitempty"`

    // 店铺信息-卖家类型。0表示集市，1表示天猫
    UserType   int64 `json:"user_type,omitempty"`

    // 商品信息-宝贝所在地
    Provcity   string `json:"provcity,omitempty"`

    // 链接-宝贝地址
    ItemUrl   string `json:"item_url,omitempty"`

    // 商品信息-是否包含营销计划
    IncludeMkt   string `json:"include_mkt,omitempty"`

    // 商品信息-是否包含定向计划
    IncludeDxjh   string `json:"include_dxjh,omitempty"`

    // 商品信息-佣金比率。1550表示15.5%
    CommissionRate   string `json:"commission_rate,omitempty"`

    // 商品信息-30天销量（饿了么卡券信息-总销量）
    Volume   int64 `json:"volume,omitempty"`

    // 店铺信息-卖家id
    SellerId   int64 `json:"seller_id,omitempty"`

    // 优惠券信息-优惠券满减信息
    CouponInfo   string `json:"coupon_info,omitempty"`

    // 商品信息-佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划
    CommissionType   string `json:"commission_type,omitempty"`

    // 店铺信息-店铺名称
    ShopTitle   string `json:"shop_title,omitempty"`

    // 店铺信息-店铺dsr评分
    ShopDsr   int64 `json:"shop_dsr,omitempty"`

    // 链接-宝贝+券二合一页面链接
    CouponShareUrl   string `json:"coupon_share_url,omitempty"`

    // 链接-宝贝推广链接
    Url   string `json:"url,omitempty"`

    // 商品信息-一级类目名称
    LevelOneCategoryName   string `json:"level_one_category_name,omitempty"`

    // 商品信息-一级类目ID
    LevelOneCategoryId   int64 `json:"level_one_category_id,omitempty"`

    // 商品信息-叶子类目名称
    CategoryName   string `json:"category_name,omitempty"`

    // 商品信息-叶子类目id
    CategoryId   int64 `json:"category_id,omitempty"`

    // 商品信息-商品短标题
    ShortTitle   string `json:"short_title,omitempty"`

    // 商品信息-商品白底图
    WhiteImage   string `json:"white_image,omitempty"`

    // 拼团专用-拼团结束时间
    Oetime   string `json:"oetime,omitempty"`

    // 拼团专用-拼团开始时间
    Ostime   string `json:"ostime,omitempty"`

    // 拼团专用-拼团几人团
    JddNum   int64 `json:"jdd_num,omitempty"`

    // 拼团专用-拼团拼成价，单位元
    JddPrice   string `json:"jdd_price,omitempty"`

    // 预售专用-预售数量
    UvSumPreSale   int64 `json:"uv_sum_pre_sale,omitempty"`

    // 链接-物料块id(测试中请勿使用)
    XId   string `json:"x_id,omitempty"`

    // 商品信息-宝贝描述(推荐理由)
    ItemDescription   string `json:"item_description,omitempty"`

    // 店铺信息-卖家昵称
    Nick   string `json:"nick,omitempty"`

    // 拼团专用-拼团一人价（原价)，单位元
    OrigPrice   string `json:"orig_price,omitempty"`

    // 拼团专用-拼团库存数量
    TotalStock   int64 `json:"total_stock,omitempty"`

    // 拼团专用-拼团已售数量
    SellNum   int64 `json:"sell_num,omitempty"`

    // 拼团专用-拼团剩余库存
    Stock   int64 `json:"stock,omitempty"`

    // 营销-天猫营销玩法
    TmallPlayActivityInfo   string `json:"tmall_play_activity_info,omitempty"`

    // 商品邮费
    RealPostFee   string `json:"real_post_fee,omitempty"`

    // 锁住的佣金率
    LockRate   string `json:"lock_rate,omitempty"`

    // 锁佣结束时间
    LockRateEndTime   int64 `json:"lock_rate_end_time,omitempty"`

    // 锁佣开始时间
    LockRateStartTime   int64 `json:"lock_rate_start_time,omitempty"`

    // 预售商品-优惠
    PresaleDiscountFeeText   string `json:"presale_discount_fee_text,omitempty"`

    // 预售商品-付尾款结束时间（毫秒）
    PresaleTailEndTime   int64 `json:"presale_tail_end_time,omitempty"`

    // 预售商品-付尾款开始时间（毫秒）
    PresaleTailStartTime   int64 `json:"presale_tail_start_time,omitempty"`

    // 预售商品-付定金结束时间（毫秒）
    PresaleEndTime   int64 `json:"presale_end_time,omitempty"`

    // 预售商品-付定金开始时间（毫秒）
    PresaleStartTime   int64 `json:"presale_start_time,omitempty"`

    // 预售商品-定金（元）
    PresaleDeposit   string `json:"presale_deposit,omitempty"`

    // 预售有礼-淘礼金发放时间
    YsylTljSendTime   string `json:"ysyl_tlj_send_time,omitempty"`

    // 预售有礼-佣金比例（ 预售有礼活动享受的推广佣金比例，注：推广该活动有特殊分成规则，请详见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=9334376 ）
    YsylCommissionRate   string `json:"ysyl_commission_rate,omitempty"`

    // 预售有礼-预估淘礼金（元）
    YsylTljFace   string `json:"ysyl_tlj_face,omitempty"`

    // 预售有礼-推广链接
    YsylClickUrl   string `json:"ysyl_click_url,omitempty"`

    // 预售有礼-淘礼金使用结束时间
    YsylTljUseEndTime   string `json:"ysyl_tlj_use_end_time,omitempty"`

    // 预售有礼-淘礼金使用开始时间
    YsylTljUseStartTime   string `json:"ysyl_tlj_use_start_time,omitempty"`

    // 本地化-销售开始时间
    SaleBeginTime   string `json:"sale_begin_time,omitempty"`

    // 本地化-销售结束时间
    SaleEndTime   string `json:"sale_end_time,omitempty"`

    // 本地化-到门店距离（米）
    Distance   string `json:"distance,omitempty"`

    // 本地化-可用店铺id
    UsableShopId   string `json:"usable_shop_id,omitempty"`

    // 本地化-可用店铺名称
    UsableShopName   string `json:"usable_shop_name,omitempty"`

    // 活动价
    SalePrice   string `json:"sale_price,omitempty"`

    // 跨店满减信息
    KuadianPromotionInfo   string `json:"kuadian_promotion_info,omitempty"`

    // 是否品牌精选，0不是，1是
    SuperiorBrand   string `json:"superior_brand,omitempty"`

    // 比价场景专用，当系统检测到入参消费者ID购买当前商品会获得《天天开彩蛋》玩法的彩蛋时，该字段显示1，否则为0
    RewardInfo   int64 `json:"reward_info,omitempty"`

    // 是否品牌快抢，0不是，1是
    IsBrandFlashSale   string `json:"is_brand_flash_sale,omitempty"`

    // 本地化-扩展信息
    LocalizationExtend   *LocalizationMapData `json:"localization_extend,omitempty"`

    // 物料评估-匹配分
    MatchScore   string `json:"match_score,omitempty"`

    // 物料评估-收益分
    CommiScore   string `json:"commi_score,omitempty"`

    // 是否是热门商品，0不是，1是
    HotFlag   string `json:"hot_flag,omitempty"`

    // 前N件佣金信息-前N件佣金生效或预热时透出以下字段
    TopnInfo   *TopNInfoDTO `json:"topn_info,omitempty"`

    // 新注册时间，仅淘宝拉新适用
    RegisterTime   string `json:"register_time,omitempty"`

    // 当前活动为淘宝拉新活动时，bind_time为新激活时间； 当前活动为支付宝拉新活动时，bind_time为绑定时间。
    BindTime   string `json:"bind_time,omitempty"`

    // 首购时间，仅淘宝，天猫拉新适用
    BuyTime   string `json:"buy_time,omitempty"`

    // 新人状态， 当前活动为淘宝拉新活动时，1: 新注册，2:激活，3:首购，4:确认收货； 当前活动为支付宝实名活动时，1：已绑定，2：拉新成功，3：无效用户；当前活动为支付宝新登活动时，3：手淘首购，4：手淘确认收货；当前活动为天猫拉新活动时，2:已领取，3:已首购，4:已收货
    Status   int64 `json:"status,omitempty"`

    // 新人手机号
    Mobile   string `json:"mobile,omitempty"`

    // 订单淘客类型:1.淘客订单；2.非淘客订单，仅淘宝，天猫拉新适用
    OrderTkType   int64 `json:"order_tk_type,omitempty"`

    // 分享用户(unionid)，仅淘宝，天猫拉新适用
    UnionId   string `json:"union_id,omitempty"`

    // 来源媒体ID(pid中mm_1_2_3)中第1位
    MemberId   int64 `json:"member_id,omitempty"`

    // 来源媒体名称
    MemberNick   string `json:"member_nick,omitempty"`

    // 来源站点ID(pid中mm_1_2_3)中第2位
    SiteId   int64 `json:"site_id,omitempty"`

    // 来源站点名称
    SiteName   string `json:"site_name,omitempty"`

    // 来源广告位ID(pid中mm_1_2_3)中第3位
    AdzoneId   int64 `json:"adzone_id,omitempty"`

    // 来源广告位名称
    AdzoneName   string `json:"adzone_name,omitempty"`

    // 淘宝订单id，仅淘宝，天猫拉新适用
    TbTradeParentId   int64 `json:"tb_trade_parent_id,omitempty"`

    // 确认收货时间，仅天猫拉新适用
    AcceptTime   string `json:"accept_time,omitempty"`

    // 领取红包时间，仅天猫拉新适用
    ReceiveTime   string `json:"receive_time,omitempty"`

    // 拉新成功时间，仅支付宝拉新适用
    SuccessTime   string `json:"success_time,omitempty"`

    // 活动类型，taobao-淘宝 alipay-支付宝 tmall-天猫
    ActivityType   string `json:"activity_type,omitempty"`

    // 活动id
    ActivityId   string `json:"activity_id,omitempty"`

    // 日期，格式为"20180202"
    BizDate   string `json:"biz_date,omitempty"`

    // 复购订单，仅适用于手淘拉新
    Orders   []OrderData `json:"orders,omitempty"`

    // 绑卡日期，仅适用于手淘拉新
    BindCardTime   string `json:"bind_card_time,omitempty"`

    // loginTime
    LoginTime   string `json:"login_time,omitempty"`

    // 银行卡是否是绑定状态：1-绑定，0-未绑定
    IsCardSave   int64 `json:"is_card_save,omitempty"`

    // 使用权益时间
    UseRightsTime   string `json:"use_rights_time,omitempty"`

    // 领取权益时间
    GetRightsTime   string `json:"get_rights_time,omitempty"`

    // 渠道关系id
    RelationId   string `json:"relation_id,omitempty"`

    // 链接-宝贝+券二合一页面链接(该字段废弃，请勿再用)
    CouponClickUrl   string `json:"coupon_click_url,omitempty"`

    // 链接-宝贝推广链接
    ClickUrl   string `json:"click_url,omitempty"`

    // 商品信息-商品关联词
    WordList   []WordMapData `json:"word_list,omitempty"`

    // 商品信息-新人价
    NewUserPrice   string `json:"new_user_price,omitempty"`

    // 聚划算信息-聚淘结束时间
    JuOnlineEndTime   string `json:"ju_online_end_time,omitempty"`

    // 聚划算信息-聚淘开始时间
    JuOnlineStartTime   string `json:"ju_online_start_time,omitempty"`

    // 猫超玩法信息-活动结束时间，精确到毫秒
    MaochaoPlayEndTime   string `json:"maochao_play_end_time,omitempty"`

    // 猫超玩法信息-活动开始时间，精确到毫秒
    MaochaoPlayStartTime   string `json:"maochao_play_start_time,omitempty"`

    // 猫超玩法信息-折扣条件，价格百分数存储，件数按数量存储。可以有多个折扣条件，与折扣字段对应，','分割
    MaochaoPlayConditions   string `json:"maochao_play_conditions,omitempty"`

    // 猫超玩法信息-折扣，折扣按照百分数存储，其余按照单位分存储。可以有多个折扣，','分割
    MaochaoPlayDiscounts   string `json:"maochao_play_discounts,omitempty"`

    // 猫超玩法信息-玩法类型，2:折扣(满n件折扣),5:减钱(满n元减m元)
    MaochaoPlayDiscountType   string `json:"maochao_play_discount_type,omitempty"`

    // 猫超玩法信息-当前是否包邮，1:是，0:否
    MaochaoPlayFreePostFee   string `json:"maochao_play_free_post_fee,omitempty"`

    // 多件券优惠比例
    MultiCouponZkRate   string `json:"multi_coupon_zk_rate,omitempty"`

    // 多件券件单价
    PriceAfterMultiCoupon   string `json:"price_after_multi_coupon,omitempty"`

    // 多件券单品件数
    MultiCouponItemCount   string `json:"multi_coupon_item_count,omitempty"`

    // 满减满折的类型（1. 满减 2. 满折）
    PromotionType   string `json:"promotion_type,omitempty"`

    // 满减满折信息
    PromotionInfo   string `json:"promotion_info,omitempty"`

    // 满减满折门槛（满2件打5折中值为2；满300减20中值为300）
    PromotionDiscount   string `json:"promotion_discount,omitempty"`

    // 满减满折优惠（满2件打5折中值为5；满300减20中值为20）
    PromotionCondition   string `json:"promotion_condition,omitempty"`

    // 聚划算满减  -结束时间（毫秒）
    JuPlayEndTime   int64 `json:"ju_play_end_time,omitempty"`

    // 聚划算满减  -开始时间（毫秒）
    JuPlayStartTime   int64 `json:"ju_play_start_time,omitempty"`

    // 1聚划算满减：满N件减X元，满N件X折，满N件X元）  2天猫限时抢：前N分钟每件X元，前N分钟满N件每件X元，前N件每件X元）
    PlayInfo   string `json:"play_info,omitempty"`

    // 天猫限时抢可售  -结束时间（毫秒）
    TmallPlayActivityEndTime   int64 `json:"tmall_play_activity_end_time,omitempty"`

    // 天猫限时抢可售  -开始时间（毫秒）
    TmallPlayActivityStartTime   int64 `json:"tmall_play_activity_start_time,omitempty"`

    // 聚划算信息-商品预热开始时间（毫秒）
    JuPreShowEndTime   string `json:"ju_pre_show_end_time,omitempty"`

    // 聚划算信息-商品预热结束时间（毫秒）
    JuPreShowStartTime   string `json:"ju_pre_show_start_time,omitempty"`

    // 选品库信息
    FavoritesInfo   *FavoritesInfo `json:"favorites_info,omitempty"`

    // 商品子标题
    SubTitle   string `json:"sub_title,omitempty"`

    // 聚划算商品价格卖点描述
    JhsPriceUspList   string `json:"jhs_price_usp_list,omitempty"`

    // 淘抢购商品专用-结束时间
    TqgOnlineEndTime   string `json:"tqg_online_end_time,omitempty"`

    // 淘抢购商品专用-开团时间
    TqgOnlineStartTime   string `json:"tqg_online_start_time,omitempty"`

    // 淘抢购商品专用-已抢购数量
    TqgSoldCount   int64 `json:"tqg_sold_count,omitempty"`

    // 淘抢购商品专用-总库存
    TqgTotalCount   int64 `json:"tqg_total_count,omitempty"`

    // 优惠门槛类型： 1 满元 2 满件
    ConditionType   string `json:"condition_type,omitempty"`

    // 优惠类型： 1 减钱 2 打折
    DiscountType   string `json:"discount_type,omitempty"`

    // 权益信息-总量（权益初始库存量）
    TotalCount   int64 `json:"total_count,omitempty"`

    // 权益信息-剩余库存（权益剩余库存量）
    RemainCount   int64 `json:"remain_count,omitempty"`

    // 权益信息展示开始时间，精确到毫秒时间戳
    DisplayStartTime   string `json:"display_start_time,omitempty"`

    // 权益信息展示结束时间，精确到毫秒时间戳
    DisplayEndTime   string `json:"display_end_time,omitempty"`

    // 权益信息
    PromotionList   []PromotionList `json:"promotion_list,omitempty"`

    // 权益扩展信息
    PromotionExtend   *PromotionExtend `json:"promotion_extend,omitempty"`

    // 店铺信息-店铺logo
    ShopPictureUrl   string `json:"shop_picture_url,omitempty"`

    // 若该用户当前无待核销的红包，则返回1，若当前有待核销的红包，则返回0
    IsNewUser   string `json:"is_new_user,omitempty"`

    // 共享字段 - 备案场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）
    RelationApp   string `json:"relation_app,omitempty"`

    // 共享字段 - 备案日期
    CreateDate   string `json:"create_date,omitempty"`

    // 渠道独有 - 渠道昵称
    AccountName   string `json:"account_name,omitempty"`

    // 渠道独有 - 渠道姓名
    RealName   string `json:"real_name,omitempty"`

    // 渠道独有 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
    OfflineScene   string `json:"offline_scene,omitempty"`

    // 渠道独有 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
    OnlineScene   string `json:"online_scene,omitempty"`

    // 渠道独有 - 媒体侧渠道备注信息
    Note   string `json:"note,omitempty"`

    // 共享字段 - 渠道/会员专属pid
    RootPid   string `json:"root_pid,omitempty"`

    // 共享字段 - 渠道/会员原始身份信息
    Rtag   string `json:"rtag,omitempty"`

    // 线下备案专属信息
    OfflineInfo   *RegisterInfoDto `json:"offline_info,omitempty"`

    // 会员独有 - 会员运营ID
    SpecialId   int64 `json:"special_id,omitempty"`

    // 渠道独有 - 处罚状态
    PunishStatus   string `json:"punish_status,omitempty"`

    // 淘宝客外部用户标记，如自身系统账户ID；微信ID等
    ExternalId   string `json:"external_id,omitempty"`

    // 1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他
    ExternalType   string `json:"external_type,omitempty"`

    // 优惠券过期时间13位时间戳
    CouponExpireTime   int64 `json:"coupon_expire_time,omitempty"`

    // 非苹果ios14以上版本的设备（即其他ios版本、Android系统等），可以用此淘口令正常在复制到手淘打开
    PasswordSimple   string `json:"password_simple,omitempty"`

    // 针对苹果ios14及以上版本的苹果设备，手淘将按照示例值信息格式读取淘口令(需包含：数字+羊角符+url，识别规则可能根据ios情况变更)。如需更改淘口令内文案、url等内容，请务必先验证更改后的淘口令在手淘可被识别打开！
    Model   string `json:"model,omitempty"`

}
