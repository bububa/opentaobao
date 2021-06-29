package feedflow

// AccessAllowedDTO 
type AccessAllowedDTO struct {
    // 是否可以使用，false不可以进行广告投放
    IsAccessAllowed   bool `json:"is_access_allowed,omitempty" xml:"is_access_allowed,omitempty"`
    // 不可以使用的原因
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
}
