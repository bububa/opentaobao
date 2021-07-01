package travel

// GroupTourTrafficInfo 结构体
type GroupTourTrafficInfo struct {
	// 交通类型 1：飞机， 2：火车，3：汽车，4：轮船
	TransportWay int64 `json:"transport_way,omitempty" xml:"transport_way,omitempty"`
	// 第几组交通信息，每一组交通信息包含一组去程交通和返程交通，当在页> 面上点击【添加交通信息】按钮后，就会出现第二组交通信息，第一组交>通信息group=1，第二组交通信息group取值为2，以此类推
	Group int64 `json:"group,omitempty" xml:"group,omitempty"`
	// 是否直飞，飞机选填，1-直飞；0-不是直飞
	NonStop int64 `json:"non_stop,omitempty" xml:"non_stop,omitempty"`
	// 交通说明，针对交通类型是汽车，轮船和其他
	TrafficDesc string `json:"traffic_desc,omitempty" xml:"traffic_desc,omitempty"`
	// 到达时间，当地时间HH:mm
	ArrivalTime string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// 出发时间，当地时间HH:mm
	DepartureTime string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// 到达城市
	Destination string `json:"destination,omitempty" xml:"destination,omitempty"`
	// 出发城市
	Departure string `json:"departure,omitempty" xml:"departure,omitempty"`
	// 飞机机型，飞机选填
	PlaneType string `json:"plane_type,omitempty" xml:"plane_type,omitempty"`
	// 交通公司名，飞机选填
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 参考班次号，飞机需要填航班号，火车需要填车次号，汽车和船可不填
	TrafficNo string `json:"traffic_no,omitempty" xml:"traffic_no,omitempty"`
	// 第几天
	Day int64 `json:"day,omitempty" xml:"day,omitempty"`
	// 是否经停
	StopOver bool `json:"stop_over,omitempty" xml:"stop_over,omitempty"`
	// 经停城市
	StopCity string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// 是否是"非红眼航班"。【红眼航班】定义：凌晨一点至六点起飞，且飞行时间少于少于正常睡眠需求（8小时）的航班。
	IsNonRedEyeFlight bool `json:"is_non_red_eye_flight,omitempty" xml:"is_non_red_eye_flight,omitempty"`
}
