package xhotelonlineorder

// TopDomesticPriceInfoDo 
type TopDomesticPriceInfoDo struct {
    // createOrderDailyPrice
    CreateOrderDailyPrice   *HbsDailyPriceDo `json:"create_order_daily_price,omitempty" xml:"create_order_daily_price,omitempty"`
    // taxPrice
    TaxPrice   int64 `json:"tax_price,omitempty" xml:"tax_price,omitempty"`
    // roomsPrice
    RoomsPrice   int64 `json:"rooms_price,omitempty" xml:"rooms_price,omitempty"`
    // servicePrice
    ServicePrice   int64 `json:"service_price,omitempty" xml:"service_price,omitempty"`
    // otherFee
    OtherFee   int64 `json:"other_fee,omitempty" xml:"other_fee,omitempty"`
    // bookingRoomsPrice
    BookingRoomsPrice   int64 `json:"booking_rooms_price,omitempty" xml:"booking_rooms_price,omitempty"`
    // bookingTaxPrice
    BookingTaxPrice   int64 `json:"booking_tax_price,omitempty" xml:"booking_tax_price,omitempty"`
    // bookingServicePrice
    BookingServicePrice   int64 `json:"booking_service_price,omitempty" xml:"booking_service_price,omitempty"`
    // settleOrderDailyPrice
    SettleOrderDailyPrice   *HbsDailyPriceDo `json:"settle_order_daily_price,omitempty" xml:"settle_order_daily_price,omitempty"`
    // basePrice
    BasePrice   int64 `json:"base_price,omitempty" xml:"base_price,omitempty"`
    // bookingBasePrice
    BookingBasePrice   int64 `json:"booking_base_price,omitempty" xml:"booking_base_price,omitempty"`
}
