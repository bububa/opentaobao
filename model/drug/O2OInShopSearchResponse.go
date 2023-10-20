package drug

import (
	"sync"
)

// O2OInShopSearchResponse 结构体
type O2OInShopSearchResponse struct {
	// itemList
	ItemList []Itemlist `json:"item_list,omitempty" xml:"item_list>itemlist,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolO2OInShopSearchResponse = sync.Pool{
	New: func() any {
		return new(O2OInShopSearchResponse)
	},
}

// GetO2OInShopSearchResponse() 从对象池中获取O2OInShopSearchResponse
func GetO2OInShopSearchResponse() *O2OInShopSearchResponse {
	return poolO2OInShopSearchResponse.Get().(*O2OInShopSearchResponse)
}

// ReleaseO2OInShopSearchResponse 释放O2OInShopSearchResponse
func ReleaseO2OInShopSearchResponse(v *O2OInShopSearchResponse) {
	v.ItemList = v.ItemList[:0]
	v.TotalCount = 0
	poolO2OInShopSearchResponse.Put(v)
}
