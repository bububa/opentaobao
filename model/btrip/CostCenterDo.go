package btrip

// CostCenterDO 
type CostCenterDO struct {
    // 成本中心ID
    CostCenterId   int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
    // 成本中心名称
    CostCenterTitle   string `json:"cost_center_title,omitempty" xml:"cost_center_title,omitempty"`
    // 成本中心编码
    CostCenterNumber   string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
    // 第三方成本中心id
    ThirdCostCenterId   string `json:"third_cost_center_id,omitempty" xml:"third_cost_center_id,omitempty"`
}
