package ascpchannel

// BatchChannelInventoryQuery 结构体
type BatchChannelInventoryQuery struct {
	// 产品Id
	ProductIds []string `json:"product_ids,omitempty" xml:"product_ids>string,omitempty"`
	// 二级渠道
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 默认不传
	ProductType string `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 一级渠道
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 省
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
}
