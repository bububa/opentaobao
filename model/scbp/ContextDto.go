package scbp

// ContextDto 
type ContextDto struct {
    // login_id
    LoginId   string `json:"login_id,omitempty" xml:"login_id,omitempty"`
    // is_admin
    IsAdmin   bool `json:"is_admin,omitempty" xml:"is_admin,omitempty"`
    // is_top
    IsTop   bool `json:"is_top,omitempty" xml:"is_top,omitempty"`
    // from
    From   string `json:"from,omitempty" xml:"from,omitempty"`
    // service_type
    ServiceType   string `json:"service_type,omitempty" xml:"service_type,omitempty"`
    // ip
    Ip   string `json:"ip,omitempty" xml:"ip,omitempty"`
}
