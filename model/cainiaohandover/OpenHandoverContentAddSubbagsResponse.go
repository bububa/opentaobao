package cainiaohandover

// OpenHandoverContentAddSubbagsResponse 结构体
type OpenHandoverContentAddSubbagsResponse struct {
	// 追加大包列表
	SubbagHandoverContentList []HandoverContentAddSubbagsDto `json:"subbag_handover_content_list,omitempty" xml:"subbag_handover_content_list>handover_content_add_subbags_dto,omitempty"`
}
