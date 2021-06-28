package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
主动预约条件更新 APIResponse
tmall.servicecenter.reservecond.update

1、设置主动预约开通条件
*/
type TmallServicecenterReservecondUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterReservecondUpdateResponse `json:"tmall_servicecenter_reservecond_update_response,omitempty"` 
    TmallServicecenterReservecondUpdateResponse
}

/* model for simplify = false
type TmallServicecenterReservecondUpdateResponse struct {

    // 是否成功
    
    MsgSuccess   bool `json:"msg_success,omitempty"`
    

    // 返回信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 返回code
    
    MsgCode   string `json:"msg_code,omitempty"`
    

}
*/

type TmallServicecenterReservecondUpdateResponse struct {

    // 是否成功
    MsgSuccess   bool `json:"msg_success,omitempty"`

    // 返回信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 返回code
    MsgCode   string `json:"msg_code,omitempty"`

}
