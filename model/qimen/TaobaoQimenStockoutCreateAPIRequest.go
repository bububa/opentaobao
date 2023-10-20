package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStockoutCreateAPIRequest 出库单创建接口 API请求
// taobao.qimen.stockout.create
//
// taobao.qimen.returnorder.create
type TaobaoQimenStockoutCreateAPIRequest struct {
	model.Params
	//
	_request *StockOutCreateRequest
}

// NewTaobaoQimenStockoutCreateRequest 初始化TaobaoQimenStockoutCreateAPIRequest对象
func NewTaobaoQimenStockoutCreateRequest() *TaobaoQimenStockoutCreateAPIRequest {
	return &TaobaoQimenStockoutCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenStockoutCreateAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStockoutCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.stockout.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStockoutCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenStockoutCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenStockoutCreateAPIRequest) SetRequest(_request *StockOutCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenStockoutCreateAPIRequest) GetRequest() *StockOutCreateRequest {
	return r._request
}

var poolTaobaoQimenStockoutCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenStockoutCreateRequest()
	},
}

// GetTaobaoQimenStockoutCreateRequest 从 sync.Pool 获取 TaobaoQimenStockoutCreateAPIRequest
func GetTaobaoQimenStockoutCreateAPIRequest() *TaobaoQimenStockoutCreateAPIRequest {
	return poolTaobaoQimenStockoutCreateAPIRequest.Get().(*TaobaoQimenStockoutCreateAPIRequest)
}

// ReleaseTaobaoQimenStockoutCreateAPIRequest 将 TaobaoQimenStockoutCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenStockoutCreateAPIRequest(v *TaobaoQimenStockoutCreateAPIRequest) {
	v.Reset()
	poolTaobaoQimenStockoutCreateAPIRequest.Put(v)
}
