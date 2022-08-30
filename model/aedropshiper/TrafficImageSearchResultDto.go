package aedropshiper

// TrafficImageSearchResultDto 结构体
type TrafficImageSearchResultDto struct {
	// products
	Products []TrafficImageProductDto `json:"products,omitempty" xml:"products>traffic_image_product_dto,omitempty"`
}
