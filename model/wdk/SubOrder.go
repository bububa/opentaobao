package wdk

// SubOrder 
/* model for simplify = false
type SubOrder struct {

    // 商品编码
    
    SkuCode   string `json:"sku_code,omitempty"`
    

    // 销售商品数量
    
    SaleQuantity   int64 `json:"sale_quantity,omitempty"`
    

    // 销售单价
    
    SalePrice   int64 `json:"sale_price,omitempty"`
    

    // 实际支付金额
    
    PayFee   int64 `json:"pay_fee,omitempty"`
    

    // 原始金额
    
    OriginFee   int64 `json:"origin_fee,omitempty"`
    

    // 优惠金额
    
    DiscountFee   int64 `json:"discount_fee,omitempty"`
    

    // 营销优惠明细
    
    DiscountInfos  struct {
        DiscountInfo  []DiscountInfo `json:"discount_info,omitempty"`
    } `json:"discount_infos,omitempty"`
    

    // 外部子单号
    
    SubOutOrderId   string `json:"sub_out_order_id,omitempty"`
    

    // 子单优惠金额商家分摊
    
    MerchantDiscountFee   int64 `json:"merchant_discount_fee,omitempty"`
    

    // 子单优惠金额平台分摊
    
    PlatformDiscountFee   int64 `json:"platform_discount_fee,omitempty"`
    

    // 子单商品总重量
    
    TotalWeight   int64 `json:"total_weight,omitempty"`
    

    // 加工方式说明, 非加工品不需要填写
    
    HandlingType   string `json:"handling_type,omitempty"`
    

    // 盒马子单号
    
    SubBizOrderId   string `json:"sub_biz_order_id,omitempty"`
    

    // 业务子订单id
    
    BizOrderId   int64 `json:"biz_order_id,omitempty"`
    

    // 商品id
    
    ItemCode   int64 `json:"item_code,omitempty"`
    

    // 正向：售价金额（购买数量*原售价）。逆向：退款金额
    
    OriginalAmt   int64 `json:"original_amt,omitempty"`
    

    // 父订单id
    
    ParentId   int64 `json:"parent_id,omitempty"`
    

    // 商品原价
    
    Price   int64 `json:"price,omitempty"`
    

    // 促销信息(json格式)。例如:[{"activity_id":"1234","activity_name":"五一促销","activity_type":1,"activity_desc":"优惠卡券"}]
    
    PromotionInfo   string `json:"promotion_info,omitempty"`
    

    // 购买数量
    
    Quantity   int64 `json:"quantity,omitempty"`
    

    // 门店编码
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 销售类型（正向销售1：逆向销售2。本接口应返回2）
    
    TrdType   int64 `json:"trd_type,omitempty"`
    

    // 促销优惠总金额
    
    PromotionDiscountAmt   int64 `json:"promotion_discount_amt,omitempty"`
    

    // 会员价优惠金额
    
    MemberDiscountAmt   int64 `json:"member_discount_amt,omitempty"`
    

    // 其它分摊优惠金额
    
    ShareDiscountAmt   int64 `json:"share_discount_amt,omitempty"`
    

    // 非标品销售单位
    
    SellUnit   string `json:"sell_unit,omitempty"`
    

    // 非标品购买数量
    
    NsQuantity   string `json:"ns_quantity,omitempty"`
    

    // 淘宝子订单号
    
    TbBizOrderId   int64 `json:"tb_biz_order_id,omitempty"`
    

    // 库存单位拣货数量
    
    PickAmountStock   string `json:"pick_amount_stock,omitempty"`
    

    // 库存单位购买数量
    
    BuyAmountStock   string `json:"buy_amount_stock,omitempty"`
    

    // memberPoint
    
    MemberPoint   string `json:"member_point,omitempty"`
    

    // 子订单类型，当前取值[COMMON|GIFT],分别表示 普通|赠品 订单
    
    OrderType   string `json:"order_type,omitempty"`
    

    // 淘鲜达平台优惠券中平台出资金额，单位分
    
    TxdPmtAmt   int64 `json:"txd_pmt_amt,omitempty"`
    

    // 拣货金额
    
    PickAmt   int64 `json:"pick_amt,omitempty"`
    

    // 揽件
    
    OrderStatus   string `json:"order_status,omitempty"`
    

    // statusChangeTime
    
    StatusChangeTime   string `json:"status_change_time,omitempty"`
    

    // 库存单位
    
    StockUnit   string `json:"stock_unit,omitempty"`
    

    // 销售单位
    
    SaleUnit   string `json:"sale_unit,omitempty"`
    

    // tradeSubAttributes
    
    TradeSubAttributes   string `json:"trade_sub_attributes,omitempty"`
    

    // 外部skucode
    
    OutSkuCode   string `json:"out_sku_code,omitempty"`
    

    // 外部订单号
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

}
*/

