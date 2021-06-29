package btrip

// PassengerFlightShoppingItemRs 
type PassengerFlightShoppingItemRs struct {

    // 乘客类型
    
    PassengerType   string `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
    

    // 机票报价
    
    ShoppingItem   *FlightShoppingItemRs `json:"shopping_item,omitempty" xml:"shopping_item,omitempty"`
    

}
