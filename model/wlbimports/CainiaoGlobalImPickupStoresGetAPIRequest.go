package wlbimports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupStoresGetAPIRequest 首公里揽收-集货仓列表查询 API请求
// cainiao.global.im.pickup.stores.get
//
// 首公里揽收-集货仓列表查询
type CainiaoGlobalImPickupStoresGetAPIRequest struct {
	model.Params
	// 请求体
	_transferstoreQueryRequest *TransferstoreQueryRequest
}

// NewCainiaoGlobalImPickupStoresGetRequest 初始化CainiaoGlobalImPickupStoresGetAPIRequest对象
func NewCainiaoGlobalImPickupStoresGetRequest() *CainiaoGlobalImPickupStoresGetAPIRequest {
	return &CainiaoGlobalImPickupStoresGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoGlobalImPickupStoresGetAPIRequest) Reset() {
	r._transferstoreQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupStoresGetAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.stores.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupStoresGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalImPickupStoresGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTransferstoreQueryRequest is TransferstoreQueryRequest Setter
// 请求体
func (r *CainiaoGlobalImPickupStoresGetAPIRequest) SetTransferstoreQueryRequest(_transferstoreQueryRequest *TransferstoreQueryRequest) error {
	r._transferstoreQueryRequest = _transferstoreQueryRequest
	r.Set("transferstore_query_request", _transferstoreQueryRequest)
	return nil
}

// GetTransferstoreQueryRequest TransferstoreQueryRequest Getter
func (r CainiaoGlobalImPickupStoresGetAPIRequest) GetTransferstoreQueryRequest() *TransferstoreQueryRequest {
	return r._transferstoreQueryRequest
}

var poolCainiaoGlobalImPickupStoresGetAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoGlobalImPickupStoresGetRequest()
	},
}

// GetCainiaoGlobalImPickupStoresGetRequest 从 sync.Pool 获取 CainiaoGlobalImPickupStoresGetAPIRequest
func GetCainiaoGlobalImPickupStoresGetAPIRequest() *CainiaoGlobalImPickupStoresGetAPIRequest {
	return poolCainiaoGlobalImPickupStoresGetAPIRequest.Get().(*CainiaoGlobalImPickupStoresGetAPIRequest)
}

// ReleaseCainiaoGlobalImPickupStoresGetAPIRequest 将 CainiaoGlobalImPickupStoresGetAPIRequest 放入 sync.Pool
func ReleaseCainiaoGlobalImPickupStoresGetAPIRequest(v *CainiaoGlobalImPickupStoresGetAPIRequest) {
	v.Reset()
	poolCainiaoGlobalImPickupStoresGetAPIRequest.Put(v)
}
