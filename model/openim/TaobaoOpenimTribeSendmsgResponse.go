package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发送群消息 APIResponse
taobao.openim.tribe.sendmsg

发送群消息，目前支持发送4种类型的群消息，普通文本，图片，语音，自定义消息
*/
type TaobaoOpenimTribeSendmsgAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimTribeSendmsgResponse
}

type TaobaoOpenimTribeSendmsgResponse struct {
    XMLName xml.Name `xml:"openim_tribe_sendmsg_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    TribeCode   int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`

    
    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
