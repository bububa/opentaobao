package cloudgame

// ScoreReportDto 
type ScoreReportDto struct {
    // 游戏结果
    Score   string `json:"score,omitempty" xml:"score,omitempty"`
    // SDK传入的gameSessionId
    GameSessionId   string `json:"game_session_id,omitempty" xml:"game_session_id,omitempty"`
    // 1 - 保存最佳战绩     2 - 强制更新战绩
    OverwriteMethod   int64 `json:"overwrite_method,omitempty" xml:"overwrite_method,omitempty"`
    // SDK传入的云游戏混合用户ID
    MixUserId   string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
    // 全局唯一uuid, 用于幂等校验
    Uuid   string `json:"uuid,omitempty" xml:"uuid,omitempty"`
    // 扩展信息JSON串
    ExtInfo   string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
    // 上报时间戳
    ReportTimestamp   int64 `json:"report_timestamp,omitempty" xml:"report_timestamp,omitempty"`
}
