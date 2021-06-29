package usergrowth2

// TaobaoUsergrowthOfflineConvertionDetailsGetE 
type TaobaoUsergrowthOfflineConvertionDetailsGetE struct {

    // 新登时间
    
    ActionTime   string `json:"action_time,omitempty" xml:"action_time,omitempty"`
    

    // 活动名称
    
    ActivityName   string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
    

    // 是否是新登用户，1是0非2新登非归因
    
    IsNewUser   string `json:"is_new_user,omitempty" xml:"is_new_user,omitempty"`
    

    // 日期
    
    Ds   int64 `json:"ds,omitempty" xml:"ds,omitempty"`
    

    // 活动ID
    
    ActivityId   string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    

    // 脱敏手机号
    
    DesensitizationMobile   string `json:"desensitization_mobile,omitempty" xml:"desensitization_mobile,omitempty"`
    

    // 任务名称
    
    TaskName   string `json:"task_name,omitempty" xml:"task_name,omitempty"`
    

    // 是否是首购用户，1是0非
    
    IsPurchaseFlag   string `json:"is_purchase_flag,omitempty" xml:"is_purchase_flag,omitempty"`
    

    // 渠道名称
    
    Company   string `json:"company,omitempty" xml:"company,omitempty"`
    

    // 首购时间
    
    FirstPayTime   string `json:"first_pay_time,omitempty" xml:"first_pay_time,omitempty"`
    

    // 求道id
    
    ChannelId   string `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
    

    // 任务id
    
    TaskId   string `json:"task_id,omitempty" xml:"task_id,omitempty"`
    

    // 扫码时间
    
    TargetTime   string `json:"target_time,omitempty" xml:"target_time,omitempty"`
    

    // 推广码code
    
    PromoterCode   string `json:"promoter_code,omitempty" xml:"promoter_code,omitempty"`
    

}
