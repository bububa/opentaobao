package promotion

// PromotionTagQuery 
/* model for simplify = false
type PromotionTagQuery struct {

    // 标签结果列表
    
    TagList  struct {
        PromotionTag  []PromotionTag `json:"promotion_tag,omitempty"`
    } `json:"tag_list,omitempty"`
    

    // 总记录数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

// PromotionTagQuery 
type PromotionTagQuery struct {

    // 标签结果列表
    TagList   []PromotionTag `json:"tag_list,omitempty"`

    // 总记录数
    TotalResults   int64 `json:"total_results,omitempty"`

}
