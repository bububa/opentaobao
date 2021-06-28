package promotion

// PagingDto 
type PagingDto struct {

    // 总数
    
    TotalSize   int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
    

    // 总页数
    
    TotalPages   int64 `json:"total_pages,omitempty" xml:"total_pages,omitempty"`
    

    // 页大小
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 第几页
    
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    

    // 抽奖活动列表
    
    LotteryActivityList   []LotteryActivityExtendDto `json:"lottery_activity_list,omitempty" xml:"lottery_activity_list,omitempty"`
    

}
