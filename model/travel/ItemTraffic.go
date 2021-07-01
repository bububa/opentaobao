package travel

// ItemTraffic 结构体
type ItemTraffic struct {
	// 关联的套餐id
	RelatedPackageId int64 `json:"related_package_id,omitempty" xml:"related_package_id,omitempty"`
	// 交通公司名，飞机选填
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 参考班次号，飞机需要填航班号，火车需要填车次号，汽车和船可不填
	TrafficNo string `json:"traffic_no,omitempty" xml:"traffic_no,omitempty"`
	// 飞机机型，飞机选填
	PlaneType string `json:"plane_type,omitempty" xml:"plane_type,omitempty"`
	// 到达城市
	Destination string `json:"destination,omitempty" xml:"destination,omitempty"`
	// 出发时间，当地时间HH:mm
	DepartureTime string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// 出发城市
	Departure string `json:"departure,omitempty" xml:"departure,omitempty"`
	// 到达时间，当地时间HH:mm
	ArrivalTime string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
}
