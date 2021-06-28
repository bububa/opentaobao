package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
开始配送工单 APIResponse
tmall.servicecenter.workcard.delivery

服务商调用该接口通知天猫服务平台服务商工人已开始配送工单
*/
type TmallServicecenterWorkcardDeliveryAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterWorkcardDeliveryResponse `json:"tmall_servicecenter_workcard_delivery_response,omitempty"` 
    TmallServicecenterWorkcardDeliveryResponse
}

/* model for simplify = false
type TmallServicecenterWorkcardDeliveryResponse struct {

    // 是否成功
    
    MsgSuccess   bool `json:"msg_success,omitempty"`
    

    // 返回信息
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // 返回code
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

}
*/

type TmallServicecenterWorkcardDeliveryResponse struct {

    // 是否成功
    MsgSuccess   bool `json:"msg_success,omitempty"`

    // 返回信息
    MsgCode   string `json:"msg_code,omitempty"`

    // 返回code
    MsgInfo   string `json:"msg_info,omitempty"`

}
