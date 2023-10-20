package tbuser

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserAvatarGetAPIRequest 淘宝用户头像查询 API请求
// taobao.user.avatar.get
//
// 根据混淆nick查询用户头像
type TaobaoUserAvatarGetAPIRequest struct {
	model.Params
	// 混淆nick
	_nick string
}

// NewTaobaoUserAvatarGetRequest 初始化TaobaoUserAvatarGetAPIRequest对象
func NewTaobaoUserAvatarGetRequest() *TaobaoUserAvatarGetAPIRequest {
	return &TaobaoUserAvatarGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUserAvatarGetAPIRequest) Reset() {
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUserAvatarGetAPIRequest) GetApiMethodName() string {
	return "taobao.user.avatar.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUserAvatarGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUserAvatarGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 混淆nick
func (r *TaobaoUserAvatarGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoUserAvatarGetAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoUserAvatarGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUserAvatarGetRequest()
	},
}

// GetTaobaoUserAvatarGetRequest 从 sync.Pool 获取 TaobaoUserAvatarGetAPIRequest
func GetTaobaoUserAvatarGetAPIRequest() *TaobaoUserAvatarGetAPIRequest {
	return poolTaobaoUserAvatarGetAPIRequest.Get().(*TaobaoUserAvatarGetAPIRequest)
}

// ReleaseTaobaoUserAvatarGetAPIRequest 将 TaobaoUserAvatarGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUserAvatarGetAPIRequest(v *TaobaoUserAvatarGetAPIRequest) {
	v.Reset()
	poolTaobaoUserAvatarGetAPIRequest.Put(v)
}
