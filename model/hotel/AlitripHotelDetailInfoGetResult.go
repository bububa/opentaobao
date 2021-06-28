package hotel

// AlitripHotelDetailInfoGetResult 
/* model for simplify = false
type AlitripHotelDetailInfoGetResult struct {

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
    

    // 返回结果
    
    Model  *struct {
        HotelTopDetailsVo  *HotelTopDetailsVo `json:"hotel_top_details_vo,omitempty"`
    } `json:"model,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlitripHotelDetailInfoGetResult 
type AlitripHotelDetailInfoGetResult struct {

    // bizExtMap
    BizExtMap   *Bizextmap `json:"biz_ext_map,omitempty"`

    // headers
    Headers   *Headers `json:"headers,omitempty"`

    // httpStatusCode
    HttpStatusCode   int64 `json:"http_status_code,omitempty"`

    // mappingCode
    MappingCode   string `json:"mapping_code,omitempty"`

    // 返回结果
    Model   *HotelTopDetailsVo `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
