package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubusersgetAPIRequest 获取指定账户的子账号简易信息列表 API请求
// taobao.subusers.get
//
// 获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息）
type TaobaosubusersgetAPIRequest struct {
	model.Params
	// 主账号用户名
	_userNick string
}

// NewTaobaosubusersgetRequest 初始化TaobaosubusersgetAPIRequest对象
func NewTaobaosubusersgetRequest() *TaobaosubusersgetAPIRequest {
	return &TaobaosubusersgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubusersgetAPIRequest) GetApiMethodName() string {
	return "taobao.subusers.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubusersgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubusersgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 主账号用户名
func (r *TaobaosubusersgetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaosubusersgetAPIRequest) GetUserNick() string {
	return r._userNick
}
