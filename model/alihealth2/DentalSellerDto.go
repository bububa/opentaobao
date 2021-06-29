package alihealth2

// DentalSellerDTO 
type DentalSellerDTO struct {
    // storeList
    StoreList   []DentalStoreDTO `json:"store_list,omitempty" xml:"store_list>dental_store_dto,omitempty"`
    // sellerName
    SellerName   string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
}
