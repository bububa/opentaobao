package wdk

// ApplyReverseReq 
/* model for simplify = false
type ApplyReverseReq struct {

    // wdk子单list
    
    BizOrderIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"biz_order_ids,omitempty"`
    

    // 礼品卡号
    
    GiftCardNos  struct {
        String  []string `json:"string,omitempty"`
    } `json:"gift_card_nos,omitempty"`
    

    // 操作者
    
    Operator  *struct {
        OperatorVo  *OperatorVo `json:"operator_vo,omitempty"`
    } `json:"operator,omitempty"`
    

    // 门店id
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 外部子单list(biz_order_ids与sub_out_order_ids 二选一)
    
    SubOutOrderIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sub_out_order_ids,omitempty"`
    

}
*/

// ApplyReverseReq 
type ApplyReverseReq struct {

    // wdk子单list
    BizOrderIds   []int64 `json:"biz_order_ids,omitempty"`

    // 礼品卡号
    GiftCardNos   []string `json:"gift_card_nos,omitempty"`

    // 操作者
    Operator   *OperatorVo `json:"operator,omitempty"`

    // 门店id
    StoreId   string `json:"store_id,omitempty"`

    // 外部子单list(biz_order_ids与sub_out_order_ids 二选一)
    SubOutOrderIds   []string `json:"sub_out_order_ids,omitempty"`

}
