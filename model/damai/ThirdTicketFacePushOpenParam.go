package damai

// ThirdTicketFacePushOpenParam 
type ThirdTicketFacePushOpenParam struct {
    // 票面id
    FaceId   int64 `json:"face_id,omitempty" xml:"face_id,omitempty"`
    // 票面模式
    FaceMode   int64 `json:"face_mode,omitempty" xml:"face_mode,omitempty"`
    // 票面类型
    FrontType   int64 `json:"front_type,omitempty" xml:"front_type,omitempty"`
    // 票纸版式id
    PaperFormatId   int64 `json:"paper_format_id,omitempty" xml:"paper_format_id,omitempty"`
    // 场次id
    PerformId   int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
    // 推送时间
    PushTime   string `json:"push_time,omitempty" xml:"push_time,omitempty"`
    // 商户密钥
    SupplierSecret   string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
    // 系统id
    SystemId   int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
}
