package wdk

// AxStoreCreateRequest 结构体
type AxStoreCreateRequest struct {
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 区编码
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 市编码
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省编码
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
	// 门店编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 门店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 门店经营状态 ：1 正常 0 关闭
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
