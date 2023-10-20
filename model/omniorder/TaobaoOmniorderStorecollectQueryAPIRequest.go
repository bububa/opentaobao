package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStorecollectQueryAPIRequest 全渠道门店自提根据核销码查询订单 API请求
// taobao.omniorder.storecollect.query
//
// 全渠道门店自提根据核销码查询订单
type TaobaoOmniorderStorecollectQueryAPIRequest struct {
	model.Params
	// 核销码
	_code string
}

// NewTaobaoOmniorderStorecollectQueryRequest 初始化TaobaoOmniorderStorecollectQueryAPIRequest对象
func NewTaobaoOmniorderStorecollectQueryRequest() *TaobaoOmniorderStorecollectQueryAPIRequest {
	return &TaobaoOmniorderStorecollectQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderStorecollectQueryAPIRequest) Reset() {
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStorecollectQueryAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.storecollect.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStorecollectQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderStorecollectQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 核销码
func (r *TaobaoOmniorderStorecollectQueryAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoOmniorderStorecollectQueryAPIRequest) GetCode() string {
	return r._code
}

var poolTaobaoOmniorderStorecollectQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderStorecollectQueryRequest()
	},
}

// GetTaobaoOmniorderStorecollectQueryRequest 从 sync.Pool 获取 TaobaoOmniorderStorecollectQueryAPIRequest
func GetTaobaoOmniorderStorecollectQueryAPIRequest() *TaobaoOmniorderStorecollectQueryAPIRequest {
	return poolTaobaoOmniorderStorecollectQueryAPIRequest.Get().(*TaobaoOmniorderStorecollectQueryAPIRequest)
}

// ReleaseTaobaoOmniorderStorecollectQueryAPIRequest 将 TaobaoOmniorderStorecollectQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderStorecollectQueryAPIRequest(v *TaobaoOmniorderStorecollectQueryAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderStorecollectQueryAPIRequest.Put(v)
}
