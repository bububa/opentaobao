package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUserAvatarGetAPIRequest
淘宝用户头像查询 API请求
taobao.user.avatar.get

根据混淆nick查询用户头像 */
type TaobaoUserAvatarGetAPIRequest struct {
	model.Params
	// 混淆nick
	_nick string
}

// NewTaobaoUserAvatarGetRequest 初始化TaobaoUserAvatarGetAPIRequest对象
func NewTaobaoUserAvatarGetRequest() *TaobaoUserAvatarGetAPIRequest {
	return &TaobaoUserAvatarGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUserAvatarGetAPIRequest) GetApiMethodName() string {
	return "taobao.user.avatar.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUserAvatarGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 混淆nick
func (r *TaobaoUserAvatarGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoUserAvatarGetAPIRequest) GetNick() string {
	return r._nick
}
