package alsc

// AlibabaAlscSaasCodecCodeAttrsQueryResult 
/* model for simplify = false
type AlibabaAlscSaasCodecCodeAttrsQueryResult struct {

    // 返回素材id
    
    Data  *struct {
        CodeBizAttributeDto  *CodeBizAttributeDto `json:"code_biz_attribute_dto,omitempty"`
    } `json:"data,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

}
*/

// AlibabaAlscSaasCodecCodeAttrsQueryResult 
type AlibabaAlscSaasCodecCodeAttrsQueryResult struct {

    // 返回素材id
    Data   *CodeBizAttributeDto `json:"data,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

}
