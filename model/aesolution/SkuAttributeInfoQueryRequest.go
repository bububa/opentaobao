package aesolution

// SkuAttributeInfoQueryRequest 
type SkuAttributeInfoQueryRequest struct {
    // aliexpress category ID. aliexpress_category_id and category_id could not be both empty.
    AliexpressCategoryId   int64 `json:"aliexpress_category_id,omitempty" xml:"aliexpress_category_id,omitempty"`
    // merchant's category ID
    CategoryId   string `json:"category_id,omitempty" xml:"category_id,omitempty"`
}
