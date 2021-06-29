package fundplatform

// CardActiveRequest 
type CardActiveRequest struct {

    // 需要激活的卡号。卡号与外部订单号不可以同时为空。
    
    CardNos   []string `json:"card_nos,omitempty" xml:"card_nos>string,omitempty"`
    

    // 在制卡时传入的外部订单号,卡号与外部订单号不可以同时为空
    
    FetchOutBizId   string `json:"fetch_out_biz_id,omitempty" xml:"fetch_out_biz_id,omitempty"`
    

    // 操作人ID，用于后续追踪
    
    OptUserId   int64 `json:"opt_user_id,omitempty" xml:"opt_user_id,omitempty"`
    

    // 当前激活操作的幂等号，请保证不重复
    
    OutBizId   string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
    

}
