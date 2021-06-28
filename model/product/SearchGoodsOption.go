package product

// SearchGoodsOption 
/* model for simplify = false
type SearchGoodsOption struct {

    // price_to_cent 最大价格
    
    PriceToCent   int64 `json:"price_to_cent,omitempty"`
    

    // price_from_cent 最小价格
    
    PriceFromCent   int64 `json:"price_from_cent,omitempty"`
    

    // sort_by 排序方式
    
    SortBy   string `json:"sort_by,omitempty"`
    

    // page_no 当前页面
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // keyword 搜索关键词
    
    Keyword   string `json:"keyword,omitempty"`
    

    // category_id 搜索类目
    
    CategoryId   string `json:"category_id,omitempty"`
    

    // page_size 每页大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // min_order_from 最小起订量区间
    
    MinOrderFrom   int64 `json:"min_order_from,omitempty"`
    

    // min_order_to 最小起订量区间
    
    MinOrderTo   int64 `json:"min_order_to,omitempty"`
    

    // 产品筛选tags
    
    ProductRefineTags  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"product_refine_tags,omitempty"`
    

}
*/

// SearchGoodsOption 
type SearchGoodsOption struct {

    // price_to_cent 最大价格
    PriceToCent   int64 `json:"price_to_cent,omitempty"`

    // price_from_cent 最小价格
    PriceFromCent   int64 `json:"price_from_cent,omitempty"`

    // sort_by 排序方式
    SortBy   string `json:"sort_by,omitempty"`

    // page_no 当前页面
    PageNo   int64 `json:"page_no,omitempty"`

    // keyword 搜索关键词
    Keyword   string `json:"keyword,omitempty"`

    // category_id 搜索类目
    CategoryId   string `json:"category_id,omitempty"`

    // page_size 每页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // min_order_from 最小起订量区间
    MinOrderFrom   int64 `json:"min_order_from,omitempty"`

    // min_order_to 最小起订量区间
    MinOrderTo   int64 `json:"min_order_to,omitempty"`

    // 产品筛选tags
    ProductRefineTags   []int64 `json:"product_refine_tags,omitempty"`

}
