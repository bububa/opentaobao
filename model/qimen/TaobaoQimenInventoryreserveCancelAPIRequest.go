package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventoryreserveCancelAPIRequest 库存预占取消接口 API请求
// taobao.qimen.inventoryreserve.cancel
//
// 库存预占取消
type TaobaoQimenInventoryreserveCancelAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenInventoryreserveCancelRequest
}

// NewTaobaoQimenInventoryreserveCancelRequest 初始化TaobaoQimenInventoryreserveCancelAPIRequest对象
func NewTaobaoQimenInventoryreserveCancelRequest() *TaobaoQimenInventoryreserveCancelAPIRequest {
	return &TaobaoQimenInventoryreserveCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenInventoryreserveCancelAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventoryreserveCancelAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventoryreserve.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventoryreserveCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenInventoryreserveCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenInventoryreserveCancelAPIRequest) SetRequest(_request *TaobaoQimenInventoryreserveCancelRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenInventoryreserveCancelAPIRequest) GetRequest() *TaobaoQimenInventoryreserveCancelRequest {
	return r._request
}

var poolTaobaoQimenInventoryreserveCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenInventoryreserveCancelRequest()
	},
}

// GetTaobaoQimenInventoryreserveCancelRequest 从 sync.Pool 获取 TaobaoQimenInventoryreserveCancelAPIRequest
func GetTaobaoQimenInventoryreserveCancelAPIRequest() *TaobaoQimenInventoryreserveCancelAPIRequest {
	return poolTaobaoQimenInventoryreserveCancelAPIRequest.Get().(*TaobaoQimenInventoryreserveCancelAPIRequest)
}

// ReleaseTaobaoQimenInventoryreserveCancelAPIRequest 将 TaobaoQimenInventoryreserveCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenInventoryreserveCancelAPIRequest(v *TaobaoQimenInventoryreserveCancelAPIRequest) {
	v.Reset()
	poolTaobaoQimenInventoryreserveCancelAPIRequest.Put(v)
}
