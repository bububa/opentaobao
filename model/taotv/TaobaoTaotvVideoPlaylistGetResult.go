package taotv

// TaobaoTaotvVideoPlaylistGetResult 
/* model for simplify = false
type TaobaoTaotvVideoPlaylistGetResult struct {

    // 数据列表
    
    ModelList  struct {
        TaobaoTaotvVideoPlaylistGetModel  []TaobaoTaotvVideoPlaylistGetModel `json:"taobao_taotv_video_playlist_get_model,omitempty"`
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

// TaobaoTaotvVideoPlaylistGetResult 
type TaobaoTaotvVideoPlaylistGetResult struct {

    // 数据列表
    ModelList   []TaobaoTaotvVideoPlaylistGetModel `json:"model_list,omitempty"`

    // httpStatusCode
    HttpStatusCode   int64 `json:"http_status_code,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
