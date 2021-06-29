package wdk

// MerchantProductResponse 
type MerchantProductResponse struct {
    // 货品id
    ScIds   []int64 `json:"sc_ids,omitempty" xml:"sc_ids>int64,omitempty"`
    // ["123","456"]
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
