package travel

// PontusTravelItemTrafficInfo 
type PontusTravelItemTrafficInfo struct {

    // 必填，交通类型。1/2/3/4 分别对应 飞机/火车/汽车/船，其他交通方式传100
    
    TrafficType   int64 `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
    

    // 详细交通信息结构。【注意】当traffic_type选择飞机或火车时，该字段为必填，汽车或轮船时该字段不用填。
    
    Traffics   []PontusTravelItemTraffic `json:"traffics,omitempty" xml:"traffics,omitempty"`
    

    // 描述，小于1500字
    
    TrafficDesc   string `json:"traffic_desc,omitempty" xml:"traffic_desc,omitempty"`
    

}
