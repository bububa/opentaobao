package alihealth2

// DentalSellerDto 结构体
type DentalSellerDto struct {
	// storeList
	StoreList []DentalStoreDto `json:"store_list,omitempty" xml:"store_list>dental_store_dto,omitempty"`
	// sellerName
	SellerName string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
}
