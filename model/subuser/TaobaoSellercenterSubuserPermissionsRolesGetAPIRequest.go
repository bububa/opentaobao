package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosellercentersubuserpermissionsrolesgetAPIRequest 查询指定的子账号的权限和角色信息 API请求
// taobao.sellercenter.subuser.permissions.roles.get
//
// 查询指定的子账号的被直接赋予的权限信息和角色信息。&lt;br/&gt;返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。
type TaobaosellercentersubuserpermissionsrolesgetAPIRequest struct {
	model.Params
	// 子账号昵称(子账号标识)
	_nick string
}

// NewTaobaosellercentersubuserpermissionsrolesgetRequest 初始化TaobaosellercentersubuserpermissionsrolesgetAPIRequest对象
func NewTaobaosellercentersubuserpermissionsrolesgetRequest() *TaobaosellercentersubuserpermissionsrolesgetAPIRequest {
	return &TaobaosellercentersubuserpermissionsrolesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosellercentersubuserpermissionsrolesgetAPIRequest) GetApiMethodName() string {
	return "taobao.sellercenter.subuser.permissions.roles.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosellercentersubuserpermissionsrolesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosellercentersubuserpermissionsrolesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 子账号昵称(子账号标识)
func (r *TaobaosellercentersubuserpermissionsrolesgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosellercentersubuserpermissionsrolesgetAPIRequest) GetNick() string {
	return r._nick
}
