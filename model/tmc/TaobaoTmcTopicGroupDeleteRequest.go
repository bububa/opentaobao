package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除消息topic分组路由 APIRequest
taobao.tmc.topic.group.delete

删除根据topic名称路由消息到不同的分组关系
*/
type TaobaoTmcTopicGroupDeleteRequest struct {
    model.Params

    // 消息分组名
    groupName   string 

    // 消息topic名称，多个以逗号(,)分割
    topics   []String 

    // 消息分组Id，一般不用填写，如果分组已经被删除，则根据问题排查工具返回的ID删除路由关系
    groupId   int64 

}

func NewTaobaoTmcTopicGroupDeleteRequest() *TaobaoTmcTopicGroupDeleteRequest{
    return &TaobaoTmcTopicGroupDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTmcTopicGroupDeleteRequest) GetApiMethodName() string {
    return "taobao.tmc.topic.group.delete"
}

func (r TaobaoTmcTopicGroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTmcTopicGroupDeleteRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

func (r TaobaoTmcTopicGroupDeleteRequest) GetGroupName() string {
    return r.groupName
}

func (r *TaobaoTmcTopicGroupDeleteRequest) SetTopics(topics []String) error {
    r.topics = topics
    r.Set("topics", topics)
    return nil
}

func (r TaobaoTmcTopicGroupDeleteRequest) GetTopics() []String {
    return r.topics
}

func (r *TaobaoTmcTopicGroupDeleteRequest) SetGroupId(groupId int64) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r TaobaoTmcTopicGroupDeleteRequest) GetGroupId() int64 {
    return r.groupId
}

