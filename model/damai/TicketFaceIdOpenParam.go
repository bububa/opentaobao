package damai

// TicketFaceIdOpenParam 结构体
type TicketFaceIdOpenParam struct {
	// 票面id
	FaceId int64 `json:"face_id,omitempty" xml:"face_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
}
