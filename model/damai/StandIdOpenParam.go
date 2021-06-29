package damai

// StandIdOpenParam 
type StandIdOpenParam struct {
    // 场次id
    PerformId   int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
    // 看台id
    StandId   int64 `json:"stand_id,omitempty" xml:"stand_id,omitempty"`
    // 系统id
    SystemId   int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
    // 商户密钥
    SupplierSecret   string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
}
