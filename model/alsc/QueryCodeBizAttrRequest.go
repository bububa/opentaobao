package alsc

// QueryCodeBizAttrRequest 
/* model for simplify = false
type QueryCodeBizAttrRequest struct {

    // 码使用的业务场景
    
    BizScenePrefix   string `json:"biz_scene_prefix,omitempty"`
    

    // 码值
    
    CodeValue   string `json:"code_value,omitempty"`
    

}
*/

// QueryCodeBizAttrRequest 
type QueryCodeBizAttrRequest struct {

    // 码使用的业务场景
    BizScenePrefix   string `json:"biz_scene_prefix,omitempty"`

    // 码值
    CodeValue   string `json:"code_value,omitempty"`

}
