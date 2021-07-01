package koubeimall

// SuperDiscountDto 
type SuperDiscountDto struct {
    // 商品信息list
    ItemList   []ItemDto `json:"item_list,omitempty" xml:"item_list>item_dto,omitempty"`
}
