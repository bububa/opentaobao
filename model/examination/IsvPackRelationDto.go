package examination

// IsvPackRelationDto 
type IsvPackRelationDto struct {
    // 加项包id
    IsvPackId   string `json:"isv_pack_id,omitempty" xml:"isv_pack_id,omitempty"`
    // 关联加项包id
    RelIsvPackId   string `json:"rel_isv_pack_id,omitempty" xml:"rel_isv_pack_id,omitempty"`
    // 关系：1、互斥.
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
}
