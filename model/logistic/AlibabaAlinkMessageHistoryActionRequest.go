package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
操作历史消息 API请求
alibaba.alink.message.history.action

阿里智能操作历史消息
*/
type AlibabaAlinkMessageHistoryActionRequest struct {
    model.Params
    // 消息id
    index   string
    // 删除：delete,已读：read
    action   string
}

// 初始化AlibabaAlinkMessageHistoryActionRequest对象
func NewAlibabaAlinkMessageHistoryActionRequest() *AlibabaAlinkMessageHistoryActionRequest{
    return &AlibabaAlinkMessageHistoryActionRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlinkMessageHistoryActionRequest) GetApiMethodName() string {
    return "alibaba.alink.message.history.action"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlinkMessageHistoryActionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Index Setter
// 消息id
func (r *AlibabaAlinkMessageHistoryActionRequest) SetIndex(index string) error {
    r.index = index
    r.Set("index", index)
    return nil
}

// Index Getter
func (r AlibabaAlinkMessageHistoryActionRequest) GetIndex() string {
    return r.index
}
// Action Setter
// 删除：delete,已读：read
func (r *AlibabaAlinkMessageHistoryActionRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

// Action Getter
func (r AlibabaAlinkMessageHistoryActionRequest) GetAction() string {
    return r.action
}
