package wdk

// MerchantProductResponse 
type MerchantProductResponse struct {

    // 货品id
    ScIds   []Number `json:"sc_ids,omitempty"`

    // ["123","456"]
    ItemId   int64 `json:"item_id,omitempty"`

}
