package hotelhstdf

// HotelExportParam 结构体
type HotelExportParam struct {
	// 标准酒店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 标准酒店所在城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 酒店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 电话号码
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 经纬度
	Jwd string `json:"jwd,omitempty" xml:"jwd,omitempty"`
	// 相似度分数
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 英文名称
	NameE string `json:"name_e,omitempty" xml:"name_e,omitempty"`
	// 英文地址
	EnAddr string `json:"en_addr,omitempty" xml:"en_addr,omitempty"`
	// 曾用名
	UsedName string `json:"used_name,omitempty" xml:"used_name,omitempty"`
	// 酒店状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 卖家酒店id
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 酒店状态0,正常，有报价可在线售卖，-1，待上线，-3 下架，-1，-3需要上线时联系平台处理
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
}
