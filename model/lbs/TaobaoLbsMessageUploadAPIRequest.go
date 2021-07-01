package lbs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
lbs数据采集 API请求
taobao.lbs.message.upload

lbs数据采集
*/
type TaobaoLbsMessageUploadAPIRequest struct {
    model.Params
    // 消息TOPIC
    _topic   string
    // 消息体 json结构
    _body   string
}

// 初始化TaobaoLbsMessageUploadAPIRequest对象
func NewTaobaoLbsMessageUploadRequest() *TaobaoLbsMessageUploadAPIRequest{
    return &TaobaoLbsMessageUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLbsMessageUploadAPIRequest) GetApiMethodName() string {
    return "taobao.lbs.message.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLbsMessageUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Topic Setter
// 消息TOPIC
func (r *TaobaoLbsMessageUploadAPIRequest) SetTopic(_topic string) error {
    r._topic = _topic
    r.Set("topic", _topic)
    return nil
}

// Topic Getter
func (r TaobaoLbsMessageUploadAPIRequest) GetTopic() string {
    return r._topic
}
// Body Setter
// 消息体 json结构
func (r *TaobaoLbsMessageUploadAPIRequest) SetBody(_body string) error {
    r._body = _body
    r.Set("body", _body)
    return nil
}

// Body Getter
func (r TaobaoLbsMessageUploadAPIRequest) GetBody() string {
    return r._body
}
