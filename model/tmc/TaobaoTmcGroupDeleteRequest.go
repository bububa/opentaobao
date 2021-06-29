package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除指定的分组或分组下的用户 API请求
taobao.tmc.group.delete

删除指定的分组或分组下的用户，授权消息使用
*/
type TaobaoTmcGroupDeleteRequest struct {
    model.Params
    // 分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。
    _groupName   string
    // 用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组
    _nicks   []string
    // 用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
    _userPlatform   string
}

// 初始化TaobaoTmcGroupDeleteRequest对象
func NewTaobaoTmcGroupDeleteRequest() *TaobaoTmcGroupDeleteRequest{
    return &TaobaoTmcGroupDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcGroupDeleteRequest) GetApiMethodName() string {
    return "taobao.tmc.group.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcGroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupName Setter
// 分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。
func (r *TaobaoTmcGroupDeleteRequest) SetGroupName(_groupName string) error {
    r._groupName = _groupName
    r.Set("group_name", _groupName)
    return nil
}

// GroupName Getter
func (r TaobaoTmcGroupDeleteRequest) GetGroupName() string {
    return r._groupName
}
// Nicks Setter
// 用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组
func (r *TaobaoTmcGroupDeleteRequest) SetNicks(_nicks []string) error {
    r._nicks = _nicks
    r.Set("nicks", _nicks)
    return nil
}

// Nicks Getter
func (r TaobaoTmcGroupDeleteRequest) GetNicks() []string {
    return r._nicks
}
// UserPlatform Setter
// 用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
func (r *TaobaoTmcGroupDeleteRequest) SetUserPlatform(_userPlatform string) error {
    r._userPlatform = _userPlatform
    r.Set("user_platform", _userPlatform)
    return nil
}

// UserPlatform Getter
func (r TaobaoTmcGroupDeleteRequest) GetUserPlatform() string {
    return r._userPlatform
}
