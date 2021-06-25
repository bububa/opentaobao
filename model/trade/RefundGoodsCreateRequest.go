package trade

// RefundGoodsCreateRequest 
type RefundGoodsCreateRequest struct {

    // 退货商品列表
    RefundGoodsSubOrders   []RefundGoodsSubOrder `json:"refund_goods_sub_orders,omitempty"`

    // 子订单号，默认传商品列表的第一个子订单号
    SubBizOrderId   string `json:"sub_biz_order_id,omitempty"`

    // 门店标识
    ShopId   string `json:"shop_id,omitempty"`

    // 取货类型（"FETCH_HOME"：上门；"ON_SHOP"：到店；"NONE"：无需取）
    RefundFetchType   string `json:"refund_fetch_type,omitempty"`

    // 主订单号
    MainBizOrderId   string `json:"main_biz_order_id,omitempty"`

    // 买家标识
    BuyerId   string `json:"buyer_id,omitempty"`

    // 买家姓名
    BuyerName   string `json:"buyer_name,omitempty"`

    // 买家联系方式
    BuyerPhone   string `json:"buyer_phone,omitempty"`

    // 买家地址
    BuyerAddress   string `json:"buyer_address,omitempty"`

    // 操作人
    Operator   string `json:"operator,omitempty"`

    // 小二备注
    OperatorMemo   string `json:"operator_memo,omitempty"`

    // 渠道来源，欧尚外卖默认是19
    InitFrom   int64 `json:"init_from,omitempty"`

}
