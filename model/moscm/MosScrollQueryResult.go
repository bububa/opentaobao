package moscm

// MosScrollQueryResult 结构体
type MosScrollQueryResult struct {
	// 返回数据
	Data []InventoryDetailDto `json:"data,omitempty" xml:"data>inventory_detail_dto,omitempty"`
	// 滚动查询id号
	ScrollId string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// 总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
