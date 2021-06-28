package cmns

import (
    "github.com/bububa/opentaobao/model"
)

/* 
消息推送接口 APIResponse
yunos.service.cmns.coa.push

调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
*/
type YunosServiceCmnsCoaPushAPIResponse struct {
    model.CommonResponse
    // Response *YunosServiceCmnsCoaPushResponse `json:"yunos_service_cmns_coa_push_response,omitempty"` 
    YunosServiceCmnsCoaPushResponse
}

/* model for simplify = false
type YunosServiceCmnsCoaPushResponse struct {

    // 200:消息发送成功
    
    Status   int64 `json:"status,omitempty"`
    

    // 消息发送提示
    
    Message   string `json:"message,omitempty"`
    

    // 消息ID，失败则为null
    
    Mid   int64 `json:"mid,omitempty"`
    

}
*/

type YunosServiceCmnsCoaPushResponse struct {

    // 200:消息发送成功
    Status   int64 `json:"status,omitempty"`

    // 消息发送提示
    Message   string `json:"message,omitempty"`

    // 消息ID，失败则为null
    Mid   int64 `json:"mid,omitempty"`

}
