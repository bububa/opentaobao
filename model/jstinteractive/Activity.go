package jstinteractive

// Activity 
type Activity struct {
    // 活动id
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    // 活动名称
    ActivityName   string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
    // 活动开始时间
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 活动结束时间
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 活动状态，0=无效，1=有效
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
}
