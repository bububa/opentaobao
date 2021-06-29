package cainiaohandover

// OpenHandoverCommitResponse 
type OpenHandoverCommitResponse struct {
    // 交接批次号，即交接单id
    HandoverOrderId   int64 `json:"handover_order_id,omitempty" xml:"handover_order_id,omitempty"`
    // 交接物编码，即大包LP号
    HandoverContentCode   string `json:"handover_content_code,omitempty" xml:"handover_content_code,omitempty"`
    // 交接物id，即大包id
    HandoverContentId   int64 `json:"handover_content_id,omitempty" xml:"handover_content_id,omitempty"`
}
