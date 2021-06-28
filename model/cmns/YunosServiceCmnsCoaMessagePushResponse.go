package cmns

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息推送接口 APIResponse
yunos.service.cmns.coa.message.push

调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
*/
type YunosServiceCmnsCoaMessagePushAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"yunos_service_cmns_coa_message_push_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 消息ID，失败则为null
    
    Mid   int64 `json:"mid,omitempty" xml:"