package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcUserCancelAPIRequest
取消用户的消息服务 API请求
taobao.tmc.user.cancel

取消用户的消息服务 */
type TaobaoTmcUserCancelAPIRequest struct {
	model.Params
	// 用户昵称
	_nick string
	// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
	_userPlatform string
}

// NewTaobaoTmcUserCancelRequest 初始化TaobaoTmcUserCancelAPIRequest对象
func NewTaobaoTmcUserCancelRequest() *TaobaoTmcUserCancelAPIRequest {
	return &TaobaoTmcUserCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcUserCancelAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.user.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcUserCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 用户昵称
func (r *TaobaoTmcUserCancelAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoTmcUserCancelAPIRequest) GetNick() string {
	return r._nick
}

// Set is UserPlatform Setter
// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
func (r *TaobaoTmcUserCancelAPIRequest) SetUserPlatform(_userPlatform string) error {
	r._userPlatform = _userPlatform
	r.Set("user_platform", _userPlatform)
	return nil
}

// Get UserPlatform Getter
func (r TaobaoTmcUserCancelAPIRequest) GetUserPlatform() string {
	return r._userPlatform
}
