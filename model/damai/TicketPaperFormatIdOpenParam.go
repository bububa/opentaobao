package damai

// TicketPaperFormatIdOpenParam 
type TicketPaperFormatIdOpenParam struct {

    // 系统id
    
    SystemId   int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
    

    // 票纸版式id
    
    TicketPaperFormatId   int64 `json:"ticket_paper_format_id,omitempty" xml:"ticket_paper_format_id,omitempty"`
    

    // 商户密钥
    
    SupplierSecret   string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
    

}
