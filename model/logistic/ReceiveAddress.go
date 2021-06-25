package logistic

// ReceiveAddress 
type ReceiveAddress struct {

    // 镇/街道
    TownName   string `json:"town_name,omitempty"`

    // 详细地址
    AddressDetail   string `json:"address_detail,omitempty"`

    // 市
    CityName   string `json:"city_name,omitempty"`

    // 区
    AreaName   string `json:"area_name,omitempty"`

    // 省
    ProvinceName   string `json:"province_name,omitempty"`

}
