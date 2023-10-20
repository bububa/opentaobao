package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcusercancelAPIRequest 取消用户的消息服务 API请求
// taobao.tmc.user.cancel
//
// 取消用户的消息服务
type TaobaotmcusercancelAPIRequest struct {
	model.Params
	// 用户昵称
	_nick string
	// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
	_userPlatform string
}

// NewTaobaotmcusercancelRequest 初始化TaobaotmcusercancelAPIRequest对象
func NewTaobaotmcusercancelRequest() *TaobaotmcusercancelAPIRequest {
	return &TaobaotmcusercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotmcusercancelAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.user.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotmcusercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotmcusercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户昵称
func (r *TaobaotmcusercancelAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaotmcusercancelAPIRequest) GetNick() string {
	return r._nick
}

// SetUserPlatform is UserPlatform Setter
// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
func (r *TaobaotmcusercancelAPIRequest) SetUserPlatform(_userPlatform string) error {
	r._userPlatform = _userPlatform
	r.Set("user_platform", _userPlatform)
	return nil
}

// GetUserPlatform UserPlatform Getter
func (r TaobaotmcusercancelAPIRequest) GetUserPlatform() string {
	return r._userPlatform
}
