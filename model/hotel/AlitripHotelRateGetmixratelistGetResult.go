package hotel

// AlitripHotelRateGetmixratelistGetResult 
/* model for simplify = false
type AlitripHotelRateGetmixratelistGetResult struct {

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
    

    // model
    
    Model  *struct {
        GetMixRateListResult  *GetMixRateListResult `json:"get_mix_rate_list_result,omitempty"`
    } `json:"model,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlitripHotelRateGetmixratelistGetResult 
type AlitripHotelRateGetmixratelistGetResult struct {

    // bizExtMap
    BizExtMap   *Bizextmap `json:"biz_ext_map,omitempty"`

    // headers
    Headers   *Headers `json:"headers,omitempty"`

    // httpStatusCode
    HttpStatusCode   int64 `json:"http_status_code,omitempty"`

    // mappingCode
    MappingCode   string `json:"mapping_code,omitempty"`

    // model
    Model   *GetMixRateListResult `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
