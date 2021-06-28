package user

// AlibabaAilabsUserSpeechGuideResult 
/* model for simplify = false
type AlibabaAilabsUserSpeechGuideResult struct {

    // 推荐信息model
    
    RetValue  *struct {
        RecommendInfo  *RecommendInfo `json:"recommend_info,omitempty"`
    } `json:"ret_value,omitempty"`
    

    // 0表示成功
    
    RetCode   int64 `json:"ret_code,omitempty"`
    

    // 出错信息
    
    RetMsg   string `json:"ret_msg,omitempty"`
    

}
*/

// AlibabaAilabsUserSpeechGuideResult 
type AlibabaAilabsUserSpeechGuideResult struct {

    // 推荐信息model
    RetValue   *RecommendInfo `json:"ret_value,omitempty"`

    // 0表示成功
    RetCode   int64 `json:"ret_code,omitempty"`

    // 出错信息
    RetMsg   string `json:"ret_msg,omitempty"`

}
