package film

// CinemaPageCinemaVo 
type CinemaPageCinemaVo struct {

    // 活动标
    ActivityTags   []ActivityTagVo `json:"activity_tags,omitempty"`

    // 影院状态
    BizStatus   int64 `json:"biz_status,omitempty"`

    // 影院特殊提醒
    SpecialRemind   string `json:"special_remind,omitempty"`

    // 排期数
    ScheduleCount   int64 `json:"schedule_count,omitempty"`

    // 场次提前几分钟停止场次售卖
    ScheduleCloseTime   int64 `json:"schedule_close_time,omitempty"`

    // 用户是否已开卡
    McardOpen   bool `json:"mcard_open,omitempty"`

    // 影城卡标
    SupportMcard   bool `json:"support_mcard,omitempty"`

    // 常去影院
    AlwaysGo   bool `json:"always_go,omitempty"`

    // 最低票价(分)
    MinPrice   int64 `json:"min_price,omitempty"`

    // 距离
    Distance   string `json:"distance,omitempty"`

    // 纬度
    Latitude   string `json:"latitude,omitempty"`

    // 经度
    Longitude   string `json:"longitude,omitempty"`

    // 特色标
    DisplaySupports   []String `json:"display_supports,omitempty"`

    // 地址
    Address   string `json:"address,omitempty"`

    // 影院名称
    CinemaName   string `json:"cinema_name,omitempty"`

    // 影院ID
    CinemaId   int64 `json:"cinema_id,omitempty"`

}
