package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消费多条消息 API请求
taobao.tmc.messages.consume

消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。
*/
type TaobaoTmcMessagesConsumeAPIRequest struct {
    model.Params
    // 用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误
    _groupName   string
    // 每次批量消费消息的条数，最小值：10；最大值：200
    _quantity   int64
}

// 初始化TaobaoTmcMessagesConsumeAPIRequest对象
func NewTaobaoTmcMessagesConsumeRequest() *TaobaoTmcMessagesConsumeAPIRequest{
    return &TaobaoTmcMessagesConsumeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcMessagesConsumeAPIRequest) GetApiMethodName() string {
    return "taobao.tmc.messages.consume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcMessagesConsumeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupName Setter
// 用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误
func (r *TaobaoTmcMessagesConsumeAPIRequest) SetGroupName(_groupName string) error {
    r._groupName = _groupName
    r.Set("group_name", _groupName)
    return nil
}

// GroupName Getter
func (r TaobaoTmcMessagesConsumeAPIRequest) GetGroupName() string {
    return r._groupName
}
// Quantity Setter
// 每次批量消费消息的条数，最小值：10；最大值：200
func (r *TaobaoTmcMessagesConsumeAPIRequest) SetQuantity(_quantity int64) error {
    r._quantity = _quantity
    r.Set("quantity", _quantity)
    return nil
}

// Quantity Getter
func (r TaobaoTmcMessagesConsumeAPIRequest) GetQuantity() int64 {
    return r._quantity
}
