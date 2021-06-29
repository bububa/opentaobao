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
type TaobaoTmcMessagesConsumeRequest struct {
    model.Params
    // 用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误
    groupName   string
    // 每次批量消费消息的条数，最小值：10；最大值：200
    quantity   int64
}

// 初始化TaobaoTmcMessagesConsumeRequest对象
func NewTaobaoTmcMessagesConsumeRequest() *TaobaoTmcMessagesConsumeRequest{
    return &TaobaoTmcMessagesConsumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcMessagesConsumeRequest) GetApiMethodName() string {
    return "taobao.tmc.messages.consume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcMessagesConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupName Setter
// 用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误
func (r *TaobaoTmcMessagesConsumeRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

// GroupName Getter
func (r TaobaoTmcMessagesConsumeRequest) GetGroupName() string {
    return r.groupName
}
// Quantity Setter
// 每次批量消费消息的条数，最小值：10；最大值：200
func (r *TaobaoTmcMessagesConsumeRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

// Quantity Getter
func (r TaobaoTmcMessagesConsumeRequest) GetQuantity() int64 {
    return r.quantity
}
