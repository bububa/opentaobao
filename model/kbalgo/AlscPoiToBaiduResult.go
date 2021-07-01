package kbalgo

// AlscPoiToBaiduResult 结构体
type AlscPoiToBaiduResult struct {
	// poi总量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// datas
	Datas []PoiToBaiduData `json:"datas,omitempty" xml:"datas>poi_to_baidu_data,omitempty"`
	// 附加信息或错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态码：0-success，1-fail
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
