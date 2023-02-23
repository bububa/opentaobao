package aedropshiper

// AeStoreInfo 结构体
type AeStoreInfo struct {
	// Shop name
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// Product description, 1-5 stars
	ItemAsDescribedRating string `json:"item_as_described_rating,omitempty" xml:"item_as_described_rating,omitempty"`
	// Seller service, 1-5 stars
	CommunicationRating string `json:"communication_rating,omitempty" xml:"communication_rating,omitempty"`
	// Logistics, 1-5 stars
	ShippingSpeedRating string `json:"shipping_speed_rating,omitempty" xml:"shipping_speed_rating,omitempty"`
	// Store ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
