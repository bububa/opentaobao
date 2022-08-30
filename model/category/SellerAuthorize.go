package category

// SellerAuthorize 结构体
type SellerAuthorize struct {
	// 品牌列表
	Brands []Brand `json:"brands,omitempty" xml:"brands>brand,omitempty"`
	// 类目列表
	ItemCats []ItemCat `json:"item_cats,omitempty" xml:"item_cats>item_cat,omitempty"`
	// 商品类目结构
	XinpinItemCats []ItemCat `json:"xinpin_item_cats,omitempty" xml:"xinpin_item_cats>item_cat,omitempty"`
}
