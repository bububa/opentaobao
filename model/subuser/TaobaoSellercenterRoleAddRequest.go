package subuser

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
子账号角色的新增（指定卖家） API请求
taobao.sellercenter.role.add

给指定的卖家创建新的子账号角色<br/><br/>如果需要授权的权限点有下级权限点或上级权限点，把该权限点的父权限点和该权限点的所有子权限都一并做赋权操作，并递归处理<br/>例如：权限点列表如下<br/><br/>code=sell 宝贝管理<br/><br/>---------|code=sm 店铺管理<br/><br/>---------|---------|code=sm-design 如店铺装修<br/><br/>---------|---------|---------|code=sm-tbd-visit内店装修入口<br/><br/>---------|---------|---------|code=sm-tbd-publish内店装修发布<br/><br/>---------|---------|code=phone 手机淘宝店铺<br/><br/>调用改接口给code=sm-design店铺装修赋权时，同时会将下列权限点都赋予默认角色<br/><br/>code=sell 宝贝管理<br/><br/>---------|code=sm 店铺管理<br/><br/>---------|---------|code=sm-design 如店铺装修<br/><br/>---------|---------|---------|code=sm-tbd-visit内店装修入口<br/><br/>---------|---------|---------|code=sm-tbd-publish内店装修发布<br/>
*/
type TaobaoSellercenterRoleAddRequest struct {
    model.Params
    // 角色名
    name   string
    // 角色描述
    description   string
    // 需要授权的权限点permission_code列表,以","分割.其code值可以通过调用taobao.sellercenter.user.permissions.get返回，其中permission.is_authorize=1的权限点可以通过本接口授权给对应角色。
    permissionCodes   []string
    // 表示卖家昵称
    nick   string
}

// 初始化TaobaoSellercenterRoleAddRequest对象
func NewTaobaoSellercenterRoleAddRequest() *TaobaoSellercenterRoleAddRequest{
    return &TaobaoSellercenterRoleAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSellercenterRoleAddRequest) GetApiMethodName() string {
    return "taobao.sellercenter.role.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSellercenterRoleAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 角色名
func (r *TaobaoSellercenterRoleAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoSellercenterRoleAddRequest) GetName() string {
    return r.name
}
// Description Setter
// 角色描述
func (r *TaobaoSellercenterRoleAddRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r TaobaoSellercenterRoleAddRequest) GetDescription() string {
    return r.description
}
// PermissionCodes Setter
// 需要授权的权限点permission_code列表,以","分割.其code值可以通过调用taobao.sellercenter.user.permissions.get返回，其中permission.is_authorize=1的权限点可以通过本接口授权给对应角色。
func (r *TaobaoSellercenterRoleAddRequest) SetPermissionCodes(permissionCodes []string) error {
    r.permissionCodes = permissionCodes
    r.Set("permission_codes", permissionCodes)
    return nil
}

// PermissionCodes Getter
func (r TaobaoSellercenterRoleAddRequest) GetPermissionCodes() []string {
    return r.permissionCodes
}
// Nick Setter
// 表示卖家昵称
func (r *TaobaoSellercenterRoleAddRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSellercenterRoleAddRequest) GetNick() string {
    return r.nick
}
