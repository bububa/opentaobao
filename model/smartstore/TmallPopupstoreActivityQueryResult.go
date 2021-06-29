package smartstore

// TmallPopupstoreActivityQueryResult 
type TmallPopupstoreActivityQueryResult struct {
    // 活动id
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    // 活动结束时间
    ActivityEndTime   string `json:"activity_end_time,omitempty" xml:"activity_end_time,omitempty"`
    // 0：正常状态，-1：删除状态，-2：活动取消
    ActivityStatus   int64 `json:"activity_status,omitempty" xml:"activity_status,omitempty"`
    // 活动名称
    ActivityName   string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
    // 活动开始时间
    ActivityStartTime   string `json:"activity_start_time,omitempty" xml:"activity_start_time,omitempty"`
}
