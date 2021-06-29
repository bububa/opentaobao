package ticket

// TicketItemResult 
type TicketItemResult struct {
    // 修改时间
    ModifyedTime   string `json:"modifyed_time,omitempty" xml:"modifyed_time,omitempty"`
    // 商品id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 标准景点id
    AliScenicId   int64 `json:"ali_scenic_id,omitempty" xml:"ali_scenic_id,omitempty"`
    // 商户景点id
    OutScenicId   string `json:"out_scenic_id,omitempty" xml:"out_scenic_id,omitempty"`
    // 标准收费项目id
    AliProductId   int64 `json:"ali_product_id,omitempty" xml:"ali_product_id,omitempty"`
    // 商户收费项目id
    OutProductId   string `json:"out_product_id,omitempty" xml:"out_product_id,omitempty"`
    // 预留，扩展字段
    Extend   string `json:"extend,omitempty" xml:"extend,omitempty"`
    // 票种
    TicketType   string `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
    // 商户票种规则id
    OutRuleIds   []string `json:"out_rule_ids,omitempty" xml:"out_rule_ids>string,omitempty"`
}
