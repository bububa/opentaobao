package alihealth2

// DentalStoreDto 
type DentalStoreDto struct {

    // itemList
    
    ItemList   []DentalItemDto `json:"item_list,omitempty" xml:"item_list,omitempty"`
    

    // storeName
    
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    

    // storeId
    
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    

}
