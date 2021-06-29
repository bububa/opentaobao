package koubeimall

// SuperDiscountDTO 
type SuperDiscountDTO struct {
    // 商品信息list
    ItemList   []ItemDTO `json:"item_list,omitempty" xml:"item_list>item_dto,omitempty"`
}
