package jms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加ONS消息同步用户 APIRequest
taobao.jushita.jms.user.add

添加ONS消息同步用户
*/
type TaobaoJushitaJmsUserAddRequest struct {
    model.Params

    // topic列表,不填则继承appkey所订阅的topic
    topicNames   []String 

}

func NewTaobaoJushitaJmsUserAddRequest() *TaobaoJushitaJmsUserAddRequest{
    return &TaobaoJushitaJmsUserAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJushitaJmsUserAddRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.user.add"
}

func (r TaobaoJushitaJmsUserAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJushitaJmsUserAddRequest) SetTopicNames(topicNames []String) error {
    r.topicNames = topicNames
    r.Set("topic_names", topicNames)
    return nil
}

func (r TaobaoJushitaJmsUserAddRequest) GetTopicNames() []String {
    return r.topicNames
}

