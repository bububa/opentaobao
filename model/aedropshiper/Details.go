package aedropshiper

// Details 结构体
type Details struct {
	// eventDesc
	EventDesc string `json:"event_desc,omitempty" xml:"event_desc,omitempty"`
	// signedName
	SignedName string `json:"signed_name,omitempty" xml:"signed_name,omitempty"`
	// status
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// address
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// eventDate
	EventDate string `json:"event_date,omitempty" xml:"event_date,omitempty"`
}
