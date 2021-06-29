package drug

// StoreDetailDTO 
type StoreDetailDTO struct {
    // tags
    Tags   []Tags `json:"tags,omitempty" xml:"tags>tags,omitempty"`
    // storeDetail
    StoreDetail   *StoreDTO `json:"store_detail,omitempty" xml:"store_detail,omitempty"`
    // cat
    Cats   []Cat `json:"cats,omitempty" xml:"cats>cat,omitempty"`
}
