package hotel

// ItemInfo 
type ItemInfo struct {

    // 实价有房
    SJYF   bool `json:"s_j_y_f,omitempty"`

    // 是否自动发货
    AutoShip   bool `json:"auto_ship,omitempty"`

    // 不含税均价
    AvgPriceWithOutTax   string `json:"avg_price_with_out_tax,omitempty"`

    // 含税均价
    AvgPriceWithTax   string `json:"avg_price_with_tax,omitempty"`

    // b2bVip
    B2bVip   bool `json:"b2b_vip,omitempty"`

    // 返现
    BackCash   int64 `json:"back_cash,omitempty"`

    // backTicketAmount
    BackTicketAmount   int64 `json:"back_ticket_amount,omitempty"`

    // 床型
    BedDesc   string `json:"bed_desc,omitempty"`

    // bookingNoticeText1
    BookingNoticeText1   string `json:"booking_notice_text1,omitempty"`

    // 早餐描述
    BreakfastDesc   string `json:"breakfast_desc,omitempty"`

    // 早餐信息
    Breakfasts   []String `json:"breakfasts,omitempty"`

    // 取消描述
    CancelDesc   string `json:"cancel_desc,omitempty"`

    // 文案
    CopyWriter   string `json:"copy_writer,omitempty"`

    // 展示价格
    CostPrice   string `json:"cost_price,omitempty"`

    // 日历外币价（酒店本地价)
    DailyPrices   []Json `json:"daily_prices,omitempty"`

    // 是否双十一
    DualEleven   bool `json:"dual_eleven,omitempty"`

    // 前端展示用：-1: 不能直接预订，0：可以预定但是库存<0,预订按钮显示为灰（卖家白名单），1: 正常预订
    Enable   bool `json:"enable,omitempty"`

    // 首住
    FirstStay   bool `json:"first_stay,omitempty"`

    // 首住最大入住x间x晚
    FirstStayLimits   []Number `json:"first_stay_limits,omitempty"`

    // 未来酒店
    Futures   []Number `json:"futures,omitempty"`

    // 礼物
    Gifts   []String `json:"gifts,omitempty"`

    // 国际卖家
    GuojiSeller   bool `json:"guoji_seller,omitempty"`

    // 发票
    HasReceipt   bool `json:"has_receipt,omitempty"`

    // 酒店id
    Hid   int64 `json:"hid,omitempty"`

    // rate级别的透传字段，json格式
    Hidden   string `json:"hidden,omitempty"`

    // 带高亮说明的title
    HighlightTitle   *HighlightContent `json:"highlight_title,omitempty"`

    // 酒店商品ID
    Id   int64 `json:"id,omitempty"`

    // IC商品数字ID
    Iid   int64 `json:"iid,omitempty"`

    // 立减
    ImmediatelySubtract   int64 `json:"immediately_subtract,omitempty"`

    // 库存描述
    InventoryDesc   string `json:"inventory_desc,omitempty"`

    // 担保
    IsGuarantee   int64 `json:"is_guarantee,omitempty"`

    // 打包
    IsHotelPackage   int64 `json:"is_hotel_package,omitempty"`

    // 会员价
    IsMember   int64 `json:"is_member,omitempty"`

    // 卖完
    IsSellOut   int64 `json:"is_sell_out,omitempty"`

    // 标签
    Labels   []Number `json:"labels,omitempty"`

    // 直减
    MinusPrice   int64 `json:"minus_price,omitempty"`

    // nop
    Nop   int64 `json:"nop,omitempty"`

    // 入住人数
    Occupancy   int64 `json:"occupancy,omitempty"`

    // 预定链接
    OrderConfirmUnits   *OrderConfirmUnits `json:"order_confirm_units,omitempty"`

    // 订单平均确认时长
    OrderShipTime   int64 `json:"order_ship_time,omitempty"`

    // 预订成功率
    OrderSucessRate   int64 `json:"order_sucess_rate,omitempty"`

    // 其他折扣
    OthersDiscounts   []String `json:"others_discounts,omitempty"`

    // 支付类型。编码取值：1:全额支付;5:前台面付;
    PaymentType   int64 `json:"payment_type,omitempty"`

    // 支付code
    PaymentTypeCode   int64 `json:"payment_type_code,omitempty"`

    // rate上的价格描述文案
    PriceDesc   string `json:"price_desc,omitempty"`

    // 大促展示的文案内容
    PromotionDescArrs   []String `json:"promotion_desc_arrs,omitempty"`

    // quickBuy逻辑：
    QuickBuy   bool `json:"quick_buy,omitempty"`

    // 库存
    Quota   int64 `json:"quota,omitempty"`

    // 商品id
    RateId   int64 `json:"rate_id,omitempty"`

    // 未落地商品库rate的唯一值，用于交易下单使用
    RateKey   string `json:"rate_key,omitempty"`

    // 退订信息
    RefundInfo   string `json:"refund_info,omitempty"`

    // 退款策略
    RefundPolicys   []String `json:"refund_policys,omitempty"`

    // 退款策略
    RefundRules   string `json:"refund_rules,omitempty"`

    // 房屋内容
    RoomContents   []HighlightContent `json:"room_contents,omitempty"`

    // 商品定价规则标题
    RpTitle   string `json:"rp_title,omitempty"`

    // 商品定价规则ID，类似SKU
    Rpid   int64 `json:"rpid,omitempty"`

    // 英文房型名称
    RtEnglishName   string `json:"rt_english_name,omitempty"`

    // 卖家房型id
    RtId   int64 `json:"rt_id,omitempty"`

    // 房型名称
    RtName   string `json:"rt_name,omitempty"`

    // 卖家ID
    SellerId   int64 `json:"seller_id,omitempty"`

    // 卖家nick
    SellerNick   string `json:"seller_nick,omitempty"`

    // 卖家综合评分
    SellerScore   string `json:"seller_score,omitempty"`

    // 卖家综合评分与同业相比结果：1:优于同业，-1:差于同业, 0:与同业相同
    SellerScoreThanAvg   int64 `json:"seller_score_than_avg,omitempty"`

    // 订单平均确认时长与同业相比结果：1:优于同业，-1:差于同业, 0:与同业相同
    ShipTimeThanAvg   int64 `json:"ship_time_than_avg,omitempty"`

    // 展示新人红包
    ShowNewPeopleCash   bool `json:"show_new_people_cash,omitempty"`

    // 展示价格
    ShowPrice   int64 `json:"show_price,omitempty"`

    // 子标题
    SubTitle   string `json:"sub_title,omitempty"`

    // 立减
    SubtractPrice   int64 `json:"subtract_price,omitempty"`

    // 预订成功率与同业相比结果:1:优于同业，-1:差于同业, 0：与同业相同
    SuccessRateThanAvg   int64 `json:"success_rate_than_avg,omitempty"`

    // 税和费
    TaxAndFee   string `json:"tax_and_fee,omitempty"`

    // 税
    TaxPrice   string `json:"tax_price,omitempty"`

    // rate上的标题：包含床型、早餐、退订政策短文案等。
    Title   string `json:"title,omitempty"`

    // 返现的总优惠总金额
    TotalBackCashAmount   int64 `json:"total_back_cash_amount,omitempty"`

    // 含税总价
    TotalPriceWithTax   string `json:"total_price_with_tax,omitempty"`

    // 立减的总优惠金额（包括立减+虚拟红包）
    TotalSubtractAmount   int64 `json:"total_subtract_amount,omitempty"`

    // virtualCardSubtractPrice
    VirtualCardSubtractPrice   int64 `json:"virtual_card_subtract_price,omitempty"`

    // discountDescs
    DiscountDescs   string `json:"discount_descs,omitempty"`

    // discountPrice
    DiscountPrice   string `json:"discount_price,omitempty"`

    // double12Desc
    Double12Desc   string `json:"double12_desc,omitempty"`

    // maxDays
    MaxDays   int64 `json:"max_days,omitempty"`

    // refundDesc
    RefundDesc   string `json:"refund_desc,omitempty"`

}
