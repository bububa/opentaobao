package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubuserdutysgetAPIRequest 获取指定账户的所有职务信息列表 API请求
// taobao.subuser.dutys.get
//
// 通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息）
type TaobaosubuserdutysgetAPIRequest struct {
	model.Params
	// 主账号用户名
	_userNick string
}

// NewTaobaosubuserdutysgetRequest 初始化TaobaosubuserdutysgetAPIRequest对象
func NewTaobaosubuserdutysgetRequest() *TaobaosubuserdutysgetAPIRequest {
	return &TaobaosubuserdutysgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubuserdutysgetAPIRequest) GetApiMethodName() string {
	return "taobao.subuser.dutys.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubuserdutysgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubuserdutysgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 主账号用户名
func (r *TaobaosubuserdutysgetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaosubuserdutysgetAPIRequest) GetUserNick() string {
	return r._userNick
}
