package wdk

// CreateContractInstanceRequest 
type CreateContractInstanceRequest struct {

    // 外部合同id，全局唯一
    OutObjectId   string `json:"out_object_id,omitempty"`

    // 采购者id
    PurchaserId   int64 `json:"purchaser_id,omitempty"`

    // 采购者名称
    PurchaserName   string `json:"purchaser_name,omitempty"`

    // 合同名称
    Title   string `json:"title,omitempty"`

    // 创建人
    Creator   string `json:"creator,omitempty"`

    // 创建人id
    CreatorId   int64 `json:"creator_id,omitempty"`

    // 合同模版
    ContractTemplate   *CreateContractTemplateRequest `json:"contract_template,omitempty"`

}
