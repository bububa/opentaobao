package lbs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
lbs数据采集 APIRequest
taobao.lbs.message.upload

lbs数据采集
*/
type TaobaoLbsMessageUploadRequest struct {
    model.Params

    // 消息TOPIC
    topic   string 

    // 消息体 json结构
    body   string 

}

func NewTaobaoLbsMessageUploadRequest() *TaobaoLbsMessageUploadRequest{
    return &TaobaoLbsMessageUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLbsMessageUploadRequest) GetApiMethodName() string {
    return "taobao.lbs.message.upload"
}

func (r TaobaoLbsMessageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLbsMessageUploadRequest) SetTopic(topic string) error {
    r.topic = topic
    r.Set("topic", topic)
    return nil
}

func (r TaobaoLbsMessageUploadRequest) GetTopic() string {
    return r.topic
}

func (r *TaobaoLbsMessageUploadRequest) SetBody(body string) error {
    r.body = body
    r.Set("body", body)
    return nil
}

func (r TaobaoLbsMessageUploadRequest) GetBody() string {
    return r.body
}

