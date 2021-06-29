package alihealth2

// DentalSellerDto 
type DentalSellerDto struct {

    // storeList
    
    StoreList   []DentalStoreDto `json:"store_list,omitempty" xml:"store_list,omitempty"`
    

    // sellerName
    
    SellerName   string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
    

}
