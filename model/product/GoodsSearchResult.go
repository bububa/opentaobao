package product

// GoodsSearchResult 
/* model for simplify = false
type GoodsSearchResult struct {

    // 搜素结果所在的类目信息
    
    Category  *struct {
        WholesaleCategory  *WholesaleCategory `json:"wholesale_category,omitempty"`
    } `json:"category,omitempty"`
    

    // 返回结果数
    
    Total   int64 `json:"total,omitempty"`
    

    // 搜索产品列表
    
    Items  struct {
        GoodsSummary  []GoodsSummary `json:"goods_summary,omitempty"`
    } `json:"items,omitempty"`
    

}
*/

// GoodsSearchResult 
type GoodsSearchResult struct {

    // 搜素结果所在的类目信息
    Category   *WholesaleCategory `json:"category,omitempty"`

    // 返回结果数
    Total   int64 `json:"total,omitempty"`

    // 搜索产品列表
    Items   []GoodsSummary `json:"items,omitempty"`

}
