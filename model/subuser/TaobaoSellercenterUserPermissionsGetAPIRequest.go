package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosellercenteruserpermissionsgetAPIRequest 获取指定用户的权限集合 API请求
// taobao.sellercenter.user.permissions.get
//
// 获取指定用户的权限集合，并不组装成树。如果是主账号，返回所有的权限列表；如果是子账号，返回所有已授权的权限。只能查询属于自己的账号信息 (如果是主账号，则是主账号以及所属子账号，如果是子账号则是对应主账号以及所属子账号)
type TaobaosellercenteruserpermissionsgetAPIRequest struct {
	model.Params
	// 用户标识，次入参必须为子账号比如zhangsan:cool。如果只输入主账号zhangsan，将报错。
	_nick string
}

// NewTaobaosellercenteruserpermissionsgetRequest 初始化TaobaosellercenteruserpermissionsgetAPIRequest对象
func NewTaobaosellercenteruserpermissionsgetRequest() *TaobaosellercenteruserpermissionsgetAPIRequest {
	return &TaobaosellercenteruserpermissionsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosellercenteruserpermissionsgetAPIRequest) GetApiMethodName() string {
	return "taobao.sellercenter.user.permissions.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosellercenteruserpermissionsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosellercenteruserpermissionsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户标识，次入参必须为子账号比如zhangsan:cool。如果只输入主账号zhangsan，将报错。
func (r *TaobaosellercenteruserpermissionsgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosellercenteruserpermissionsgetAPIRequest) GetNick() string {
	return r._nick
}
