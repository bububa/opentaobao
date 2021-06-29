package topoaid

// OrderMerge 
type OrderMerge struct {

    // 订单ID
    
    Tid   string `json:"tid,omitempty" xml:"tid,omitempty"`
    

    // 收件人ID (Open Addressee ID)，长度在128个字符之内。
    
    Oaid   string `json:"oaid,omitempty" xml:"oaid,omitempty"`
    

}
