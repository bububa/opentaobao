package tmallsc

// FulfilplatformResult 
/* model for simplify = false
type FulfilplatformResult struct {

    // 错误描述
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 错误码
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // true：成功；false：失败
    
    Success   bool `json:"success,omitempty"`
    

    // 结算明细list
    
    ResultDatas  struct {
        Resultdata  []Resultdata `json:"resultdata,omitempty"`
    } `json:"result_datas,omitempty"`
    

    // 每页数据信息
    
    ResultData  *struct {
        PagedResult  *PagedResult `json:"paged_result,omitempty"`
    } `json:"result_data,omitempty"`
    

}
*/

// FulfilplatformResult 
type FulfilplatformResult struct {

    // 错误描述
    MsgInfo   string `json:"msg_info,omitempty"`

    // 错误码
    MsgCode   string `json:"msg_code,omitempty"`

    // true：成功；false：失败
    Success   bool `json:"success,omitempty"`

    // 结算明细list
    ResultDatas   []Resultdata `json:"result_datas,omitempty"`

    // 每页数据信息
    ResultData   *PagedResult `json:"result_data,omitempty"`

}
