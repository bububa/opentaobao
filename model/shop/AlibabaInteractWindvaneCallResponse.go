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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_windvane_call_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 客户端鉴权使用，实际不会发送或接收数据
    
    Unnamed   string `json:"unnamed,omitempty" xml:"