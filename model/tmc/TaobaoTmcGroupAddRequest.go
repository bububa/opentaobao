package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
为已开通用户添加用户分组 APIRequest
taobao.tmc.group.add

为已开通用户添加用户分组，授权消息使用
*/
type TaobaoTmcGroupAddRequest struct {
    model.Params

    // 分组名称，同一个应用下需要保证唯一性，最长32个字符。添加分组后，消息通道会为用户的消息分配独立分组，但之前的消息还是存储于默认分组中。不能以default开头，default开头为系统默认组。
    groupName   string 

    // 用户昵称列表，以半角逗号分隔，支持子账号，支持增量添加用户
    nicks   []String 

    // 用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
    userPlatform   string 

}

func NewTaobaoTmcGroupAddRequest() *TaobaoTmcGroupAddRequest{
    return &TaobaoTmcGroupAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTmcGroupAddRequest) GetApiMethodName() string {
    return "taobao.tmc.group.add"
}

func (r TaobaoTmcGroupAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTmcGroupAddRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

func (r TaobaoTmcGroupAddRequest) GetGroupName() string {
    return r.groupName
}

func (r *TaobaoTmcGroupAddRequest) SetNicks(nicks []String) error {
    r.nicks = nicks
    r.Set("nicks", nicks)
    return nil
}

func (r TaobaoTmcGroupAddRequest) GetNicks() []String {
    return r.nicks
}

func (r *TaobaoTmcGroupAddRequest) SetUserPlatform(userPlatform string) error {
    r.userPlatform = userPlatform
    r.Set("user_platform", userPlatform)
    return nil
}

func (r TaobaoTmcGroupAddRequest) GetUserPlatform() string {
    return r.userPlatform
}

