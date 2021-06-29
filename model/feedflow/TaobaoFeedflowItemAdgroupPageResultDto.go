package feedflow

// TaobaoFeedflowItemAdgroupPageResultDto 
type TaobaoFeedflowItemAdgroupPageResultDto struct {

    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 返回数据结果
    
    Results   []AdgroupDTo `json:"results,omitempty" xml:"results,omitempty"`
    

    // 条数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    

    // 结果
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
