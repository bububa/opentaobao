package scbp

// AccountEffectDto 
type AccountEffectDto struct {

    // 日期(yyyy-MM-dd)
    StatDate   string `json:"stat_date,omitempty"`

    // 曝光
    Impr   string `json:"impr,omitempty"`

    // 点击
    Click   string `json:"click,omitempty"`

    // 消耗
    Cost   string `json:"cost,omitempty"`

    // 推广时长
    OnlineMin   string `json:"online_min,omitempty"`

    // 曝光
    ImpressionCnt   string `json:"impression_cnt,omitempty"`

    // 点击率
    Ctr   string `json:"ctr,omitempty"`

    // 平均点击花费
    ClickCostAvg   string `json:"click_cost_avg,omitempty"`

    // 单位小时，保留一位小数，例如13.5表示13.5小时
    OnlineTime   string `json:"online_time,omitempty"`

    // 点击量
    ClickCnt   string `json:"click_cnt,omitempty"`

}
