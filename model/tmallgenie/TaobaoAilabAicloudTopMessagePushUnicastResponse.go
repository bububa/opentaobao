package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵消息中心单播推送消息接口 APIResponse
taobao.ailab.aicloud.top.message.push.unicast

天猫精灵运营平台消息能力开放广播接口，主要开放给b端用户，用户可调用接口进行广播推送，将消息推送到天猫精灵设备或者天猫精灵APP中。
*/
type TaobaoAilabAicloudTopMessagePushUnicastAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopMessagePushUnicastResponse
}

type TaobaoAilabAicloudTopMessagePushUnicastResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_message_push_unicast_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 本次调用trace
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`

    
    // 调用结果code 0:成功 非0：失败
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 调用结果描述
    
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`

    
}
