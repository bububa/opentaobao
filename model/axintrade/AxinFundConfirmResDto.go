package axintrade

// AxinFundConfirmResDto 
type AxinFundConfirmResDto struct {

    // 外部订单号
    
    OuterOrderId   string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
    

    // 结算金额
    
    SettleAmount   int64 `json:"settle_amount,omitempty" xml:"settle_amount,omitempty"`
    

}
