package feedflow

// LaunchTimeDto 
type LaunchTimeDto struct {

    // 是否永远生效
    
    LaunchForever   bool `json:"launch_forever,omitempty" xml:"launch_forever,omitempty"`
    

    // 开始时间
    
    BeginTime   string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
    

    // 结束时间
    
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    

}
