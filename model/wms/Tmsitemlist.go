package wms

// Tmsitemlist 结构体
type Tmsitemlist struct {
	// 包裹里面商品
	TmsItem *Tmsitem `json:"tms_item,omitempty" xml:"tms_item,omitempty"`
}
