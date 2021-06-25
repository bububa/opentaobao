package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量发送消息 APIResponse
taobao.tmc.messages.produce

批量发送消息
*/
type TaobaoTmcMessagesProduceAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTmcMessagesProduceResponse `json:"taobao_tmc_messages_produce_response,omitempty"`
}

type TaobaoTmcMessagesProduceResponse struct {

    // 是否全部成功
    IsAllSuccess   bool `json:"is_all_success,omitempty"`

    // 发送结果，与发送时的参数顺序一致。如果is_all_success为true时，不用校验result是否成功
    Results   []TmcProduceResult `json:"results,omitempty"`

}
