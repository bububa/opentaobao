package promotion

// LotteryActivityQueryDto 
type LotteryActivityQueryDto struct {
    // 是否需要奖品信息
    IsNeedAward   bool `json:"is_need_award,omitempty" xml:"is_need_award,omitempty"`
    // 页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 第几页
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    // 抽奖活动id
    ActivityIds   []int64 `json:"activity_ids,omitempty" xml:"activity_ids>int64,omitempty"`
}
