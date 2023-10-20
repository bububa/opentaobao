package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpAccountGetBalanceAPIRequest 获取账户余额，现金余额 API请求
// taobao.universalbp.account.get.balance
//
// 获取账户实时现金余额
type TaobaoUniversalbpAccountGetBalanceAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaoUniversalbpAccountGetBalanceRequest 初始化TaobaoUniversalbpAccountGetBalanceAPIRequest对象
func NewTaobaoUniversalbpAccountGetBalanceRequest() *TaobaoUniversalbpAccountGetBalanceAPIRequest {
	return &TaobaoUniversalbpAccountGetBalanceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpAccountGetBalanceAPIRequest) Reset() {
	r._topServiceContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpAccountGetBalanceAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.account.get.balance"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpAccountGetBalanceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpAccountGetBalanceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpAccountGetBalanceAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpAccountGetBalanceAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

var poolTaobaoUniversalbpAccountGetBalanceAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpAccountGetBalanceRequest()
	},
}

// GetTaobaoUniversalbpAccountGetBalanceRequest 从 sync.Pool 获取 TaobaoUniversalbpAccountGetBalanceAPIRequest
func GetTaobaoUniversalbpAccountGetBalanceAPIRequest() *TaobaoUniversalbpAccountGetBalanceAPIRequest {
	return poolTaobaoUniversalbpAccountGetBalanceAPIRequest.Get().(*TaobaoUniversalbpAccountGetBalanceAPIRequest)
}

// ReleaseTaobaoUniversalbpAccountGetBalanceAPIRequest 将 TaobaoUniversalbpAccountGetBalanceAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpAccountGetBalanceAPIRequest(v *TaobaoUniversalbpAccountGetBalanceAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpAccountGetBalanceAPIRequest.Put(v)
}
