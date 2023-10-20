package wdk

import (
	"sync"
)

// CreateContractInstanceRequest 结构体
type CreateContractInstanceRequest struct {
	// 外部合同id，全局唯一
	OutObjectId string `json:"out_object_id,omitempty" xml:"out_object_id,omitempty"`
	// 采购者名称
	PurchaserName string `json:"purchaser_name,omitempty" xml:"purchaser_name,omitempty"`
	// 合同名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 创建人
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 采购者id
	PurchaserId int64 `json:"purchaser_id,omitempty" xml:"purchaser_id,omitempty"`
	// 创建人id
	CreatorId int64 `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// 合同模版
	ContractTemplate *CreateContractTemplateRequest `json:"contract_template,omitempty" xml:"contract_template,omitempty"`
}

var poolCreateContractInstanceRequest = sync.Pool{
	New: func() any {
		return new(CreateContractInstanceRequest)
	},
}

// GetCreateContractInstanceRequest() 从对象池中获取CreateContractInstanceRequest
func GetCreateContractInstanceRequest() *CreateContractInstanceRequest {
	return poolCreateContractInstanceRequest.Get().(*CreateContractInstanceRequest)
}

// ReleaseCreateContractInstanceRequest 释放CreateContractInstanceRequest
func ReleaseCreateContractInstanceRequest(v *CreateContractInstanceRequest) {
	v.OutObjectId = ""
	v.PurchaserName = ""
	v.Title = ""
	v.Creator = ""
	v.PurchaserId = 0
	v.CreatorId = 0
	v.ContractTemplate = nil
	poolCreateContractInstanceRequest.Put(v)
}
