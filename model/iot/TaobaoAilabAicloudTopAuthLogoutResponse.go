package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
登出 APIResponse
taobao.ailab.aicloud.top.auth.logout

登出
*/
type TaobaoAilabAicloudTopAuthLogoutAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopAuthLogoutResponse
}

type TaobaoAilabAicloudTopAuthLogoutResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_auth_logout_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // msgInfo错误码信息，成功返回null
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
