package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Weex页面调用windvane APIResponse
alibaba.interact.windvane.call

客户端鉴权使用，实际不会发送或接收数据
*/
type AlibabaInteractWindvaneCallAPIResponse struct {
    model.CommonResponse
    AlibabaInteractWindvaneCallResponse
}

type AlibabaInteractWindvaneCallResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_windvane_call_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 客户端鉴权使用，实际不会发送或接收数据
    
    Unnamed   string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`

    
}
