package koubeimall

// SuperDiscountDto 结构体
type SuperDiscountDto struct {
	// 商品信息list
	ItemList []ItemDto `json:"item_list,omitempty" xml:"item_list>item_dto,omitempty"`
}
