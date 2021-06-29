package cmns

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息回执查询 APIResponse
yunos.service.cmns.coa.message.ack

第三方应用开发者调用此接口查询设备是否收到消息，只能查询此appKey床发的消息
*/
type YunosServiceCmnsCoaMessageAckAPIResponse struct {
    model.CommonResponse
    YunosServiceCmnsCoaMessageAckResponse
}

type YunosServiceCmnsCoaMessageAckResponse struct {
    XMLName xml.Name `xml:"yunos_service_cmns_coa_message_ack_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口调用成功
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`

    
    // 接口出错提示信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 接口调用成功返回信息<br/>0:未到达 1：已到达
    
    Data   int64 `json:"data,omitempty" xml:"data,omitempty"`

    
}
