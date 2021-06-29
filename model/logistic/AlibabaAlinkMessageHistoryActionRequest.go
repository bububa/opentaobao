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
    _index   string
    // 删除：delete,已读：read
    _action   string
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
func (r *AlibabaAlinkMessageHistoryActionRequest) SetIndex(_index string) error {
    r._index = _index
    r.Set("index", _index)
    return nil
}

// Index Getter
func (r AlibabaAlinkMessageHistoryActionRequest) GetIndex() string {
    return r._index
}
// Action Setter
// 删除：delete,已读：read
func (r *AlibabaAlinkMessageHistoryActionRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r AlibabaAlinkMessageHistoryActionRequest) GetAction() string {
    return r._action
}
