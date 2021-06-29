package usergrowth2

// OfflineConvertionSyncInfoQuery 
type OfflineConvertionSyncInfoQuery struct {
    // 任务类型
    TaskType   string `json:"task_type,omitempty" xml:"task_type,omitempty"`
    // 渠道id
    ChannelId   int64 `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
    // 日期
    Ds   int64 `json:"ds,omitempty" xml:"ds,omitempty"`
}
