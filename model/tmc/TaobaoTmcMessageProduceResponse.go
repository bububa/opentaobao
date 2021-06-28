package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发布单条消息 APIResponse
taobao.tmc.message.produce

发布单条消息
*/
type TaobaoTmcMessageProduceAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTmcMessageProduceResponse `json:"tmc_message_produce_response,omitempty"` 
    TaobaoTmcMessageProduceResponse
}

/* model for simplify = false
type TaobaoTmcMessageProduceResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 投递目标数
    
    Total   int64 `json:"total,omitempty"`
    

    // 消息ID
    
    MsgIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"msg_ids,omitempty"`
    

}
*/

type TaobaoTmcMessageProduceResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 投递目标数
    Total   int64 `json:"total,omitempty"`

    // 消息ID
    MsgIds   []string `json:"msg_ids,omitempty"`

}
