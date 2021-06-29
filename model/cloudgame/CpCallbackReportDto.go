package cloudgame

// CpCallbackReportDTO 
type CpCallbackReportDTO struct {
    // 云游戏业务类型, 不同业务类型对应bizData格式不同
    BizType   string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    // 启动游戏票据,可以用来换取gameSessionId
    Ticket   string `json:"ticket,omitempty" xml:"ticket,omitempty"`
    // 云游戏加密用户Id
    MixUserId   string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
    // 约定的appKey
    GameAppKey   string `json:"game_app_key,omitempty" xml:"game_app_key,omitempty"`
    // 全局唯一id, 用于接口幂等
    Uuid   string `json:"uuid,omitempty" xml:"uuid,omitempty"`
    // 时间戳, 毫秒级别
    ReportTimestamp   int64 `json:"report_timestamp,omitempty" xml:"report_timestamp,omitempty"`
    // 回传业务数据, 针对biztype=0, 格式包含: 过关数, 剩余生命, 每关使用生命,每关耗时.
    CallbackBizData   string `json:"callback_biz_data,omitempty" xml:"callback_biz_data,omitempty"`
}
