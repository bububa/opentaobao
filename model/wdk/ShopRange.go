package wdk

// ShopRange 结构体
type ShopRange struct {
	// 经纬度点
	Points []Point `json:"points,omitempty" xml:"points>point,omitempty"`
}
