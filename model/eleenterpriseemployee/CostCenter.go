package eleenterpriseemployee

// CostCenter 
type CostCenter struct {
    // 删除成本中心列表
    DeleteItemIds   []string `json:"delete_item_ids,omitempty" xml:"delete_item_ids>string,omitempty"`
    // 新增成本中心列表
    AddItemIds   []string `json:"add_item_ids,omitempty" xml:"add_item_ids>string,omitempty"`
}
