package wdk

// ApplyReverseReq 
type ApplyReverseReq struct {

    // wdk子单list
    BizOrderIds   []Number `json:"biz_order_ids,omitempty"`

    // 礼品卡号
    GiftCardNos   []String `json:"gift_card_nos,omitempty"`

    // 操作者
    Operator   *OperatorVo `json:"operator,omitempty"`

    // 门店id
    StoreId   string `json:"store_id,omitempty"`

    // 外部子单list(biz_order_ids与sub_out_order_ids 二选一)
    SubOutOrderIds   []String `json:"sub_out_order_ids,omitempty"`

}
