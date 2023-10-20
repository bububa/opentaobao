package travel

import (
	"sync"
)

// ItemTrafficInfo 结构体
type ItemTrafficInfo struct {
	// 详细交通信息结构。【注意】当traffic_type选择飞机或火车时，该字段为必填，汽车或轮船时该字段不用填。
	Traffics []ItemTraffic `json:"traffics,omitempty" xml:"traffics>item_traffic,omitempty"`
	// 必填，交通类型。1/2/3/4 分别对应 飞机/火车/汽车/船
	TrafficType int64 `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
}

var poolItemTrafficInfo = sync.Pool{
	New: func() any {
		return new(ItemTrafficInfo)
	},
}

// GetItemTrafficInfo() 从对象池中获取ItemTrafficInfo
func GetItemTrafficInfo() *ItemTrafficInfo {
	return poolItemTrafficInfo.Get().(*ItemTrafficInfo)
}

// ReleaseItemTrafficInfo 释放ItemTrafficInfo
func ReleaseItemTrafficInfo(v *ItemTrafficInfo) {
	v.Traffics = v.Traffics[:0]
	v.TrafficType = 0
	poolItemTrafficInfo.Put(v)
}
