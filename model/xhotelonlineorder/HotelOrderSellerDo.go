package xhotelonlineorder

// HotelOrderSellerDo 
type HotelOrderSellerDo struct {
    // 支付类型
    Payment   *OrderSellerPaymentDo `json:"payment,omitempty" xml:"payment,omitempty"`
    // 卖家结账单信息
    SettlePayment   *OrderSettleDo `json:"settle_payment,omitempty" xml:"settle_payment,omitempty"`
}
