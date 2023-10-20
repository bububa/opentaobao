package car

import (
	"sync"
)

// RentCarDepositNode 结构体
type RentCarDepositNode struct {
	// 节点流程
	DepositFlows []RentCarDepositFlow `json:"deposit_flows,omitempty" xml:"deposit_flows>rent_car_deposit_flow,omitempty"`
	// nodeName
	NodeName string `json:"node_name,omitempty" xml:"node_name,omitempty"`
	// status
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolRentCarDepositNode = sync.Pool{
	New: func() any {
		return new(RentCarDepositNode)
	},
}

// GetRentCarDepositNode() 从对象池中获取RentCarDepositNode
func GetRentCarDepositNode() *RentCarDepositNode {
	return poolRentCarDepositNode.Get().(*RentCarDepositNode)
}

// ReleaseRentCarDepositNode 释放RentCarDepositNode
func ReleaseRentCarDepositNode(v *RentCarDepositNode) {
	v.DepositFlows = v.DepositFlows[:0]
	v.NodeName = ""
	v.Status = 0
	poolRentCarDepositNode.Put(v)
}
