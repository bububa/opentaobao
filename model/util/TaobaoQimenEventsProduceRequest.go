package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量发送奇门事件 APIRequest
taobao.qimen.events.produce

批量发送消息
*/
type TaobaoQimenEventsProduceRequest struct {
    model.Params

    // 奇门事件列表, 最多50条
    messages   []QimenEvent 

}

func NewTaobaoQimenEventsProduceRequest() *TaobaoQimenEventsProduceRequest{
    return &TaobaoQimenEventsProduceRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenEventsProduceRequest) GetApiMethodName() string {
    return "taobao.qimen.events.produce"
}

func (r TaobaoQimenEventsProduceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenEventsProduceRequest) SetMessages(messages []QimenEvent) error {
    r.messages = messages
    r.Set("messages", messages)
    return nil
}

func (r TaobaoQimenEventsProduceRequest) GetMessages() []QimenEvent {
    return r.messages
}

