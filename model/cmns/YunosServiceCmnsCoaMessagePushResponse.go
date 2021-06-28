package cmns

import (
    "github.com/bububa/opentaobao/model"
)

/* 
消息推送接口 APIResponse
yunos.service.cmns.coa.message.push

调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
*/
type YunosServiceCmnsCoaMessagePushAPIResponse struct {
    model.CommonResponse
    // Response *YunosServiceCmnsCoaMessagePushResponse `json:"yunos_service_cmns_coa_message_push_response,omitempty"` 
    YunosServiceCmnsCoaMessagePushResponse
}

/* model for simplify = false
type YunosServiceCmnsCoaMessagePushResponse struct {

    // 消息ID，失败则为null
    
    Mid   int64 `json:"mid,omitempty"`
    

    // 消息发送提示
    
    Message   string `json:"message,omitempty"`
    

    // 200:消息发送成功
    
    Status   int64 `json:"status,omitempty"`
    

}
*/

type YunosServiceCmnsCoaMessagePushResponse struct {

    // 消息ID，失败则为null
    Mid   int64 `json:"mid,omitempty"`

    // 消息发送提示
    Message   string `json:"message,omitempty"`

    // 200:消息发送成功
    Status   int64 `json:"status,omitempty"`

}
