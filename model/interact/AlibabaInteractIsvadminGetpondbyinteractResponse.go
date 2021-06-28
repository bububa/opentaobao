package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据互动实例ID查询奖池信息 APIResponse
alibaba.interact.isvadmin.getpondbyinteract

根据互动实例ID查询奖池信息
*/
type AlibabaInteractIsvadminGetpondbyinteractAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractIsvadminGetpondbyinteractResponse `json:"alibaba_interact_isvadmin_getpondbyinteract_response,omitempty"` 
    AlibabaInteractIsvadminGetpondbyinteractResponse
}

/* model for simplify = false
type AlibabaInteractIsvadminGetpondbyinteractResponse struct {

    // 奖池信息
    
    Data  *struct {
        PrizePondDTO  *PrizePondDTO `json:"prize_pond_dto,omitempty"`
    } `json:"data,omitempty"`
    

    // 是否调用成功
    
    Succ   bool `json:"succ,omitempty"`
    

    // 调用错误原因
    
    Error   string `json:"error,omitempty"`
    

    // 错误描述
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

type AlibabaInteractIsvadminGetpondbyinteractResponse struct {

    // 奖池信息
    Data   *PrizePondDTO `json:"data,omitempty"`

    // 是否调用成功
    Succ   bool `json:"succ,omitempty"`

    // 调用错误原因
    Error   string `json:"error,omitempty"`

    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty"`

}
