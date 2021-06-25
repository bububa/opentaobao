package travel

// ItemTraffic 
type ItemTraffic struct {

    // 关联的套餐id
    RelatedPackageId   int64 `json:"related_package_id,omitempty"`

    // 交通公司名，飞机选填
    Vendor   string `json:"vendor,omitempty"`

    // 参考班次号，飞机需要填航班号，火车需要填车次号，汽车和船可不填
    TrafficNo   string `json:"traffic_no,omitempty"`

    // 飞机机型，飞机选填
    PlaneType   string `json:"plane_type,omitempty"`

    // 到达城市
    Destination   string `json:"destination,omitempty"`

    // 出发时间，当地时间HH:mm
    DepartureTime   string `json:"departure_time,omitempty"`

    // 出发城市
    Departure   string `json:"departure,omitempty"`

    // 到达时间，当地时间HH:mm
    ArrivalTime   string `json:"arrival_time,omitempty"`

}
