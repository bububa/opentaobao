package cainiaohandover

// HandoverContentAddSubbagsDto 结构体
type HandoverContentAddSubbagsDto struct {
	// 追加的大包LP
	SubbagHandoverContentCode string `json:"subbag_handover_content_code,omitempty" xml:"subbag_handover_content_code,omitempty"`
	// 交接单id
	HandoverOrderId int64 `json:"handover_order_id,omitempty" xml:"handover_order_id,omitempty"`
	// 追加的大包id
	SubbagHandoverContentId int64 `json:"subbag_handover_content_id,omitempty" xml:"subbag_handover_content_id,omitempty"`
}
