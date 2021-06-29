package usergrowth2

// OfflineMapiQuery 
type OfflineMapiQuery struct {

    // 总条数
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

    // 结束日期（long类型）
    
    EndDate   int64 `json:"end_date,omitempty" xml:"end_date,omitempty"`
    

    // 页大小
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 页码
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    

    // 任务ID
    
    TaskId   string `json:"task_id,omitempty" xml:"task_id,omitempty"`
    

    // 渠道id
    
    ChannelId   string `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
    

    // 开始日期（long类型）
    
    StartDate   int64 `json:"start_date,omitempty" xml:"start_date,omitempty"`
    

    // 推广码code
    
    PromoterCode   string `json:"promoter_code,omitempty" xml:"promoter_code,omitempty"`
    

}
