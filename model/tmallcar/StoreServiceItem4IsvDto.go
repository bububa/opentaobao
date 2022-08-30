package tmallcar

// StoreServiceItem4IsvDto 结构体
type StoreServiceItem4IsvDto struct {
	// 商品列表
	StoreServiceItems []Item4IsvDto `json:"store_service_items,omitempty" xml:"store_service_items>item4isv_dto,omitempty"`
}
