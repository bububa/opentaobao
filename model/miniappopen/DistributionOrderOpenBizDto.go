package miniappopen

// DistributionOrderOpenBizDto 
type DistributionOrderOpenBizDto struct {

    // 唯一标识的id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 名字
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 开始时间，订购期有效的情况无该返回值
    
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    

    // 结束时间，订购期有效的投放计划无该返回值
    
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    

    // 状态，0:未开始， 1：进行中，2/3:已结束，其他为平台状态
    
    Status   *model.File `json:"status,omitempty" xml:"status,omitempty"`
    

    // 投放的小程序url
    
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    

    // 小程序id
    
    AppId   int64 `json:"app_id,omitempty" xml:"app_id,omitempty"`
    

    // 时效类型.0: 商家自定义时效，1：订购期内有效
    
    TimeType   *model.File `json:"time_type,omitempty" xml:"time_type,omitempty"`
    

    // 关联的小部件id
    
    WidgetId   int64 `json:"widget_id,omitempty" xml:"widget_id,omitempty"`
    

    // 关联的小部件版本号
    
    WidgetVersion   string `json:"widget_version,omitempty" xml:"widget_version,omitempty"`
    

}
