package taotv

// TaobaoTaotvCarouselCategoryListResult 
/* model for simplify = false
type TaobaoTaotvCarouselCategoryListResult struct {

    // 数据列表
    
    ModelList  struct {
        TaobaoTaotvCarouselCategoryListModel  []TaobaoTaotvCarouselCategoryListModel `json:"taobao_taotv_carousel_category_list_model,omitempty"`
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

// TaobaoTaotvCarouselCategoryListResult 
type TaobaoTaotvCarouselCategoryListResult struct {

    // 数据列表
    ModelList   []TaobaoTaotvCarouselCategoryListModel `json:"model_list,omitempty"`

    // httpStatusCode
    HttpStatusCode   int64 `json:"http_status_code,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
