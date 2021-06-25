package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量发送消息 APIRequest
taobao.tmc.messages.produce

批量发送消息
*/
type TaobaoTmcMessagesProduceRequest struct {
    model.Params

    // tmc消息列表, 最多50条，元素结构与taobao.tmc.message.produce一致，用json表示的消息列表。例如：[{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"},{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"}]
    messages   []TmcPublishMessage 

}

func NewTaobaoTmcMessagesProduceRequest() *TaobaoTmcMessagesProduceRequest{
    return &TaobaoTmcMessagesProduceRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTmcMessagesProduceRequest) GetApiMethodName() string {
    return "taobao.tmc.messages.produce"
}

func (r TaobaoTmcMessagesProduceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTmcMessagesProduceRequest) SetMessages(messages []TmcPublishMessage) error {
    r.messages = messages
    r.Set("messages", messages)
    return nil
}

func (r TaobaoTmcMessagesProduceRequest) GetMessages() []TmcPublishMessage {
    return r.messages
}