// SubOrder 
type SubOrder struct {

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 销售商品数量
    SaleQuantity   int64 `json:"sale_quantity,omitempty"`

    // 销售单价
    SalePrice   int64 `json:"sale_price,omitempty"`

    // 实际支付金额
    PayFee   int64 `json:"pay_fee,omitempty"`

    // 原始金额
    OriginFee   int64 `json:"origin_fee,omitempty"`

    // 优惠金额
    DiscountFee   int64 `json:"discount_fee,omitempty"`

    // 营销优惠明细
    DiscountInfos   []DiscountInfo `json:"discount_infos,omitempty"`

    // 外部子单号
    SubOutOrderId   string `json:"sub_out_order_id,omitempty"`

    // 子单优惠金额商家分摊
    MerchantDiscountFee   int64 `json:"merchant_discount_fee,omitempty"`

    // 子单优惠金额平台分摊
    PlatformDiscountFee   int64 `json:"platform_discount_fee,omitempty"`

    // 子单商品总重量
    TotalWeight   int64 `json:"total_weight,omitempty"`

    // 加工方式说明, 非加工品不需要填写
    HandlingType   string `json:"handling_type,omitempty"`

    // 盒马子单号
    SubBizOrderId   string `json:"sub_biz_order_id,omitempty"`

    // 业务子订单id
    BizOrderId   int64 `json:"biz_order_id,omitempty"`

    // 商品id
    ItemCode   int64 `json:"item_code,omitempty"`

    // 正向：售价金额（购买数量*原售价）。逆向：退款金额
    OriginalAmt   int64 `json:"original_amt,omitempty"`

    // 父订单id
    ParentId   int64 `json:"parent_id,omitempty"`

    // 商品原价
    Price   int64 `json:"price,omitempty"`

    // 促销信息(json格式)。例如:[{"activity_id":"1234","activity_name":"五一促销","activity_type":1,"activity_desc":"优惠卡券"}]
    PromotionInfo   string `json:"promotion_info,omitempty"`

    // 购买数量
    Quantity   int64 `json:"quantity,omitempty"`

    // 门店编码
    StoreId   string `json:"store_id,omitempty"`

    // 销售类型（正向销售1：逆向销售2。本接口应返回2）
    TrdType   int64 `json:"trd_type,omitempty"`

    // 促销优惠总金额
    PromotionDiscountAmt   int64 `json:"promotion_discount_amt,omitempty"`

    // 会员价优惠金额
    MemberDiscountAmt   int64 `json:"member_discount_amt,omitempty"`

    // 其它分摊优惠金额
    ShareDiscountAmt   int64 `json:"share_discount_amt,omitempty"`

    // 非标品销售单位
    SellUnit   string `json:"sell_unit,omitempty"`

    // 非标品购买数量
    NsQuantity   string `json:"ns_quantity,omitempty"`

    // 淘宝子订单号
    TbBizOrderId   int64 `json:"tb_biz_order_id,omitempty"`

    // 库存单位拣货数量
    PickAmountStock   string `json:"pick_amount_stock,omitempty"`

    // 库存单位购买数量
    BuyAmountStock   string `json:"buy_amount_stock,omitempty"`

    // memberPoint
    MemberPoint   string `json:"member_point,omitempty"`

    // 子订单类型，当前取值[COMMON|GIFT],分别表示 普通|赠品 订单
    OrderType   string `json:"order_type,omitempty"`

    // 淘鲜达平台优惠券中平台出资金额，单位分
    TxdPmtAmt   int64 `json:"txd_pmt_amt,omitempty"`

    // 拣货金额
    PickAmt   int64 `json:"pick_amt,omitempty"`

    // 揽件
    OrderStatus   string `json:"order_status,omitempty"`

    // statusChangeTime
    StatusChangeTime   string `json:"status_change_time,omitempty"`

    // 库存单位
    StockUnit   string `json:"stock_unit,omitempty"`

    // 销售单位
    SaleUnit   string `json:"sale_unit,omitempty"`

    // tradeSubAttributes
    TradeSubAttributes   string `json:"trade_sub_attributes,omitempty"`

    // 外部skucode
    OutSkuCode   string `json:"out_sku_code,omitempty"`

    // 外部订单号
    OutOrderId   string `json:"out_order_id,omitempty"`

}
