package drug

// Cat 结构体
type Cat struct {
	// catId
	CatId string `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// catName
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// itemCount
	ItemCount int64 `json:"item_count,omitempty" xml:"item_count,omitempty"`
}
