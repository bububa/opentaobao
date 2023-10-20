package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaLoginAuthsignGetAPIRequest 获取登陆权限签名 API请求
// taobao.simba.login.authsign.get
//
// 获取登陆权限签名
type TaobaoSimbaLoginAuthsignGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
}

// NewTaobaoSimbaLoginAuthsignGetRequest 初始化TaobaoSimbaLoginAuthsignGetAPIRequest对象
func NewTaobaoSimbaLoginAuthsignGetRequest() *TaobaoSimbaLoginAuthsignGetAPIRequest {
	return &TaobaoSimbaLoginAuthsignGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaLoginAuthsignGetAPIRequest) Reset() {
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaLoginAuthsignGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.login.authsign.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaLoginAuthsignGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaLoginAuthsignGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaLoginAuthsignGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaLoginAuthsignGetAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoSimbaLoginAuthsignGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaLoginAuthsignGetRequest()
	},
}

// GetTaobaoSimbaLoginAuthsignGetRequest 从 sync.Pool 获取 TaobaoSimbaLoginAuthsignGetAPIRequest
func GetTaobaoSimbaLoginAuthsignGetAPIRequest() *TaobaoSimbaLoginAuthsignGetAPIRequest {
	return poolTaobaoSimbaLoginAuthsignGetAPIRequest.Get().(*TaobaoSimbaLoginAuthsignGetAPIRequest)
}

// ReleaseTaobaoSimbaLoginAuthsignGetAPIRequest 将 TaobaoSimbaLoginAuthsignGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaLoginAuthsignGetAPIRequest(v *TaobaoSimbaLoginAuthsignGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaLoginAuthsignGetAPIRequest.Put(v)
}
