package ascpchannel

// SalecategoryQueryResponse 
type SalecategoryQueryResponse struct {
    // 货品ID
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 货品品类ID
    ItemSalecategory   int64 `json:"item_salecategory,omitempty" xml:"item_salecategory,omitempty"`
    // 货品品类名称
    ItemSalecategoryName   string `json:"item_salecategory_name,omitempty" xml:"item_salecategory_name,omitempty"`
}
