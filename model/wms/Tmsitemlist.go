package wms

// Tmsitemlist 
type Tmsitemlist struct {
    // 包裹里面商品
    TmsItem   *Tmsitem `json:"tms_item,omitempty" xml:"tms_item,omitempty"`
}
