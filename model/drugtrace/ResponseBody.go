package drugtrace

// ResponseBody 
type ResponseBody struct {

    // 流向对象
    
    FlowList   []FlowEntity `json:"flow_list,omitempty" xml:"flow_list,omitempty"`
    

    // 商品信息、生产信息
    
    ProductInfoList   *ProductInfoList `json:"product_info_list,omitempty" xml:"product_info_list,omitempty"`
    

    // 首次查询时间
    
    FirstQueryTime   string `json:"first_query_time,omitempty" xml:"first_query_time,omitempty"`
    

    // 最后操作时间
    
    LastBizDate   string `json:"last_biz_date,omitempty" xml:"last_biz_date,omitempty"`
    

    // 药品状态
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 查询次数
    
    QueryCount   string `json:"query_count,omitempty" xml:"query_count,omitempty"`
    

}
