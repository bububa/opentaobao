package security

// JaqImageDetectResultCollection 
/* model for simplify = false
type JaqImageDetectResultCollection struct {

    // 响应消息结构体
    
    JaqImageDetectResultList  struct {
        JaqImageDetectResult  []JaqImageDetectResult `json:"jaq_image_detect_result,omitempty"`
    } `json:"jaq_image_detect_result_list,omitempty"`
    

}
*/

// JaqImageDetectResultCollection 
type JaqImageDetectResultCollection struct {

    // 响应消息结构体
    JaqImageDetectResultList   []JaqImageDetectResult `json:"jaq_image_detect_result_list,omitempty"`

}
