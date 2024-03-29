package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAccountBalanceGetAPIRequest 获取实时余额，”元”为单位 API请求
// taobao.simba.account.balance.get
//
// 获取实时余额，”元”为单位
type TaobaoSimbaAccountBalanceGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
}

// NewTaobaoSimbaAccountBalanceGetRequest 初始化TaobaoSimbaAccountBalanceGetAPIRequest对象
func NewTaobaoSimbaAccountBalanceGetRequest() *TaobaoSimbaAccountBalanceGetAPIRequest {
	return &TaobaoSimbaAccountBalanceGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaAccountBalanceGetAPIRequest) Reset() {
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAccountBalanceGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.account.balance.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAccountBalanceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaAccountBalanceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaAccountBalanceGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaAccountBalanceGetAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoSimbaAccountBalanceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaAccountBalanceGetRequest()
	},
}

// GetTaobaoSimbaAccountBalanceGetRequest 从 sync.Pool 获取 TaobaoSimbaAccountBalanceGetAPIRequest
func GetTaobaoSimbaAccountBalanceGetAPIRequest() *TaobaoSimbaAccountBalanceGetAPIRequest {
	return poolTaobaoSimbaAccountBalanceGetAPIRequest.Get().(*TaobaoSimbaAccountBalanceGetAPIRequest)
}

// ReleaseTaobaoSimbaAccountBalanceGetAPIRequest 将 TaobaoSimbaAccountBalanceGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaAccountBalanceGetAPIRequest(v *TaobaoSimbaAccountBalanceGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaAccountBalanceGetAPIRequest.Put(v)
}
