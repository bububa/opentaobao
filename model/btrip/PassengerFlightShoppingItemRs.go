package btrip

// PassengerFlightShoppingItemRS 
type PassengerFlightShoppingItemRS struct {
    // 乘客类型
    PassengerType   string `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
    // 机票报价
    ShoppingItem   *FlightShoppingItemRS `json:"shopping_item,omitempty" xml:"shopping_item,omitempty"`
}
