package promotion

// LotteryActivityQueryDto 
/* model for simplify = false
type LotteryActivityQueryDto struct {

    // 是否需要奖品信息
    
    IsNeedAward   bool `json:"is_need_award,omitempty"`
    

    // 页大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 第几页
    
    CurrentPage   int64 `json:"current_page,omitempty"`
    

    // 抽奖活动id
    
    ActivityIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"activity_ids,omitempty"`
    

}
*/

// LotteryActivityQueryDto 
type LotteryActivityQueryDto struct {

    // 是否需要奖品信息
    IsNeedAward   bool `json:"is_need_award,omitempty"`

    // 页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 第几页
    CurrentPage   int64 `json:"current_page,omitempty"`

    // 抽奖活动id
    ActivityIds   []int64 `json:"activity_ids,omitempty"`

}
