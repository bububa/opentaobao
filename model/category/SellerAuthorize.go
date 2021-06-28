package category

// SellerAuthorize 
/* model for simplify = false
type SellerAuthorize struct {

    // 类目列表
    
    ItemCats  struct {
        ItemCat  []ItemCat `json:"item_cat,omitempty"`
    } `json:"item_cats,omitempty"`
    

    // 品牌列表
    
    Brands  struct {
        Brand  []Brand `json:"brand,omitempty"`
    } `json:"brands,omitempty"`
    

    // 被授权的新品类目列表
    
    XinpinItemCats  struct {
        ItemCat  []ItemCat `json:"item_cat,omitempty"`
    } `json:"xinpin_item_cats,omitempty"`
    

}
*/

// SellerAuthorize 
type SellerAuthorize struct {

    // 类目列表
    ItemCats   []ItemCat `json:"item_cats,omitempty"`

    // 品牌列表
    Brands   []Brand `json:"brands,omitempty"`

    // 被授权的新品类目列表
    XinpinItemCats   []ItemCat `json:"xinpin_item_cats,omitempty"`

}
