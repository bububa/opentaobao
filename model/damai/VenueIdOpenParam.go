package damai

// VenueIdOpenParam 
type VenueIdOpenParam struct {
    // 商家id
    CpId   int64 `json:"cp_id,omitempty" xml:"cp_id,omitempty"`
    // 推送接口返回的id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 操作员
    Operator   string `json:"operator,omitempty" xml:"operator,omitempty"`
    // 操作员id
    OperatorId   int64 `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    // 商家id
    SystemId   int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
    // 场馆id
    VenueId   int64 `json:"venue_id,omitempty" xml:"venue_id,omitempty"`
    // 商家密钥，监权使用
    SupplierSecret   string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
}
