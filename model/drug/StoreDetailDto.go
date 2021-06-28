package drug

// StoreDetailDto 
/* model for simplify = false
type StoreDetailDto struct {

    // tags
    
    Tags  struct {
        Tags  []Tags `json:"tags,omitempty"`
    } `json:"tags,omitempty"`
    

    // storeDetail
    
    StoreDetail  *struct {
        StoreDto  *StoreDto `json:"store_dto,omitempty"`
    } `json:"store_detail,omitempty"`
    

    // cat
    
    Cats  struct {
        Cat  []Cat `json:"cat,omitempty"`
    } `json:"cats,omitempty"`
    

}
*/

// StoreDetailDto 
type StoreDetailDto struct {

    // tags
    Tags   []Tags `json:"tags,omitempty"`

    // storeDetail
    StoreDetail   *StoreDto `json:"store_detail,omitempty"`

    // cat
    Cats   []Cat `json:"cats,omitempty"`

}
