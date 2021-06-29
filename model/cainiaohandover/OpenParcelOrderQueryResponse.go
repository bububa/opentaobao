package cainiaohandover

// OpenParcelOrderQueryResponse 
type OpenParcelOrderQueryResponse struct {
    // 交接仓编码，快递揽收场景,大包交接目的地国际分拨
    HandoverWarehouseCode   string `json:"handover_warehouse_code,omitempty" xml:"handover_warehouse_code,omitempty"`
    // 交接仓名称，快递揽收场景,大包交接目的地国际分拨
    HandoverWarehouseName   string `json:"handover_warehouse_name,omitempty" xml:"handover_warehouse_name,omitempty"`
    // 该小包是否已经组包
    HasBeenHandover   bool `json:"has_been_handover,omitempty" xml:"has_been_handover,omitempty"`
    // 关联的交接单ID
    HandoverOrderId   int64 `json:"handover_order_id,omitempty" xml:"handover_order_id,omitempty"`
    // 关联的大包的编码
    HandoverContentCode   string `json:"handover_content_code,omitempty" xml:"handover_content_code,omitempty"`
    // 关联的大包ID
    HandoverContentId   int64 `json:"handover_content_id,omitempty" xml:"handover_content_id,omitempty"`
    // 是否能组包
    CanCreateHandover   bool `json:"can_create_handover,omitempty" xml:"can_create_handover,omitempty"`
}
