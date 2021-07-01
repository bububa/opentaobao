package xhotelitem

// Coordinate 结构体
type Coordinate struct {
	// 批次号
	BatchId int64 `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 外部经纬度标识id，可以是酒店或城市的id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 飞猪国家编码
	Country int64 `json:"country,omitempty" xml:"country,omitempty"`
	// 飞猪城市编码
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 飞猪城市中文名称
	CityCnName string `json:"city_cn_name,omitempty" xml:"city_cn_name,omitempty"`
	// 飞猪城市英文名称
	CityEnName string `json:"city_en_name,omitempty" xml:"city_en_name,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
}
