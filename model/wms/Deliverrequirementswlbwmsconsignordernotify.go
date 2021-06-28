package wms

// Deliverrequirementswlbwmsconsignordernotify 
/* model for simplify = false
type Deliverrequirementswlbwmsconsignordernotify struct {

    // 投递时延要求:  1-工作日 2-节假日 101,当日达102次晨达103次日达 111 活动标  104 预约达
    
    ScheduleType   int64 `json:"schedule_type,omitempty"`
    

    // 送达日期
    
    ScheduleDay   string `json:"schedule_day,omitempty"`
    

    // 送达开始时间
    
    ScheduleStart   string `json:"schedule_start,omitempty"`
    

    // 送达结束时间
    
    ScheduleEnd   string `json:"schedule_end,omitempty"`
    

    // 配送类型： PTPS-常温配送 LLPS-冷链配送
    
    DeliveryType   string `json:"delivery_type,omitempty"`
    

}
*/

// Deliverrequirementswlbwmsconsignordernotify 
type Deliverrequirementswlbwmsconsignordernotify struct {

    // 投递时延要求:  1-工作日 2-节假日 101,当日达102次晨达103次日达 111 活动标  104 预约达
    ScheduleType   int64 `json:"schedule_type,omitempty"`

    // 送达日期
    ScheduleDay   string `json:"schedule_day,omitempty"`

    // 送达开始时间
    ScheduleStart   string `json:"schedule_start,omitempty"`

    // 送达结束时间
    ScheduleEnd   string `json:"schedule_end,omitempty"`

    // 配送类型： PTPS-常温配送 LLPS-冷链配送
    DeliveryType   string `json:"delivery_type,omitempty"`

}
