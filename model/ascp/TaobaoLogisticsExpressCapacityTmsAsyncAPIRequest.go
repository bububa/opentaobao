package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest 上门取退产能信息同步/更新 API请求
// taobao.logistics.express.capacity.tms.async
//
// 上门取退产能信息同步/更新
type TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest struct {
	model.Params
	// 上门取退产能信息同步/更新
	_capacityRequest *CapacityRequest
}

// NewTaobaoLogisticsExpressCapacityTmsAsyncRequest 初始化TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest对象
func NewTaobaoLogisticsExpressCapacityTmsAsyncRequest() *TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest {
	return &TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest) Reset() {
	r._capacityRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.capacity.tms.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCapacityRequest is CapacityRequest Setter
// 上门取退产能信息同步/更新
func (r *TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest) SetCapacityRequest(_capacityRequest *CapacityRequest) error {
	r._capacityRequest = _capacityRequest
	r.Set("capacity_request", _capacityRequest)
	return nil
}

// GetCapacityRequest CapacityRequest Getter
func (r TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest) GetCapacityRequest() *CapacityRequest {
	return r._capacityRequest
}

var poolTaobaoLogisticsExpressCapacityTmsAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsExpressCapacityTmsAsyncRequest()
	},
}

// GetTaobaoLogisticsExpressCapacityTmsAsyncRequest 从 sync.Pool 获取 TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest
func GetTaobaoLogisticsExpressCapacityTmsAsyncAPIRequest() *TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest {
	return poolTaobaoLogisticsExpressCapacityTmsAsyncAPIRequest.Get().(*TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest)
}

// ReleaseTaobaoLogisticsExpressCapacityTmsAsyncAPIRequest 将 TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsExpressCapacityTmsAsyncAPIRequest(v *TaobaoLogisticsExpressCapacityTmsAsyncAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsExpressCapacityTmsAsyncAPIRequest.Put(v)
}
