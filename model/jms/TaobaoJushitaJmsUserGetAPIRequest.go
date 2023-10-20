package jms

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsUserGetAPIRequest 查询某个用户是否同步消息 API请求
// taobao.jushita.jms.user.get
//
// 查询某个用户是否同步消息，只支持单个查询
type TaobaoJushitaJmsUserGetAPIRequest struct {
	model.Params
	// 需要查询的用户名
	_userNick string
}

// NewTaobaoJushitaJmsUserGetRequest 初始化TaobaoJushitaJmsUserGetAPIRequest对象
func NewTaobaoJushitaJmsUserGetRequest() *TaobaoJushitaJmsUserGetAPIRequest {
	return &TaobaoJushitaJmsUserGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJushitaJmsUserGetAPIRequest) Reset() {
	r._userNick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsUserGetAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jms.user.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsUserGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJushitaJmsUserGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 需要查询的用户名
func (r *TaobaoJushitaJmsUserGetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoJushitaJmsUserGetAPIRequest) GetUserNick() string {
	return r._userNick
}

var poolTaobaoJushitaJmsUserGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJushitaJmsUserGetRequest()
	},
}

// GetTaobaoJushitaJmsUserGetRequest 从 sync.Pool 获取 TaobaoJushitaJmsUserGetAPIRequest
func GetTaobaoJushitaJmsUserGetAPIRequest() *TaobaoJushitaJmsUserGetAPIRequest {
	return poolTaobaoJushitaJmsUserGetAPIRequest.Get().(*TaobaoJushitaJmsUserGetAPIRequest)
}

// ReleaseTaobaoJushitaJmsUserGetAPIRequest 将 TaobaoJushitaJmsUserGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoJushitaJmsUserGetAPIRequest(v *TaobaoJushitaJmsUserGetAPIRequest) {
	v.Reset()
	poolTaobaoJushitaJmsUserGetAPIRequest.Put(v)
}
