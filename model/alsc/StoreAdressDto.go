package alsc

// StoreAdressDto 
type StoreAdressDto struct {
    // 省份
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    // 城市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 区
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 街道
    Town   string `json:"town,omitempty" xml:"town,omitempty"`
    // 详细地址
    DetailAddress   string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
    // 经度
    PosX   string `json:"pos_x,omitempty" xml:"pos_x,omitempty"`
    // 维度
    PosY   string `json:"pos_y,omitempty" xml:"pos_y,omitempty"`
    // 营业面积
    BusinessArea   string `json:"business_area,omitempty" xml:"business_area,omitempty"`
}
