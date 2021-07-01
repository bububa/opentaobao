package drug

// O2OInShopSearchResponse 结构体
type O2OInShopSearchResponse struct {
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// itemList
	ItemList []Itemlist `json:"item_list,omitempty" xml:"item_list>itemlist,omitempty"`
}
