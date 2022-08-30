package car

// RentCarDepositNode 结构体
type RentCarDepositNode struct {
	// 节点流程
	DepositFlows []RentCarDepositFlow `json:"deposit_flows,omitempty" xml:"deposit_flows>rent_car_deposit_flow,omitempty"`
	// nodeName
	NodeName string `json:"node_name,omitempty" xml:"node_name,omitempty"`
	// status
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
