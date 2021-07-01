package idle

// ItemPageQuery 结构体
type ItemPageQuery struct {
	// 类目集
	CategoryIds []int64 `json:"category_ids,omitempty" xml:"category_ids>int64,omitempty"`
	// 页号
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 商品状态 0:在线 1售出
	Status []int64 `json:"status,omitempty" xml:"status>int64,omitempty"`
}
