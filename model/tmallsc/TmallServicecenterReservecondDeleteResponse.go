package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除主动预约开通条件 APIResponse
tmall.servicecenter.reservecond.delete

删除主动预约开通条件
*/
type TmallServicecenterReservecondDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterReservecondDeleteResponse `json:"tmall_servicecenter_reservecond_delete_response,omitempty"` 
    TmallServicecenterReservecondDeleteResponse
}

/* model for simplify = false
type TmallServicecenterReservecondDeleteResponse struct {

    // 是否成功
    
    MsgSuccess   bool `json:"msg_success,omitempty"`
    

    // 返回编码
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // 返回信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

}
*/

type TmallServicecenterReservecondDeleteResponse struct {

    // 是否成功
    MsgSuccess   bool `json:"msg_success,omitempty"`

    // 返回编码
    MsgCode   string `json:"msg_code,omitempty"`

    // 返回信息
    MsgInfo   string `json:"msg_info,omitempty"`

}
