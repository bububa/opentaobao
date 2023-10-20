package icbudropshipping

import (
	"sync"
)

// TrackingEvent 结构体
type TrackingEvent struct {
	// event code
	EventCode string `json:"event_code,omitempty" xml:"event_code,omitempty"`
	// event location
	EventLocation string `json:"event_location,omitempty" xml:"event_location,omitempty"`
	// event name
	EventName string `json:"event_name,omitempty" xml:"event_name,omitempty"`
	// event time
	EventTime string `json:"event_time,omitempty" xml:"event_time,omitempty"`
}

var poolTrackingEvent = sync.Pool{
	New: func() any {
		return new(TrackingEvent)
	},
}

// GetTrackingEvent() 从对象池中获取TrackingEvent
func GetTrackingEvent() *TrackingEvent {
	return poolTrackingEvent.Get().(*TrackingEvent)
}

// ReleaseTrackingEvent 释放TrackingEvent
func ReleaseTrackingEvent(v *TrackingEvent) {
	v.EventCode = ""
	v.EventLocation = ""
	v.EventName = ""
	v.EventTime = ""
	poolTrackingEvent.Put(v)
}
