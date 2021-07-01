package btrip

// FlightSearchRS 
type FlightSearchRS struct {
    // 组合商品列表
    ItemList   []GroupItemRS `json:"item_list,omitempty" xml:"item_list>group_item_rs,omitempty"`
}
