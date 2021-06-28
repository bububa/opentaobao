package game

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新用户激活通知接口 APIResponse
taobao.apple.newuser.activate.notify

资和信主动通知激活结果
*/
type TaobaoAppleNewuserActivateNotifyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAppleNewuserActivateNotifyResponse `json:"apple_newuser_activate_notify_response,omitempty"` 
    TaobaoAppleNewuserActivateNotifyResponse
}

/* model for simplify = false
type TaobaoAppleNewuserActivateNotifyResponse struct {

    // 处理结果说明
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // 处理结果码
    
    ResultCode   string `json:"result_code,omitempty"`
    

}
*/

type TaobaoAppleNewuserActivateNotifyResponse struct {

    // 处理结果说明
    ResultMsg   string `json:"result_msg,omitempty"`

    // 处理结果码
    ResultCode   string `json:"result_code,omitempty"`

}
