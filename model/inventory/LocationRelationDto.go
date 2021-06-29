package inventory

// LocationRelationDTO 
type LocationRelationDTO struct {
    // 状态  0 正常  -1 删除
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 实体类型 2：仓库 6：门店
    TargetInvStoreType   int64 `json:"target_inv_store_type,omitempty" xml:"target_inv_store_type,omitempty"`
    // 实体code
    TargetStoreCode   string `json:"target_store_code,omitempty" xml:"target_store_code,omitempty"`
    // 实体类型 2：仓库 6：门店
    SourceInvStoreType   int64 `json:"source_inv_store_type,omitempty" xml:"source_inv_store_type,omitempty"`
    // 实体code
    SourceStoreCode   string `json:"source_store_code,omitempty" xml:"source_store_code,omitempty"`
}
