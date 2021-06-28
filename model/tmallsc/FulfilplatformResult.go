package tmallsc

// FulfilplatformResult 
type FulfilplatformResult struct {

    // 错误描述
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    

    // 错误码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    

    // true：成功；false：失败
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 结算明细list
    
    ResultDatas   []Resultdata `json:"result_datas,omitempty" xml:"result_datas,omitempty"`
    

    // 每页数据信息
    
    ResultData   *PagedResult `json:"result_data,omitempty" xml:"result_data,omitempty"`
    

}
