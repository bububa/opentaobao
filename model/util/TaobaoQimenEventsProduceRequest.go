package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量发送奇门事件 API请求
taobao.qimen.events.produce

批量发送消息
*/
type TaobaoQimenEventsProduceRequest struct {
    model.Params
    // 奇门事件列表, 最多50条
    messages   []QimenEvent
}

// 初始化TaobaoQimenEventsProduceRequest对象
func NewTaobaoQimenEventsProduceRequest() *TaobaoQimenEventsProduceRequest{
    return &TaobaoQimenEventsProduceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenEventsProduceRequest) GetApiMethodName() string {
    return "taobao.qimen.events.produce"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenEventsProduceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Messages Setter
// 奇门事件列表, 最多50条
func (r *TaobaoQimenEventsProduceRequest) SetMessages(messages []QimenEvent) error {
    r.messages = messages
    r.Set("messages", messages)
    return nil
}

// Messages Getter
func (r TaobaoQimenEventsProduceRequest) GetMessages() []QimenEvent {
    return r.messages
}
