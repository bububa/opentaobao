package drug

// O2OInShopSearchResponse 
/* model for simplify = false
type O2OInShopSearchResponse struct {

    // totalCount
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // itemList
    
    ItemList  struct {
        Itemlist  []Itemlist `json:"itemlist,omitempty"`
    } `json:"item_list,omitempty"`
    

}
*/

// O2OInShopSearchResponse 
type O2OInShopSearchResponse struct {

    // totalCount
    TotalCount   int64 `json:"total_count,omitempty"`

    // itemList
    ItemList   []Itemlist `json:"item_list,omitempty"`

}
