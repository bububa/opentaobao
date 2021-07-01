package koubeimall

// ItemStoreDto 结构体
type ItemStoreDto struct {
	// 门店LOGO
	StoreLogo string `json:"store_logo,omitempty" xml:"store_logo,omitempty"`
	// 门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 门店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
