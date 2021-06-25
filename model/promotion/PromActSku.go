package promotion

// PromActSku 
type PromActSku struct {

    // 商家逻辑分组序号
    LogicGroupNumber   int64 `json:"logic_group_number,omitempty"`

    // 参与者id
    ParticipateId   string `json:"participate_id,omitempty"`

    // 参与者名称
    ParticipateName   string `json:"participate_name,omitempty"`

    // 参与者类型 SKU_CODE(1, "商品skuCode"), SHOP(2, "店铺"), CATEGORY(3, "品类")
    ParticipateType   int64 `json:"participate_type,omitempty"`

}
