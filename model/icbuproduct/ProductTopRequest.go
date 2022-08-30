package icbuproduct

// ProductTopRequest 结构体
type ProductTopRequest struct {
	// 返回结果语种
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
}
