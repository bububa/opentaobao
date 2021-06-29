package alscmerchant

// MerchantTicketUseDetail 
type MerchantTicketUseDetail struct {

    // 券核销流水号
    
    TicketTransId   string `json:"ticket_trans_id,omitempty" xml:"ticket_trans_id,omitempty"`
    

    // 核销的券码
    
    TicketCode   string `json:"ticket_code,omitempty" xml:"ticket_code,omitempty"`
    

    // 用户购买券的时候实际支付的金额，单位为元，精确到小数点后两位
    
    BuyerPayAmount   string `json:"buyer_pay_amount,omitempty" xml:"buyer_pay_amount,omitempty"`
    

    // 商家实收金额，单位为元，精确到小数点后两位
    
    ReceiptAmount   string `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
    

    // 商家优惠金额，单位为元，精确到小数点后两位
    
    DiscountAmount   string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
    

    // 口碑补贴金额，单位为元，精确到小数点后两位
    
    MerchantSubsidyAmount   string `json:"merchant_subsidy_amount,omitempty" xml:"merchant_subsidy_amount,omitempty"`
    

    // 交易中可给用户开具发票的金额，单位为元，精确到小数点后两位
    
    InvoiceAmount   string `json:"invoice_amount,omitempty" xml:"invoice_amount,omitempty"`
    

}
