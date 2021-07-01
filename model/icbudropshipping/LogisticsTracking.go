package icbudropshipping

// LogisticsTracking 结构体
type LogisticsTracking struct {
	// carrier
	Carrier string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// The latest logistics event code
	CurrentEventCode string `json:"current_event_code,omitempty" xml:"current_event_code,omitempty"`
	// event list
	EventList []TrackingEvent `json:"event_list,omitempty" xml:"event_list>tracking_event,omitempty"`
	// tracking number
	TrackingNumber string `json:"tracking_number,omitempty" xml:"tracking_number,omitempty"`
}
