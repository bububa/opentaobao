package promotion

// PromotionTagQuery 
type PromotionTagQuery struct {
    // 标签结果列表
    TagList   []PromotionTag `json:"tag_list,omitempty" xml:"tag_list>promotion_tag,omitempty"`
    // 总记录数
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
