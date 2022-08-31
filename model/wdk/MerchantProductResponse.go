package wdk

// MerchantProductResponse 结构体
type MerchantProductResponse struct {
	// 货品id
	ScIds []int64 `json:"sc_ids,omitempty" xml:"sc_ids>int64,omitempty"`
	// [&#34;123&#34;,&#34;456&#34;]
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
