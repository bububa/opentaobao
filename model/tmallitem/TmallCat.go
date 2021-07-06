package tmallitem

// TmallCat 结构体
type TmallCat struct {
	// 搜索前台类目名字
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// 搜索前台类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
}
