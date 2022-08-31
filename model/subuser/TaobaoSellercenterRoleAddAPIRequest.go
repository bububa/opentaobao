package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercenterRoleAddAPIRequest 子账号角色的新增（指定卖家） API请求
// taobao.sellercenter.role.add
//
// 给指定的卖家创建新的子账号角色&lt;br/&gt;&lt;br/&gt;如果需要授权的权限点有下级权限点或上级权限点，把该权限点的父权限点和该权限点的所有子权限都一并做赋权操作，并递归处理&lt;br/&gt;例如：权限点列表如下&lt;br/&gt;&lt;br/&gt;code=sell 宝贝管理&lt;br/&gt;&lt;br/&gt;---------|code=sm 店铺管理&lt;br/&gt;&lt;br/&gt;---------|---------|code=sm-design 如店铺装修&lt;br/&gt;&lt;br/&gt;---------|---------|---------|code=sm-tbd-visit内店装修入口&lt;br/&gt;&lt;br/&gt;---------|---------|---------|code=sm-tbd-publish内店装修发布&lt;br/&gt;&lt;br/&gt;---------|---------|code=phone 手机淘宝店铺&lt;br/&gt;&lt;br/&gt;调用改接口给code=sm-design店铺装修赋权时，同时会将下列权限点都赋予默认角色&lt;br/&gt;&lt;br/&gt;code=sell 宝贝管理&lt;br/&gt;&lt;br/&gt;---------|code=sm 店铺管理&lt;br/&gt;&lt;br/&gt;---------|---------|code=sm-design 如店铺装修&lt;br/&gt;&lt;br/&gt;---------|---------|---------|code=sm-tbd-visit内店装修入口&lt;br/&gt;&lt;br/&gt;---------|---------|---------|code=sm-tbd-publish内店装修发布&lt;br/&gt;
type TaobaoSellercenterRoleAddAPIRequest struct {
	model.Params
	// 需要授权的权限点permission_code列表,以","分割.其code值可以通过调用taobao.sellercenter.user.permissions.get返回，其中permission.is_authorize=1的权限点可以通过本接口授权给对应角色。
	_permissionCodes []string
	// 表示卖家昵称
	_nick string
	// 角色名
	_name string
	// 角色描述
	_description string
}

// NewTaobaoSellercenterRoleAddRequest 初始化TaobaoSellercenterRoleAddAPIRequest对象
func NewTaobaoSellercenterRoleAddRequest() *TaobaoSellercenterRoleAddAPIRequest {
	return &TaobaoSellercenterRoleAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSellercenterRoleAddAPIRequest) GetApiMethodName() string {
	return "taobao.sellercenter.role.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSellercenterRoleAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPermissionCodes is PermissionCodes Setter
// 需要授权的权限点permission_code列表,以&#34;,&#34;分割.其code值可以通过调用taobao.sellercenter.user.permissions.get返回，其中permission.is_authorize=1的权限点可以通过本接口授权给对应角色。
func (r *TaobaoSellercenterRoleAddAPIRequest) SetPermissionCodes(_permissionCodes []string) error {
	r._permissionCodes = _permissionCodes
	r.Set("permission_codes", _permissionCodes)
	return nil
}

// GetPermissionCodes PermissionCodes Getter
func (r TaobaoSellercenterRoleAddAPIRequest) GetPermissionCodes() []string {
	return r._permissionCodes
}

// SetNick is Nick Setter
// 表示卖家昵称
func (r *TaobaoSellercenterRoleAddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSellercenterRoleAddAPIRequest) GetNick() string {
	return r._nick
}

// SetName is Name Setter
// 角色名
func (r *TaobaoSellercenterRoleAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoSellercenterRoleAddAPIRequest) GetName() string {
	return r._name
}

// SetDescription is Description Setter
// 角色描述
func (r *TaobaoSellercenterRoleAddAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r TaobaoSellercenterRoleAddAPIRequest) GetDescription() string {
	return r._description
}
