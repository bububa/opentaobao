package promotion

// LotteryActivityQueryDto 
type LotteryActivityQueryDto struct {

    // 是否需要奖品信息
    IsNeedAward   bool `json:"is_need_award,omitempty"`

    // 页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 第几页
    CurrentPage   int64 `json:"current_page,omitempty"`

    // 抽奖活动id
    ActivityIds   []Number `json:"activity_ids,omitempty"`

}
