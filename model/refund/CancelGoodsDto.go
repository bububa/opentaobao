package refund

// CancelGoodsDto 
type CancelGoodsDto struct {

    // 子订单ID
    Oid   int64 `json:"oid,omitempty"`

    // 退款单ID
    RefundId   int64 `json:"refund_id,omitempty"`

    // 操作时间
    OperateTime   string `json:"operate_time,omitempty"`

    // 退款金额 单位 分
    RefundFee   int64 `json:"refund_fee,omitempty"`

    // 状态SUCCESS、FAIL
    Status   string `json:"status,omitempty"`

    // 商家商品ID
    OuterId   string `json:"outer_id,omitempty"`

    // 商品ID
    AuctionId   int64 `json:"auction_id,omitempty"`

    // 商品数量
    AuctionNum   int64 `json:"auction_num,omitempty"`

    // 主订单ID
    Tid   int64 `json:"tid,omitempty"`

    // 描述
    Msg   string `json:"msg,omitempty"`

}
