package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发布单条消息 APIResponse
taobao.tmc.message.produce

发布单条消息
*/
type TaobaoTmcMessageProduceAPIResponse struct {
    model.CommonResponse
    TaobaoTmcMessageProduceResponse
}

type TaobaoTmcMessageProduceResponse struct {
    XMLName xml.Name `xml:"tmc_message_produce_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 投递目标数
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`

    
    // 消息ID
    
    MsgIds   []string `json:"msg_ids,omitempty" xml:"msg_ids>string,omitempty"`
    
    
}
