package promotion

// PagingDto 
/* model for simplify = false
type PagingDto struct {

    // 总数
    
    TotalSize   int64 `json:"total_size,omitempty"`
    

    // 总页数
    
    TotalPages   int64 `json:"total_pages,omitempty"`
    

    // 页大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 第几页
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // 抽奖活动列表
    
    LotteryActivityList  struct {
        LotteryActivityExtendDto  []LotteryActivityExtendDto `json:"lottery_activity_extend_dto,omitempty"`
    } `json:"lottery_activity_list,omitempty"`
    

}
*/

// PagingDto 
type PagingDto struct {

    // 总数
    TotalSize   int64 `json:"total_size,omitempty"`

    // 总页数
    TotalPages   int64 `json:"total_pages,omitempty"`

    // 页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 第几页
    PageNo   int64 `json:"page_no,omitempty"`

    // 抽奖活动列表
    LotteryActivityList   []LotteryActivityExtendDto `json:"lottery_activity_list,omitempty"`

}
