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
type TaobaoLbsMessageUploadRequest struct {
    model.Params
    // 消息TOPIC
    _topic   string
    // 消息体 json结构
    _body   string
}

// 初始化TaobaoLbsMessageUploadRequest对象
func NewTaobaoLbsMessageUploadRequest() *TaobaoLbsMessageUploadRequest{
    return &TaobaoLbsMessageUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLbsMessageUploadRequest) GetApiMethodName() string {
    return "taobao.lbs.message.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLbsMessageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Topic Setter
// 消息TOPIC
func (r *TaobaoLbsMessageUploadRequest) SetTopic(_topic string) error {
    r._topic = _topic
    r.Set("topic", _topic)
    return nil
}

// Topic Getter
func (r TaobaoLbsMessageUploadRequest) GetTopic() string {
    return r._topic
}
// Body Setter
// 消息体 json结构
func (r *TaobaoLbsMessageUploadRequest) SetBody(_body string) error {
    r._body = _body
    r.Set("body", _body)
    return nil
}

// Body Getter
func (r TaobaoLbsMessageUploadRequest) GetBody() string {
    return r._body
}
