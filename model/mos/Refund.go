package mos

// Refund 
/* model for simplify = false
type Refund struct {

    // 退货金额
    
    Amount   string `json:"amount,omitempty"`
    

    // 退货数量
    
    Quantity   string `json:"quantity,omitempty"`
    

    // 原小票号
    
    Oldseqno   string `json:"oldseqno,omitempty"`
    

    // 原小票店号
    
    Oldstoreno   string `json:"oldstoreno,omitempty"`
    

    // 当前商品行对应小票的N09记录字符串
    
    Refundpara   string `json:"refundpara,omitempty"`
    

    // 商品行号
    
    Comorder   string `json:"comorder,omitempty"`
    

}
*/

// Refund 
type Refund struct {

    // 退货金额
    Amount   string `json:"amount,omitempty"`

    // 退货数量
    Quantity   string `json:"quantity,omitempty"`

    // 原小票号
    Oldseqno   string `json:"oldseqno,omitempty"`

    // 原小票店号
    Oldstoreno   string `json:"oldstoreno,omitempty"`

    // 当前商品行对应小票的N09记录字符串
    Refundpara   string `json:"refundpara,omitempty"`

    // 商品行号
    Comorder   string `json:"comorder,omitempty"`

}
