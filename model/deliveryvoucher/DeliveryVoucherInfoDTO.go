package deliveryvoucher

// DeliveryVoucherInfoDTO 
type DeliveryVoucherInfoDTO struct {
    // 券信息
    Vouchers   []DeliveryVoucherDTO `json:"vouchers,omitempty" xml:"vouchers>delivery_voucher_dto,omitempty"`
    // 商品id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
