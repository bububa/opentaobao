package wdk

// MerchantProductResponse 
/* model for simplify = false
type MerchantProductResponse struct {

    // 货品id
    
    ScIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"sc_ids,omitempty"`
    

    // ["123","456"]
    
    ItemId   int64 `json:"item_id,omitempty"`
    

}
*/

// MerchantProductResponse 
type MerchantProductResponse struct {

    // 货品id
    ScIds   []int64 `json:"sc_ids,omitempty"`

    // ["123","456"]
    ItemId   int64 `json:"item_id,omitempty"`

}
