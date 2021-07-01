package omniorder

// OmniItemCategoryDto 
type OmniItemCategoryDto struct {
    // 类目名称
    CategoryName   string `json:"category_name,omitempty" xml:"category_name,omitempty"`
    // 类目cid
    Cid   int64 `json:"cid,omitempty" xml:"cid,omitempty"`
    // props
    Props   []OmniItemCategoryPropDto `json:"props,omitempty" xml:"props>omni_item_category_prop_dto,omitempty"`
}
