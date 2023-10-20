package travel

import (
	"sync"
)

// PontusTravelItemTrafficInfo 结构体
type PontusTravelItemTrafficInfo struct {
	// 详细交通信息结构。【注意】当traffic_type选择飞机或火车时，该字段为必填，汽车或轮船时该字段不用填。
	Traffics []PontusTravelItemTraffic `json:"traffics,omitempty" xml:"traffics>pontus_travel_item_traffic,omitempty"`
	// 必填，交通类型。1/2/3/4 分别对应 飞机/火车/汽车/船
	TrafficType int64 `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
}

var poolPontusTravelItemTrafficInfo = sync.Pool{
	New: func() any {
		return new(PontusTravelItemTrafficInfo)
	},
}

// GetPontusTravelItemTrafficInfo() 从对象池中获取PontusTravelItemTrafficInfo
func GetPontusTravelItemTrafficInfo() *PontusTravelItemTrafficInfo {
	return poolPontusTravelItemTrafficInfo.Get().(*PontusTravelItemTrafficInfo)
}

// ReleasePontusTravelItemTrafficInfo 释放PontusTravelItemTrafficInfo
func ReleasePontusTravelItemTrafficInfo(v *PontusTravelItemTrafficInfo) {
	v.Traffics = v.Traffics[:0]
	v.TrafficType = 0
	poolPontusTravelItemTrafficInfo.Put(v)
}
