package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取天猫互动奖池列表 APIResponse
alibaba.interact.isvadmin.allponds

获取天猫互动奖池列表
*/
type AlibabaInteractIsvadminAllpondsAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractIsvadminAllpondsResponse `json:"alibaba_interact_isvadmin_allponds_response,omitempty"` 
    AlibabaInteractIsvadminAllpondsResponse
}

/* model for simplify = false
type AlibabaInteractIsvadminAllpondsResponse struct {

    // 是否成功
    
    Succ   bool `json:"succ,omitempty"`
    

    // 错误码
    
    InteractErrorCode   string `json:"interact_error_code,omitempty"`
    

    // 错误描述
    
    InteractErrorMsg   string `json:"interact_error_msg,omitempty"`
    

    // 奖池列表
    
    Allponds  struct {
        PrizePondDTO  []PrizePondDTO `json:"prize_pond_dto,omitempty"`
    } `json:"allponds,omitempty"`
    

}
*/

type AlibabaInteractIsvadminAllpondsResponse struct {

    // 是否成功
    Succ   bool `json:"succ,omitempty"`

    // 错误码
    InteractErrorCode   string `json:"interact_error_code,omitempty"`

    // 错误描述
    InteractErrorMsg   string `json:"interact_error_msg,omitempty"`

    // 奖池列表
    Allponds   []PrizePondDTO `json:"allponds,omitempty"`

}
