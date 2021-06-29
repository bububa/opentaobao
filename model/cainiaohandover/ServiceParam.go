package cainiaohandover

// ServiceParam 
type ServiceParam struct {
    // DOOR_PICKUP：揽收仓资源、SELF_SEND：自送dropOff
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
}
