package omniorder

// StoreCollectConfig 
type StoreCollectConfig struct {

    // 当activity为true时返回，活动结束时间
    
    ActivityEndTime   string `json:"activity_end_time,omitempty" xml:"activity_end_time,omitempty"`
    

    // 接单时间段，格式为 "09:00-12:00", "" 表示一直开启
    
    WorkingTime   string `json:"working_time,omitempty" xml:"working_time,omitempty"`
    

    // 每日接单阈值
    
    CollectThreshold   int64 `json:"collect_threshold,omitempty" xml:"collect_threshold,omitempty"`
    

    // 当activity为true时返回,活动开始时间
    
    ActivityStartTime   string `json:"activity_start_time,omitempty" xml:"activity_start_time,omitempty"`
    

    // 是否是活动期
    
    Activity   bool `json:"activity,omitempty" xml:"activity,omitempty"`
    

}
