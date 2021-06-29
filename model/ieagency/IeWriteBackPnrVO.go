package ieagency

// IeWriteBackPnrVO 
type IeWriteBackPnrVO struct {

    // 预定订单pnr信息
    
    BookPnrVos   []IeBookPnrVo `json:"book_pnr_vos,omitempty" xml:"book_pnr_vos,omitempty"`
    

    // 预定订单ID
    
    BookOrderId   int64 `json:"book_order_id,omitempty" xml:"book_order_id,omitempty"`
    

}
