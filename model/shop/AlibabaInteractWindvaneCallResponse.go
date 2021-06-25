package shop

import (
    "github.com/bububa/opentaobao/model"
)

/* 
Weex页面调用windvane APIResponse
alibaba.interact.windvane.call

客户端鉴权使用，实际不会发送或接收数据
*/
type AlibabaInteractWindvaneCallAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractWindvaneCallResponse `json:"alibaba_interact_windvane_call_response,omitempty"`
}

type AlibabaInteractWindvaneCallResponse struct {

    // 客户端鉴权使用，实际不会发送或接收数据
    Unnamed   string `json:"unnamed,omitempty"`

}
