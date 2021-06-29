package tmallnr

// NrtStoreDTO 
type NrtStoreDTO struct {
    // 门店ID
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 门店名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 地址
    ShortAddress   string `json:"short_address,omitempty" xml:"short_address,omitempty"`
    // 纬度
    Lat   string `json:"lat,omitempty" xml:"lat,omitempty"`
    // 经度
    Lng   string `json:"lng,omitempty" xml:"lng,omitempty"`
}
