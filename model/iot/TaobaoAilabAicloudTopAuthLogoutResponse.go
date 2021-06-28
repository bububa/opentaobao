package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
登出 APIResponse
taobao.ailab.aicloud.top.auth.logout

登出
*/
type TaobaoAilabAicloudTopAuthLogoutAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopAuthLogoutResponse `json:"ailab_aicloud_top_auth_logout_response,omitempty"` 
    TaobaoAilabAicloudTopAuthLogoutResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopAuthLogoutResponse struct {

    // msgInfo错误码信息，成功返回null
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopAuthLogoutResponse struct {

    // msgInfo错误码信息，成功返回null
    MsgInfo   string `json:"msg_info,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
