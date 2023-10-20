package tmallcar

// StoreServiceItem4isvDto 结构体
type StoreServiceItem4isvDto struct {
	// 商品列表
	StoreServiceItems []Item4isvDto `json:"store_service_items,omitempty" xml:"store_service_items>item4isv_dto,omitempty"`
}
