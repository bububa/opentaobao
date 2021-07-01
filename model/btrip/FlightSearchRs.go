package btrip

// FlightSearchRs 
type FlightSearchRs struct {
    // 组合商品列表
    ItemList   []GroupItemRs `json:"item_list,omitempty" xml:"item_list>group_item_rs,omitempty"`
}
