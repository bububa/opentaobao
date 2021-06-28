package taotv

// TaobaoTaotvVideoPlaylistOttnavGetResult 
/* model for simplify = false
type TaobaoTaotvVideoPlaylistOttnavGetResult struct {

    // model
    
    Model  *struct {
        PlayListNavRbo  *PlayListNavRbo `json:"play_list_nav_rbo,omitempty"`
    } `json:"model,omitempty"`
    

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

// TaobaoTaotvVideoPlaylistOttnavGetResult 
type TaobaoTaotvVideoPlaylistOttnavGetResult struct {

    // model
    Model   *PlayListNavRbo `json:"model,omitempty"`

    // httpStatusCode
    HttpStatusCode   int64 `json:"http_status_code,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
