package aedropshiper

// AeStoreSimpleInfo 结构体
type AeStoreSimpleInfo struct {
	// Store name
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// Store address
	StoreUrl string `json:"store_url,omitempty" xml:"store_url,omitempty"`
	// Store ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
