package icbudropshipping

import (
	"sync"
)

// LogisticsTracking 结构体
type LogisticsTracking struct {
	// event list
	EventList []TrackingEvent `json:"event_list,omitempty" xml:"event_list>tracking_event,omitempty"`
	// carrier
	Carrier string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// The latest logistics event code
	CurrentEventCode string `json:"current_event_code,omitempty" xml:"current_event_code,omitempty"`
	// tracking number
	TrackingNumber string `json:"tracking_number,omitempty" xml:"tracking_number,omitempty"`
	// tracking url
	TrackingUrl string `json:"tracking_url,omitempty" xml:"tracking_url,omitempty"`
}

var poolLogisticsTracking = sync.Pool{
	New: func() any {
		return new(LogisticsTracking)
	},
}

// GetLogisticsTracking() 从对象池中获取LogisticsTracking
func GetLogisticsTracking() *LogisticsTracking {
	return poolLogisticsTracking.Get().(*LogisticsTracking)
}

// ReleaseLogisticsTracking 释放LogisticsTracking
func ReleaseLogisticsTracking(v *LogisticsTracking) {
	v.EventList = v.EventList[:0]
	v.Carrier = ""
	v.CurrentEventCode = ""
	v.TrackingNumber = ""
	v.TrackingUrl = ""
	poolLogisticsTracking.Put(v)
}
