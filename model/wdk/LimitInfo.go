package wdk

// LimitInfo 
type LimitInfo struct {
    // 用户总限购
    UserLimitCnt   int64 `json:"user_limit_cnt,omitempty" xml:"user_limit_cnt,omitempty"`
    // 用户每日限购
    UserDailyLimitCnt   int64 `json:"user_daily_limit_cnt,omitempty" xml:"user_daily_limit_cnt,omitempty"`
    // 每日总限购
    DailyLimitCnt   int64 `json:"daily_limit_cnt,omitempty" xml:"daily_limit_cnt,omitempty"`
    // 总限购
    TotalLimitCnt   int64 `json:"total_limit_cnt,omitempty" xml:"total_limit_cnt,omitempty"`
    // 活动每日总限购（非标小数）
    DailyLimitCntDouble   string `json:"daily_limit_cnt_double,omitempty" xml:"daily_limit_cnt_double,omitempty"`
    // 每人每日限购（非标小数）
    UserDailyLimitCntDouble   string `json:"user_daily_limit_cnt_double,omitempty" xml:"user_daily_limit_cnt_double,omitempty"`
    // 每人活动期间总限购（非标小数）
    UserLimitCntDouble   string `json:"user_limit_cnt_double,omitempty" xml:"user_limit_cnt_double,omitempty"`
    // 活动期间总限购（非标小数）
    TotalLimitCntDouble   string `json:"total_limit_cnt_double,omitempty" xml:"total_limit_cnt_double,omitempty"`
}
