package koubeimall

// PageResult 
type PageResult struct {
    // 已授权商圈列表data模型
    MallInfoList   []MallDTO `json:"mall_info_list,omitempty" xml:"mall_info_list>mall_dto,omitempty"`
    // 是否有更多信息
    HasMore   bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
    // 下一页起始值
    NextStart   int64 `json:"next_start,omitempty" xml:"next_start,omitempty"`
    // 分页长度
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 门店评论信息dto
    CommentInfo   *CommentInfoDTO `json:"comment_info,omitempty" xml:"comment_info,omitempty"`
    // 门店信息模型
    StoreList   []StoreDTO `json:"store_list,omitempty" xml:"store_list>store_dto,omitempty"`
}
