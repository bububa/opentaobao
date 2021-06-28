package elife

// ConsumeRequest 
type ConsumeRequest struct {

    // 消费金额, 分
    
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`
    

    // 商家的操作流水号, 唯一键
    
    OpId   string `json:"op_id,omitempty" xml:"op_id,omitempty"`
    

    // 收银员名称
    
    Operator   string `json:"operator,omitempty" xml:"operator,omitempty"`
    

    // 门店编号
    
    OuterStoreId   string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
    

    // 码枪读入的手淘条码
    
    PayCode   string `json:"pay_code,omitempty" xml:"pay_code,omitempty"`
    

    // 小票信息
    
    SaleTicket   string `json:"sale_ticket,omitempty" xml:"sale_ticket,omitempty"`
    

}
