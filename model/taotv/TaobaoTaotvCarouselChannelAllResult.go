package taotv

// TaobaoTaotvCarouselChannelAllResult 
/* model for simplify = false
type TaobaoTaotvCarouselChannelAllResult struct {

    // 频道列表
    
    ModelList  struct {
        TaobaoTaotvCarouselChannelAllModel  []TaobaoTaotvCarouselChannelAllModel `json:"taobao_taotv_carousel_channel_all_model,omitempty"`
    } `json:"model_list,omitempty"`
    

    // httpStatusCode
    
    HttpStatusCode   int64 `json:"http_status_code,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TaobaoTaotvCarouselChannelAllResult 
type TaobaoTaotvCarouselChannelAllResult struct {

    // 频道列表
    ModelList   []TaobaoTaotvCarouselChannelAllModel `json:"model_list,omitempty"`

    // httpStatusCode
    HttpStatusCode   int64 `json:"http_status_code,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
