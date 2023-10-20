package tmc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcUserCancelAPIRequest 取消用户的消息服务 API请求
// taobao.tmc.user.cancel
//
// 取消用户的消息服务
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTmcUserCancelAPIRequest) Reset() {
	r._nick = ""
	r._userPlatform = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcUserCancelAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.user.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcUserCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTmcUserCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户昵称
func (r *TaobaoTmcUserCancelAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoTmcUserCancelAPIRequest) GetNick() string {
	return r._nick
}

// SetUserPlatform is UserPlatform Setter
// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
func (r *TaobaoTmcUserCancelAPIRequest) SetUserPlatform(_userPlatform string) error {
	r._userPlatform = _userPlatform
	r.Set("user_platform", _userPlatform)
	return nil
}

// GetUserPlatform UserPlatform Getter
func (r TaobaoTmcUserCancelAPIRequest) GetUserPlatform() string {
	return r._userPlatform
}

var poolTaobaoTmcUserCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTmcUserCancelRequest()
	},
}

// GetTaobaoTmcUserCancelRequest 从 sync.Pool 获取 TaobaoTmcUserCancelAPIRequest
func GetTaobaoTmcUserCancelAPIRequest() *TaobaoTmcUserCancelAPIRequest {
	return poolTaobaoTmcUserCancelAPIRequest.Get().(*TaobaoTmcUserCancelAPIRequest)
}

// ReleaseTaobaoTmcUserCancelAPIRequest 将 TaobaoTmcUserCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoTmcUserCancelAPIRequest(v *TaobaoTmcUserCancelAPIRequest) {
	v.Reset()
	poolTaobaoTmcUserCancelAPIRequest.Put(v)
}
