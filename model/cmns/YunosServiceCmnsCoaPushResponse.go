package cmns

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息推送接口 APIResponse
yunos.service.cmns.coa.push

调用CMNS系统的pushMessage接口实现消息通知到YUNOS设备的第三方应用软件。
*/
type YunosServiceCmnsCoaPushAPIResponse struct {
    model.CommonResponse
    YunosServiceCmnsCoaPushResponse
}

type YunosServiceCmnsCoaPushResponse struct {
    XMLName xml.Name `xml:"yunos_service_cmns_coa_push_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 200:消息发送成功
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`

    
    // 消息发送提示
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 消息ID，失败则为null
    
    Mid   int64 `json:"mid,omitempty" xml:"mid,omitempty"`

    
}
