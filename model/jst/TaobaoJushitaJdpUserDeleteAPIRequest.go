package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJdpUserDeleteAPIRequest 删除数据推送用户 API请求
// taobao.jushita.jdp.user.delete
//
// 删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
type TaobaoJushitaJdpUserDeleteAPIRequest struct {
	model.Params
	// 要删除用户的昵称
	_nick string
}

// NewTaobaoJushitaJdpUserDeleteRequest 初始化TaobaoJushitaJdpUserDeleteAPIRequest对象
func NewTaobaoJushitaJdpUserDeleteRequest() *TaobaoJushitaJdpUserDeleteAPIRequest {
	return &TaobaoJushitaJdpUserDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJushitaJdpUserDeleteAPIRequest) Reset() {
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJdpUserDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jdp.user.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJdpUserDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJushitaJdpUserDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 要删除用户的昵称
func (r *TaobaoJushitaJdpUserDeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoJushitaJdpUserDeleteAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoJushitaJdpUserDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJushitaJdpUserDeleteRequest()
	},
}

// GetTaobaoJushitaJdpUserDeleteRequest 从 sync.Pool 获取 TaobaoJushitaJdpUserDeleteAPIRequest
func GetTaobaoJushitaJdpUserDeleteAPIRequest() *TaobaoJushitaJdpUserDeleteAPIRequest {
	return poolTaobaoJushitaJdpUserDeleteAPIRequest.Get().(*TaobaoJushitaJdpUserDeleteAPIRequest)
}

// ReleaseTaobaoJushitaJdpUserDeleteAPIRequest 将 TaobaoJushitaJdpUserDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoJushitaJdpUserDeleteAPIRequest(v *TaobaoJushitaJdpUserDeleteAPIRequest) {
	v.Reset()
	poolTaobaoJushitaJdpUserDeleteAPIRequest.Put(v)
}
