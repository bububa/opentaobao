package taotv

// TaobaoTaotvVideoPlaylistAllResult 
/* model for simplify = false
type TaobaoTaotvVideoPlaylistAllResult struct {

    // model
    
    ModelList  struct {
        TaobaoTaotvVideoPlaylistAllModel  []TaobaoTaotvVideoPlaylistAllModel `json:"taobao_taotv_video_playlist_all_model,omitempty"`
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

// TaobaoTaotvVideoPlaylistAllResult 
type TaobaoTaotvVideoPlaylistAllResult struct {

    // model
    ModelList   []TaobaoTaotvVideoPlaylistAllModel `json:"model_list,omitempty"`

    // httpStatusCode
    HttpStatusCode   int64 `json:"http_status_code,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
