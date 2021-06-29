package damai

// TicketIdOpenParam 
type TicketIdOpenParam struct {

    // 场次id
    
    PerformId   int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
    

    // 系统id
    
    SystemId   int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
    

    // 票单id
    
    VoucherId   int64 `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
    

    // 商家密钥
    
    SupplierSecret   string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
    

}
