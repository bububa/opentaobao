package subuser

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest 查询指定的子账号的权限和角色信息 API请求
// taobao.sellercenter.subuser.permissions.roles.get
//
// 查询指定的子账号的被直接赋予的权限信息和角色信息。&lt;br/&gt;返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。
type TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest struct {
	model.Params
	// 子账号昵称(子账号标识)
	_nick string
}

// NewTaobaoSellercenterSubuserPermissionsRolesGetRequest 初始化TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest对象
func NewTaobaoSellercenterSubuserPermissionsRolesGetRequest() *TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest {
	return &TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest) Reset() {
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest) GetApiMethodName() string {
	return "taobao.sellercenter.subuser.permissions.roles.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 子账号昵称(子账号标识)
func (r *TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoSellercenterSubuserPermissionsRolesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSellercenterSubuserPermissionsRolesGetRequest()
	},
}

// GetTaobaoSellercenterSubuserPermissionsRolesGetRequest 从 sync.Pool 获取 TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest
func GetTaobaoSellercenterSubuserPermissionsRolesGetAPIRequest() *TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest {
	return poolTaobaoSellercenterSubuserPermissionsRolesGetAPIRequest.Get().(*TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest)
}

// ReleaseTaobaoSellercenterSubuserPermissionsRolesGetAPIRequest 将 TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSellercenterSubuserPermissionsRolesGetAPIRequest(v *TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest) {
	v.Reset()
	poolTaobaoSellercenterSubuserPermissionsRolesGetAPIRequest.Put(v)
}
