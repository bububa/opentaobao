package hotel

// AlitripHotelSearchCitysuggestGetResult 
/* model for simplify = false
type AlitripHotelSearchCitysuggestGetResult struct {

    // bizExtMap
    
    BizExtMap  *struct {
        Bizextmap  *Bizextmap `json:"bizextmap,omitempty"`
    } `json:"biz_ext_map,omitempty"`
    

    // headers
    
    Headers  *struct {
        Headers  *Headers `json:"headers,omitempty"`
    } `json:"headers,omitempty"`
    

    // httpStatusCode
    
    HttpStatusCode   int64 `json:"http_status_code,omitempty"`
    

    // mappingCode
    
    MappingCode   string `json:"mapping_code,omitempty"`
    

    // 具体的返回数据
    
    Model  *struct {
        TopCitySuggestListVo  *TopCitySuggestListVo `json:"top_city_suggest_list_vo,omitempty"`
    } `json:"model,omitempty"`
    

    // 错误码
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // 错误消息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlitripHotelSearchCitysuggestGetResult 
type AlitripHotelSearchCitysuggestGetResult struct {

    // bizExtMap
    BizExtMap   *Bizextmap `json:"biz_ext_map,omitempty"`

    // headers
    Headers   *Headers `json:"headers,omitempty"`

    // httpStatusCode
    HttpStatusCode   int64 `json:"http_status_code,omitempty"`

    // mappingCode
    MappingCode   string `json:"mapping_code,omitempty"`

    // 具体的返回数据
    Model   *TopCitySuggestListVo `json:"model,omitempty"`

    // 错误码
    MsgCode   string `json:"msg_code,omitempty"`

    // 错误消息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
