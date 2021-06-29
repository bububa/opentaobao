package xhotelonlineorder

// HotelOrderSellerDO 
type HotelOrderSellerDO struct {
    // 支付类型
    Payment   *OrderSellerPaymentDO `json:"payment,omitempty" xml:"payment,omitempty"`
    // 卖家结账单信息
    SettlePayment   *OrderSettleDO `json:"settle_payment,omitempty" xml:"settle_payment,omitempty"`
}
