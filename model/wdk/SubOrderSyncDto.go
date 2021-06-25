package wdk

// SubOrderSyncDto 
type SubOrderSyncDto struct {

    // 会员优惠金额
    MemberDiscountAmt   int64 `json:"member_discount_amt,omitempty"`

    // 优惠金额
    PromotionDiscountAmt   int64 `json:"promotion_discount_amt,omitempty"`

    // 原价
    OriginalAmt   int64 `json:"original_amt,omitempty"`

    // 商品单价
    Price   int64 `json:"price,omitempty"`

    // 履约时间
    StatusChangeTime   string `json:"status_change_time,omitempty"`

    // 子订单状态
    OrderStatus   string `json:"order_status,omitempty"`

    // 淘鲜达分摊金额
    TxdPmtAmt   int64 `json:"txd_pmt_amt,omitempty"`

    // 订单类型  COMMON，GIFT
    OrderType   string `json:"order_type,omitempty"`

    // 购买库存单位数量
    BuyAmountStock   string `json:"buy_amount_stock,omitempty"`

    // 拣货金额
    PickAmt   int64 `json:"pick_amt,omitempty"`

    // 拣货库存单位数量
    PickAmountStock   string `json:"pick_amount_stock,omitempty"`

    // 外部订单号
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 淘宝订单号
    TbBizOrderId   int64 `json:"tb_biz_order_id,omitempty"`

    // 盒马订单号
    BizOrderId   int64 `json:"biz_order_id,omitempty"`

    // 售卖单位
    SaleUnit   string `json:"sale_unit,omitempty"`

    // 库存单位
    StockUnit   string `json:"stock_unit,omitempty"`

    // 扩展属性map
    TradeSubAttributes   string `json:"trade_sub_attributes,omitempty"`

    // sku码
    SkuCode   string `json:"sku_code,omitempty"`

    // 外部sku码
    OutSkuCode   string `json:"out_sku_code,omitempty"`

    // 优惠明细字段，json格式的字符串
    PromotionInfo   string `json:"promotion_info,omitempty"`

}
