package jms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加ONS消息同步用户 API请求
taobao.jushita.jms.user.add

添加ONS消息同步用户
*/
type TaobaoJushitaJmsUserAddRequest struct {
    model.Params
    // topic列表,不填则继承appkey所订阅的topic
    _topicNames   []string
}

// 初始化TaobaoJushitaJmsUserAddRequest对象
func NewTaobaoJushitaJmsUserAddRequest() *TaobaoJushitaJmsUserAddRequest{
    return &TaobaoJushitaJmsUserAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsUserAddRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.user.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsUserAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopicNames Setter
// topic列表,不填则继承appkey所订阅的topic
func (r *TaobaoJushitaJmsUserAddRequest) SetTopicNames(_topicNames []string) error {
    r._topicNames = _topicNames
    r.Set("topic_names", _topicNames)
    return nil
}

// TopicNames Getter
func (r TaobaoJushitaJmsUserAddRequest) GetTopicNames() []string {
    return r._topicNames
}
