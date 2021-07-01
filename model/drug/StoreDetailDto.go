package drug

// StoreDetailDto 
type StoreDetailDto struct {
    // tags
    Tags   []Tags `json:"tags,omitempty" xml:"tags>tags,omitempty"`
    // storeDetail
    StoreDetail   *StoreDto `json:"store_detail,omitempty" xml:"store_detail,omitempty"`
    // cat
    Cats   []Cat `json:"cats,omitempty" xml:"cats>cat,omitempty"`
}
