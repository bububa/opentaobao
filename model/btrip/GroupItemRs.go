package btrip

// GroupItemRs 
type GroupItemRs struct {

    // 行程信息
    
    RouteInfo   *RouteInfoRs `json:"route_info,omitempty" xml:"route_info,omitempty"`
    

    // 成人、儿童、老人对应的报价信息
    
    ShoppingItems   []PassengerFlightShoppingItemRs `json:"shopping_items,omitempty" xml:"shopping_items,omitempty"`
    

}
