package alsc

// LevelPointOpenInfo 
type LevelPointOpenInfo struct {

    // 是否已删除
    Deleted   bool `json:"deleted,omitempty"`

    // 等级ID
    LevelId   string `json:"level_id,omitempty"`

    // 等级名称
    LevelName   string `json:"level_name,omitempty"`

    // 积分奖励倍数
    Times   string `json:"times,omitempty"`

    // 是否参与积分奖励
    UseSwitch   bool `json:"use_switch,omitempty"`

}
