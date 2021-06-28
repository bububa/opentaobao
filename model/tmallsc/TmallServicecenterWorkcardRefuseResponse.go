package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
买家拒收 APIResponse
tmall.servicecenter.workcard.refuse

买家拒收通知接口
*/
type TmallServicecenterWorkcardRefuseAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterWorkcardRefuseResponse `json:"tmall_servicecenter_workcard_refuse_response,omitempty"` 
    TmallServicecenterWorkcardRefuseResponse
}

/* model for simplify = false
type TmallServicecenterWorkcardRefuseResponse struct {

    // 返回信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 返回code
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // 是否成功
    
    MsgSuccess   bool `json:"msg_success,omitempty"`
    

}
*/

type TmallServicecenterWorkcardRefuseResponse struct {

    // 返回信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 返回code
    MsgCode   string `json:"msg_code,omitempty"`

    // 是否成功
    MsgSuccess   bool `json:"msg_success,omitempty"`

}
