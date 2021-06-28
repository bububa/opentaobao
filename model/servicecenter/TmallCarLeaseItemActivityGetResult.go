package servicecenter

// TmallCarLeaseItemActivityGetResult 
type TmallCarLeaseItemActivityGetResult struct {

    // 当前时间
    
    GmtCurrentTime   int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
    

    // 动态返回扩展参数：<br/>extInfo:扩展参数字符串
    
    Piggyback   string `json:"piggyback,omitempty" xml:"piggyback,omitempty"`
    

    // 返回码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    

    // 返回信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    

    // 活动返回列表对象<br/> itemId：商品ID <br/> preEndTime：预热结束时间<br/> preStartTime：预热开始时间<br/> <br/> extendInfo：扩展信息<br/> refActivityId:外部活动ID<br/> <br/> timeRangeList：活动时间范围对象<br/> startTime：活动开始时间<br/> endTime：活动结束时间<br/> amount：名额<br/> extendInfo：扩展信息<br/>
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 耗费时间
    
    CostTime   int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
    

}
