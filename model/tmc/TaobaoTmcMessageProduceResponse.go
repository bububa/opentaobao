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
	RequestId     string         `json:"request_id,omitempty" xml:"tmc_message_produce_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"