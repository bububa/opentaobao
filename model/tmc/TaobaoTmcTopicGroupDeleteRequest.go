package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除消息topic分组路由 API请求
taobao.tmc.topic.group.delete

删除根据topic名称路由消息到不同的分组关系
*/
type TaobaoTmcTopicGroupDeleteRequest struct {
    model.Params
    // 消息分组名
    _groupName   string
    // 消息topic名称，多个以逗号(,)分割
    _topics   []string
    // 消息分组Id，一般不用填写，如果分组已经被删除，则根据问题排查工具返回的ID删除路由关系
    _groupId   int64
}

// 初始化TaobaoTmcTopicGroupDeleteRequest对象
func NewTaobaoTmcTopicGroupDeleteRequest() *TaobaoTmcTopicGroupDeleteRequest{
    return &TaobaoTmcTopicGroupDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcTopicGroupDeleteRequest) GetApiMethodName() string {
    return "taobao.tmc.topic.group.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcTopicGroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupName Setter
// 消息分组名
func (r *TaobaoTmcTopicGroupDeleteRequest) SetGroupName(_groupName string) error {
    r._groupName = _groupName
    r.Set("group_name", _groupName)
    return nil
}

// GroupName Getter
func (r TaobaoTmcTopicGroupDeleteRequest) GetGroupName() string {
    return r._groupName
}
// Topics Setter
// 消息topic名称，多个以逗号(,)分割
func (r *TaobaoTmcTopicGroupDeleteRequest) SetTopics(_topics []string) error {
    r._topics = _topics
    r.Set("topics", _topics)
    return nil
}

// Topics Getter
func (r TaobaoTmcTopicGroupDeleteRequest) GetTopics() []string {
    return r._topics
}
// GroupId Setter
// 消息分组Id，一般不用填写，如果分组已经被删除，则根据问题排查工具返回的ID删除路由关系
func (r *TaobaoTmcTopicGroupDeleteRequest) SetGroupId(_groupId int64) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r TaobaoTmcTopicGroupDeleteRequest) GetGroupId() int64 {
    return r._groupId
}
