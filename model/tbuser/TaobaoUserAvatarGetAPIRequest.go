package tbuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouseravatargetAPIRequest 淘宝用户头像查询 API请求
// taobao.user.avatar.get
//
// 根据混淆nick查询用户头像
type TaobaouseravatargetAPIRequest struct {
	model.Params
	// 混淆nick
	_nick string
}

// NewTaobaouseravatargetRequest 初始化TaobaouseravatargetAPIRequest对象
func NewTaobaouseravatargetRequest() *TaobaouseravatargetAPIRequest {
	return &TaobaouseravatargetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouseravatargetAPIRequest) GetApiMethodName() string {
	return "taobao.user.avatar.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouseravatargetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouseravatargetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 混淆nick
func (r *TaobaouseravatargetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaouseravatargetAPIRequest) GetNick() string {
	return r._nick
}
