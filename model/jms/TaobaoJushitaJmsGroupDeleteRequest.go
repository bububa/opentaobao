package jms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除ONS分组 APIRequest
taobao.jushita.jms.group.delete

删除ONS分组
*/
type TaobaoJushitaJmsGroupDeleteRequest struct {
    model.Params

    // 分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。
    groupName   string 

    // 用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组
    nicks   []string 

    // 用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户
    userPlatform   string 

}

func NewTaobaoJushitaJmsGroupDeleteRequest() *TaobaoJushitaJmsGroupDeleteRequest{
    return &TaobaoJushitaJmsGroupDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJushitaJmsGroupDeleteRequest) GetApiMethodName() string {
    return "taobao.jushita.jms.group.delete"
}

func (r TaobaoJushitaJmsGroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJushitaJmsGroupDeleteRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

func (r TaobaoJushitaJmsGroupDeleteRequest) GetGroupName() string {
    return r.groupName
}

func (r *TaobaoJushitaJmsGroupDeleteRequest) SetNicks(nicks []string) error {
    r.nicks = nicks
    r.Set("nicks", nicks)
    return nil
}

func (r TaobaoJushitaJmsGroupDeleteRequest) GetNicks() []string {
    return r.nicks
}

func (r *TaobaoJushitaJmsGroupDeleteRequest) SetUserPlatform(userPlatform string) error {
    r.userPlatform = userPlatform
    r.Set("user_platform", userPlatform)
    return nil
}

func (r TaobaoJushitaJmsGroupDeleteRequest) GetUserPlatform() string {
    return r.userPlatform
}

