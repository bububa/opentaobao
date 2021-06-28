package mos

// AlibabaMjOcBigposBanksaleQueryData 
/* model for simplify = false
type AlibabaMjOcBigposBanksaleQueryData struct {

    // 行号
    
    RowNo   int64 `json:"row_no,omitempty"`
    

    // 交易时间
    
    OperTime   string `json:"oper_time,omitempty"`
    

    // 查询流水号
    
    PosTraceNo   string `json:"pos_trace_no,omitempty"`
    

    // 银联系统参考号
    
    Refnum   string `json:"refnum,omitempty"`
    

    // 交易金额
    
    Amount   int64 `json:"amount,omitempty"`
    

    // 已调账金额
    
    AdjustedAmount   int64 `json:"adjusted_amount,omitempty"`
    

    // 小票号后7位
    
    Fphm   string `json:"fphm,omitempty"`
    

}
*/

// AlibabaMjOcBigposBanksaleQueryData 
type AlibabaMjOcBigposBanksaleQueryData struct {

    // 行号
    RowNo   int64 `json:"row_no,omitempty"`

    // 交易时间
    OperTime   string `json:"oper_time,omitempty"`

    // 查询流水号
    PosTraceNo   string `json:"pos_trace_no,omitempty"`

    // 银联系统参考号
    Refnum   string `json:"refnum,omitempty"`

    // 交易金额
    Amount   int64 `json:"amount,omitempty"`

    // 已调账金额
    AdjustedAmount   int64 `json:"adjusted_amount,omitempty"`

    // 小票号后7位
    Fphm   string `json:"fphm,omitempty"`

}
