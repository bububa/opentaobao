package alihealth2

// DentalStoreDTO 
type DentalStoreDTO struct {
    // itemList
    ItemList   []DentalItemDTO `json:"item_list,omitempty" xml:"item_list>dental_item_dto,omitempty"`
    // storeName
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    // storeId
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
