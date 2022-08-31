package aesolution

// SkuAttributeInfoQueryRequest 结构体
type SkuAttributeInfoQueryRequest struct {
	// merchant&#39;s category ID
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// aliexpress category ID. aliexpress_category_id and category_id could not be both empty.
	AliexpressCategoryId int64 `json:"aliexpress_category_id,omitempty" xml:"aliexpress_category_id,omitempty"`
}
