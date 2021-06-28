package hotel

// AlitripHotelDetailStaticinfoGetResult 
/* model for simplify = false
type AlitripHotelDetailStaticinfoGetResult struct {

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
    

    // 详情页静态信息
    
    Model  *struct {
        HotelTopStaticDetailsVo  *HotelTopStaticDetailsVo `json:"hotel_top_static_details_vo,omitempty"`
    } `json:"model,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlitripHotelDetailStaticinfoGetResult 
type AlitripHotelDetailStaticinfoGetResult struct {

    // bizExtMap
    BizExtMap   *Bizextmap `json:"biz_ext_map,omitempty"`

    // headers
    Headers   *Headers `json:"headers,omitempty"`

    // httpStatusCode
    HttpStatusCode   int64 `json:"http_status_code,omitempty"`

    // mappingCode
    MappingCode   string `json:"mapping_code,omitempty"`

    // 详情页静态信息
    Model   *HotelTopStaticDetailsVo `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
