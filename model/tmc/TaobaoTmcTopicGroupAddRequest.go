package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
topic分组路由 API请求
taobao.tmc.topic.group.add

根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由
*/
type TaobaoTmcTopicGroupAddRequest struct {
    model.Params
    // 消息分组名，如果不存在，会自动创建
    groupName   string
    // 消息topic名称，多个以逗号(,)分割
    topics   []string
}

// 初始化TaobaoTmcTopicGroupAddRequest对象
func NewTaobaoTmcTopicGroupAddRequest() *TaobaoTmcTopicGroupAddRequest{
    return &TaobaoTmcTopicGroupAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcTopicGroupAddRequest) GetApiMethodName() string {
    return "taobao.tmc.topic.group.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcTopicGroupAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupName Setter
// 消息分组名，如果不存在，会自动创建
func (r *TaobaoTmcTopicGroupAddRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

// GroupName Getter
func (r TaobaoTmcTopicGroupAddRequest) GetGroupName() string {
    return r.groupName
}
// Topics Setter
// 消息topic名称，多个以逗号(,)分割
func (r *TaobaoTmcTopicGroupAddRequest) SetTopics(topics []string) error {
    r.topics = topics
    r.Set("topics", topics)
    return nil
}

// Topics Getter
func (r TaobaoTmcTopicGroupAddRequest) GetTopics() []string {
    return r.topics
}
